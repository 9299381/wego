package convert

import "reflect"

func Struct2Map(obj interface{}) map[string]interface{} {
	elem := reflect.ValueOf(obj).Elem()
	relType := elem.Type()

	var data = make(map[string]interface{})
	for i := 0; i < relType.NumField(); i++ {
		data[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return data
}