package configs

type MqttConfig struct {
	Host     string `json:"host"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Parallel bool   `json:"parallel"`
}

func (s *MqttConfig) Load() *MqttConfig {
	var b bool
	if Env("MQTT_PARALLEL", "no") == "yes" {
		b = true
	} else {
		b = false
	}
	config := &MqttConfig{
		Host:     Env("MQTT_HOST", "tcp://127.0.0.1:1883"),
		UserName: Env("MQTT_USERNAME", ""),
		PassWord: Env("MQTT_PASSWORD", ""),
		Parallel: b,
	}
	return config
}
