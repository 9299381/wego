package wego

import (
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/readers"
	"github.com/coocood/freecache"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-xorm/xorm"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Status bool

	MySql        *xorm.Engine
	Redis 	  *redis.Pool
	Logger *logrus.Logger
	Cache  *freecache.Cache


	env    map[string]interface{}
	config    map[string]contracts.Iconfig
	handler map[string]endpoint.Endpoint
	routers map[string]contracts.IRouter
}


var App *Application

//初始化成全局变量
func init() {
	App = &Application{
		Status:true,
		config: make(map[string]contracts.Iconfig),
		env:    make(map[string]interface{}),
		handler : make(map[string]endpoint.Endpoint),
		routers: make(map[string]contracts.IRouter),
	}
	App.env = App.Read(&readers.IniReader{})
}



func (it *Application) Read(reader contracts.IReader) map[string]interface{} {
	data := reader.Read(".env").(map[string]map[string]interface{})
	ret, _ := data["common"]
	envSection, _ := data[ret["APP_ENV"].(string)]
	for k, v := range envSection {
		ret[k] = v
	}
	return ret
}