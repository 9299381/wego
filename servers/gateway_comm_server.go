package servers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/9299381/wego"
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/filters"
	"github.com/9299381/wego/tools/errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"
)

type GateWayCommServer struct {
	e map[string]endpoint.Endpoint
}

func NewGateWayCommServer() *GateWayCommServer {
	return &GateWayCommServer{
		e: map[string]endpoint.Endpoint{},
	}
}

func (it *GateWayCommServer) Route(method, path string, endpoint endpoint.Endpoint) {
	key := method + "_" + path
	it.e[key] = endpoint
}

func (it *GateWayCommServer) Load() {
	//注册通用路由
	it.Route("GET", "/health", (&filters.HealthEndpoint{}).Make())

}

func (it *GateWayCommServer) Start() error {
	config := (&configs.HttpConfig{}).Load().(*configs.HttpConfig)
	address := config.HttpHost + ":" + config.HttpPort
	wego.App.Logger.Info("Http Server Start ", address)
	handler := it
	return http.ListenAndServe(address, handler)
}

func (it *GateWayCommServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	var resp contracts.Response
	//通过编解码 进行 路由路由处理
	ctx := r.Context()
	req, err := it.decodeRequest(ctx, r)
	if err != nil {
		resp = contracts.ResponseFailed(err)
		return
	}
	key := req.Method + "_" + r.URL.Path
	filter, ok := it.e[key]
	if ok && filter != nil {
		// 如果有注册管理,则注册管理处理
		//注意filter的endpoint可以只过滤,不进行service处理,
		// gateway_endpoint负责返回GATEWAY,h或者error
		resp = it.runFilter(filter, ctx, req)

	}
	if !ok || resp.Data == "GATEWAY" {

		defer func() {
			//注意这里开始记录 外部请求结束
			//todo 发送到内部队列 taskserver
			//请求记录:时间,时长,url,汇总,grpc/http,异步记录数据库
			fmt.Printf("defer")
		}()

		//服务发现
		entity, err := it.getConsulEntity(req.Service)
		if err != nil {
			resp = contracts.ResponseFailed(err)
		}
		tag := entity.Service.Tags[rand.Int()%len(entity.Service.Tags)]
		host := fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
		if tag == "http" {
			director := func(r *http.Request) {
				r.URL.Scheme = "http"
				r.URL.Host = host
				r.URL.Path = req.Dest
				r.Method = req.Method
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
	err = it.encodeResponse(ctx, w, resp)
	if err != nil {
		panic(err)
		return
	}
}

func (it *GateWayCommServer) decodeRequest(ctx context.Context, r *http.Request) (*contracts.GateWayRequest, error) {

	var req *contracts.GateWayRequest
	//按照分隔符'/'对路径进行分解，获取服务名称service 和 路由名称
	_ = r.ParseForm()
	vars := r.Form
	requestId, ok := vars["request_id"]
	if ok == false {
		requestId = make([]string, 1)
		requestId[0] = uuid.NewV4().String()
	}
	data := make(map[string]interface{})
	for k, v := range vars {
		data[k] = v[0]
	}
	if strings.Index(r.RemoteAddr, "::") > 0 {
		data["client_ip"] = "127.0.0.1"
	} else {
		data["client_ip"] = r.RemoteAddr
	}
	if autoken := r.Header.Get("autoken"); autoken != "" {
		data["autoken"] = autoken
	}

	var service, route, dest string
	pathArray := strings.Split(r.URL.Path, "/")
	if len(pathArray) <= 2 {
		//这是本地的
		//如果是health,则返回 SERVING
		service = pathArray[1]
		route = ""
		dest = r.URL.Path

	} else {
		service = pathArray[1]
		route = ""
		for _, v := range pathArray[2:] {
			route += v + "."
		}
		route = route[:len(route)-1]
		dest = "/" + strings.Join(pathArray[2:], "/")
	}
	req = &contracts.GateWayRequest{
		Dest:    dest,
		Method:  r.Method,
		Id:      requestId[0],
		Service: service,
		Route:   route,
		Data:    data,
	}
	return req, nil
}
func (it *GateWayCommServer) encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,authToken")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	return json.NewEncoder(w).Encode(response)
}

func (it *GateWayCommServer) runFilter(filter endpoint.Endpoint, ctx context.Context, req *contracts.GateWayRequest) contracts.Response {
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

func (it *GateWayCommServer) getConsulEntity(service string) (*api.ServiceEntry, error) {
	client := clients.GetConsullClient()
	entitys, _, err := client.Service(service, "", false, &api.QueryOptions{})
	if err != nil || len(entitys) == 0 {
		return nil, errors.New("9999", "没有找到响应的服务")
	}
	entity := entitys[rand.Int()%len(entitys)]
	return entity, nil
}
