package controller

import (
	"github.com/9299381/wego/contracts"
	"time"
)

type MqttEventController struct {
}

func (s *MqttEventController) Handle(ctx contracts.Context) (interface{}, error) {
	if ctx.Get("request.connected_at") != nil {
		connect := int64(ctx.Get("request.connected_at").(float64))
		var conn string = time.Unix(connect, 0).Format("2006-01-02 15:04:05")
		println(ctx.Get("request.clientid").(string), "connect_at", conn)
	}
	if ctx.Get("request.disconnected_at") != nil {
		connect := int64(ctx.Get("request.disconnected_at").(float64))
		var disconn string = time.Unix(connect, 0).Format("2006-01-02 15:04:05")
		println(ctx.Get("request.clientid").(string), "disconnect_at", disconn)
	}
	return nil, nil
}
