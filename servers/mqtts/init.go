package mqtts

import (
	"github.com/9299381/wego/tools/idwork"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mc mqtt.Client

func init() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID(idwork.ID())
	mc = mqtt.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
