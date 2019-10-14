package mqtts

import (
	"encoding/json"
	"errors"
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
	token := GetIns().Publish(topic, 0, false, param)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
