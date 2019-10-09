package configs

type MqttConfig struct {
	Host     string `json:"host"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func (it *MqttConfig) Load() *MqttConfig {

	config := &MqttConfig{
		Host:     Env("MQTT_HOST", "tcp://127.0.0.1:1883"),
		UserName: Env("MQTT_USERNAME", ""),
		PassWord: Env("MQTT_PASSWORD", ""),
	}
	return config
}
