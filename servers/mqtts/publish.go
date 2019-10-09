package mqtts

import "encoding/json"

func Publish(topic string, payload interface{}) error {

	param, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	token := mc.Publish(topic, 0, false, param)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
