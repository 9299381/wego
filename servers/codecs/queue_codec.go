package codecs

import (
	"context"
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

func QueueServerDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	id := wego.ID()
	return contracts.Request{
		Id:   id,
		Data: req.(map[string]interface{}),
	}, nil
}

func QueueServerEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
