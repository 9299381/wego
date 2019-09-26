package gateways

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/9299381/wego"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"
)

type Server struct {
	handlers map[string]endpoint.Endpoint
}

func NewServer() *Server {
	ss := &Server{
		handlers: map[string]endpoint.Endpoint{},
	}
	return ss
}

func (it *Server) Register(method, path string, endpoint endpoint.Endpoint) {
	key := method + "_" + path
	it.handlers[key] = endpoint
}

func (it *Server) Serve() error {
	config := (&configs.HttpConfig{}).Load().(*configs.HttpConfig)
	address := config.HttpHost + ":" + config.HttpPort
	wego.App.Logger.Info("Http Server Start ", address)
	handler := it
	return http.ListenAndServe(address, handler)
}

func (it *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	var resp contracts.Response
	//通过编解码 进行 路由路由处理
	ctx := r.Context()
	req, err := decodeRequest(ctx, r)
	if err != nil {
		resp = contracts.ResponseFailed(err)
		return
	}
	key := req.Method + "_" + r.URL.Path

	filter, ok := it.handlers[key]
	if ok && filter != nil {
		// 如果有注册管理,则注册管理处理
		//注意filter的endpoint可以只过滤,不进行service处理,
		// gateway_endpoint负责返回GATEWAY,h或者error
		resp = it.runFilter(filter, ctx, req)
	}
	if !ok || resp.Data == "GATEWAY" {
		var tag, host string
		defer func(begin time.Time, tag, host *string) {
			params := make(map[string]interface{})
			params["url"] = key
			params["begin"] = begin.Format(constants.YmdHis)
			params["took"] = time.Since(begin)
			params["tag"] = *tag
			params["host"] = *host
			payload := &contracts.Payload{
				Route:  wego.Env("GATEWAY_EVENT_HANDLER", "GATEWAY_EVENT_HANDLER"),
				Params: params,
			}
			wego.Event(payload)

		}(time.Now(), &tag, &host)
		//服务发现
		entity, err := clients.GetConsulService(req.Service)
		if err != nil {
			resp = contracts.ResponseFailed(err)
			return
		}
		tag = entity.Service.Tags[rand.Int()%len(entity.Service.Tags)]
		host = fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
		if tag == "http" {
			//
			director := func(dr *http.Request) {
				dr.URL.Scheme = "http"
				dr.URL.Host = host
				dr.URL.Path = req.Dest
				dr.Method = req.Method
			}
			gateway := &httputil.ReverseProxy{Director: director}
			gateway.ServeHTTP(w, r)
			return

		} else if tag == "grpc" {
			gc, err := clients.NewGrpcClient(host, req.Route, req.Data)
			if err != nil {
				resp = contracts.ResponseFailed(err)
			} else {
				m := make(map[string]interface{})
				err := json.Unmarshal([]byte(gc.GetData()), &m)
				m["gateway"] = "grpc"
				if err != nil {
					resp = contracts.ResponseFailed(err)
				} else {
					resp.Code = gc.GetCode()
					resp.Ret = 200
					resp.Data = m
					resp.Message = gc.GetMsg()
				}
			}
		}
	}
	err = encodeResponse(ctx, w, resp)
	if err != nil {
		panic(err)
		return
	}
}

func (it *Server) runFilter(filter endpoint.Endpoint, ctx context.Context, req *contracts.GateWayRequest) contracts.Response {
	filterResp, err := filter(ctx, contracts.Request{
		Id:   req.Id,
		Data: req.Data,
	})
	if err != nil {
		return contracts.ResponseFailed(err)
	} else {
		return filterResp.(contracts.Response)
	}
}
