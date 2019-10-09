package mqtts

import (
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/tools/idwork"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mc mqtt.Client

func init() {
	config := (&configs.MqttConfig{}).Load()
	opts := mqtt.NewClientOptions().AddBroker(config.Host)
	opts.SetUsername(config.UserName)
	opts.SetPassword(config.PassWord)
	opts.SetClientID(idwork.ID())
	mc = mqtt.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
