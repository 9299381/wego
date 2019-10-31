package configs

type MqttConfig struct {
	Host         string `json:"host"`
	UserName     string `json:"user_name"`
	PassWord     string `json:"pass_word"`
	Parallel     bool   `json:"parallel"`
	SubscribeQos uint8  `json:"subscribe_qos"`
	PublishQos   uint8  `json:"publish_qos"`
}

func LoadMqttConfig() *MqttConfig {
	config := &MqttConfig{
		Host:         EnvString("mqtt.host", "tcp://127.0.0.1:1883"),
		UserName:     EnvString("mqtt.username", ""),
		PassWord:     EnvString("mqtt.password", ""),
		Parallel:     EnvBool("mqtt.parallel", false),
		SubscribeQos: uint8(EnvInt("mqtt.subscribe_qos", 2)),
		PublishQos:   uint8(EnvInt("mqtt.publish_qos", 2)),
	}
	return config
}
