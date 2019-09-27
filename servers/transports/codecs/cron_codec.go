package codecs

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/idwork"
)

func CronDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(map[string]interface{})
	request["request_id"] = idwork.ID()
	return contracts.Request{
		Id:   request["request_id"].(string),
		Data: request,
	}, nil
}

func CronEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
