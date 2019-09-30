package events

import (
	"github.com/9299381/wego/contracts"
)

func Fire(payload *contracts.Payload) {
	//发送事件需要判断是否有处理器,否则不处理
	_, isExist := Handlers[payload.Route]
	if isExist {
		event := newEvent(payload)
		addEvent(event)
	}
}
