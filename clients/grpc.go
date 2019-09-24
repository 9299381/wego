package clients

import (
	"context"
	"encoding/json"
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers/transports/protobuf"
	"google.golang.org/grpc"
	"log"
)

func NewGrpcClient(serviceAddress string, service string, params interface{}) (*protobuf.Response, error) {
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
	err = conn.Invoke(context.Background(), method, in, out)
	return out, err
}
