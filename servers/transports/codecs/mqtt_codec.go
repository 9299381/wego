package codecs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/idwork"
)

func MqttSubscribeDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal(req.([]byte), &mapResult)
	if err != nil {
		return nil, errors.New(constants.ErrJson)
	}
	requestId, ok := mapResult["request_id"].(string)
	if ok == false {
		requestId = idwork.ID()
	}
	return contracts.Request{
		Id:   requestId,
		Data: mapResult,
	}, nil
}

func MqttSubscribeEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
