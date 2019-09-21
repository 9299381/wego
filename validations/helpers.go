package validations

import "errors"

func Valid(obj interface{}) error {
	valid := Validation{}
	b, _ := valid.Valid(obj)
	if !b {
		msg := ""
		for _, err := range valid.Errors {
			msg += err.Field + ":" + err.Message + ";"
		}
		return errors.New(msg)
	}
	return nil
}
