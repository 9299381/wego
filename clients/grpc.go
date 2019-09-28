package clients

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/transports/protobuf"
	"github.com/9299381/wego/tools/idwork"
	"google.golang.org/grpc"
	"log"
)

func NewGrpcClient(serviceAddress string, service string, params map[string]interface{}) (*protobuf.Response, error) {
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	jsonParam, _ := json.Marshal(params)
	in := &protobuf.Request{
		Id:    idwork.ID(),
		Param: string(jsonParam),
	}

	out := new(protobuf.Response)

	method := "/protobuf." + service + "/Handle"
	err = conn.Invoke(context.Background(), method, in, out)
	return out, err
}

func NewGrpcCall(host, service string, params map[string]interface{}) (ret contracts.Response) {
	resp, err := NewGrpcClient(host, service, params)
	if err != nil {
		ret = contracts.ResponseFailed(errors.New("没有响应的服务:" + service))
	} else {
		m := make(map[string]interface{})
		m["call_method"] = "grpc"
		err := json.Unmarshal([]byte(resp.GetData()), &m)
		if err != nil {
			ret = contracts.ResponseFailed(err)
		} else {
			ret.Code = resp.Code
			ret.Ret = 200
			ret.Message = resp.Msg
			ret.Data = m
		}
	}
	return
}
