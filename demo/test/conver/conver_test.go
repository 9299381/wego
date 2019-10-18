package test

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"testing"
)

func TestMap2Struct(t *testing.T) {

	m := getMap()
	obj := &DTOStruct{}
	//要去掉map中key的下划线
	err := mapstructure.WeakDecode(m, obj)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(obj)
	}

}

func getMap() map[string]interface{} {

	//jjson := "[{\"id\":\"123123\",\"name\":\"abcde\"},{\"id\":\"234234\",\"name\":\"asdfasdf\"}]"

	m := make(map[string]interface{})
	//m["name"] = "asdf"
	//m["age"] = "123"
	//m["desc"] = "随便练"
	//m["Username"] = string(getJson())
	m["user_name"] = string(getJson())

	return m
}

func getJson() []byte {

	var jj []map[string]interface{}
	ss1 := make(map[string]interface{})
	ss1["id"] = "123123"
	ss1["name"] = "abcde"
	ss2 := make(map[string]interface{})
	ss2["id"] = "234234"
	ss2["name"] = "asdfasdf"
	jj = append(jj, ss1)
	jj = append(jj, ss2)
	jjson, _ := json.Marshal(jj)
	return jjson

}

type DTOStruct struct {
	User_Name string `json:"user_name"`
}
