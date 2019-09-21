package convert

import (
	"errors"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/validations"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	elem := reflect.ValueOf(obj).Elem()
	relType := elem.Type()

	var data = make(map[string]interface{})
	for i := 0; i < relType.NumField(); i++ {
		data[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return data
}

func Map2Struct(req, obj interface{}) error {
	request, ok := req.(map[string]interface{})
	if ok == false {
		return errors.New(constants.ErrConvert)
	}
	err := mapstructure.WeakDecode(request, obj)
	if err != nil {
		return errors.New(constants.ErrConvert)
	}
	err = validations.Valid(obj)
	if err != nil {
		return err
	}
	return nil
}
