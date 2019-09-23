package codecs

import (
	"context"
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

func WebSocketDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	id := wego.ID()
	return contracts.Request{
		Id:   id,
		Data: req.(map[string]interface{}),
	}, nil
}

func WebSocketEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
