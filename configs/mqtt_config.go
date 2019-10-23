package configs

import "strconv"

type MqttConfig struct {
	Host         string `json:"host"`
	UserName     string `json:"user_name"`
	PassWord     string `json:"pass_word"`
	Parallel     bool   `json:"parallel"`
	SubscribeQos uint8  `json:"subscribe_qos"`
	PublishQos   uint8  `json:"publish_qos"`
}

func (s *MqttConfig) Load() *MqttConfig {
	var b bool
	if Env("MQTT_PARALLEL", "no") == "yes" {
		b = true
	} else {
		b = false
	}
	subscribeQos, _ := strconv.Atoi(Env("SUBSCRIBE_QOS", "2"))
	publishQos, _ := strconv.Atoi(Env("PUBLISH_QOS", "2"))
	config := &MqttConfig{
		Host:         Env("MQTT_HOST", "tcp://127.0.0.1:1883"),
		UserName:     Env("MQTT_USERNAME", ""),
		PassWord:     Env("MQTT_PASSWORD", ""),
		Parallel:     b,
		SubscribeQos: uint8(subscribeQos),
		PublishQos:   uint8(publishQos),
	}
	return config
}
