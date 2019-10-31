package mqtts

import (
	"encoding/json"
	"errors"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/constants"
)

func Publish(topic string, payload interface{}) error {
	if GetIns() == nil {
		return errors.New(constants.ErrMQTTConnect)
	}
	param, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	config := configs.LoadMqttConfig()
	token := GetIns().Publish(topic, config.PublishQos, false, param)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
