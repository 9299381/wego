package convert

import (
	"encoding/json"
	"errors"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/validations"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"reflect"
	"strconv"
	"strings"
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
	// map转struct时, key与struct的字段应该相同,忽略大小写,
	// 注意 不可随意增加_,json中可以有_
	request, ok := req.(map[string]interface{})
	if ok == false {
		return errors.New(constants.ErrConvert)
	}
	for k, v := range request {
		if strings.Contains(k, "_") {
			kk := strings.ReplaceAll(k, "_", "")
			request[kk] = v
		}
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

func FormEncode(params map[string]interface{}) url.Values {
	data := url.Values{}
	for k, param := range params {
		paramsType := reflect.TypeOf(param)
		switch paramsType.String() {
		case "string":
			data.Set(k, param.(string))
		case "int":
			data.Set(k, strconv.Itoa(param.(int)))
		default:
			str, _ := json.Marshal(param)
			data.Set(k, string(str))

		}
	}
	return data
}
