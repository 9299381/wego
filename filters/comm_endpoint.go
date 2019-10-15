package filters

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/tools/convert"
	"github.com/9299381/wego/validations"
	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

type CommEndpoint struct {
	Controller contracts.IController
	next       endpoint.Endpoint
}

func (it *CommEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	it.next = next
	return it
}

func (it *CommEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//生成请求参数
		req := request.(contracts.Request)
		//生成context上下文
		cc := it.makeContext(ctx, req)
		//成成线程log,统一处理ip,request_id等
		cc.Log = it.makeLog(cc, req)
		//参数验证
		err := it.valid(cc, req)
		if err != nil {
			cc.Log.Info(err.Error())
			return nil, err
		}
		//逻辑处理
		ret, err := it.Controller.Handle(cc)
		if err != nil {
			cc.Log.Info(err.Error())
		}
		return ret, err
	}
}
func (it *CommEndpoint) valid(ctx contracts.Context, request contracts.Request) error {
	obj := it.Controller.GetRules()
	if obj != nil {
		err := convert.Map2Struct(request.Data, obj)
		if err != nil {
			return err
		}
		//验证obj
		err = validations.Valid(obj)
		if err != nil {
			return err
		}
		ctx.SetValue("RequestDTO", obj)
	}
	return nil
}

func (it *CommEndpoint) makeLog(ctx contracts.Context, req contracts.Request) *logrus.Entry {
	//初始化日志字段,放到context中
	ip := (req.Data)["client_ip"]
	if ip == nil {
		ip = "LAN"
	}
	entity := loggers.GetLog().WithFields(logrus.Fields{
		"request_id": req.Id,
		"client_ip":  ip,
	})
	return entity
}

func (it *CommEndpoint) makeContext(ctx context.Context, req contracts.Request) contracts.Context {
	cc := contracts.Context{
		Context: ctx,
		Keys:    make(map[string]interface{}),
	}
	cc.SetValue("request", req.Data)
	cc.SetValue("request.id", req.Id)

	return cc
}
