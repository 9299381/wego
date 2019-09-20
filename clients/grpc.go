package clients

import (
	"context"
	"encoding/json"
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers/transports/protobuf"
	"google.golang.org/grpc"
	"log"
)

func NewGrpcClient(service string,params interface{}) *protobuf.Response {

	//这个应该从consul中读取 ,定时任务更新 内存中存储

	serviceAddress := "127.0.0.1:9341"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	jsonParam, _ := json.Marshal(params)
	in := &protobuf.Request{
		Id:    wego.ID(),
		Param: string(jsonParam),
	}
	
	out := new(protobuf.Response)
	method := "/protobuf." + service + "/Handle"
	_ = conn.Invoke(context.Background(), method, in, out)

	return out
}
