package filters

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/loggers"
	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

type CommEndpoint struct {
	Service contracts.IService
	next    endpoint.Endpoint
}

func (it *CommEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	it.next = next
	return it
}

func (it *CommEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(contracts.Request)
		cc := it.makeContext(ctx, req)
		cc.Log = it.makeLog(cc, req)
		err = it.Service.Handle(cc)
		if err != nil {
			cc.Log.Info(err.Error())
			return contracts.ResponseFailed(err), nil
		} else {
			return contracts.ResponseSucess(cc.GetValue("response")), nil
		}
	}
}

func (it *CommEndpoint) makeLog(ctx contracts.Context, req contracts.Request) *logrus.Entry {
	//初始化日志字段,放到context中
	ip := (req.Data)["client_ip"]
	if ip == nil {
		ip = "LAN"
	}
	entity := loggers.Log.WithFields(logrus.Fields{
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
