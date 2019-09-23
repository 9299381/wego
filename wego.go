package wego

import (
	"github.com/9299381/wego/contracts"
	"github.com/coocood/freecache"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-xorm/xorm"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Status bool

	MySql  *xorm.Engine
	Redis  *redis.Pool
	Logger *logrus.Logger
	Cache  *freecache.Cache

	Env     map[string]interface{}
	config  map[string]contracts.Iconfig
	handler map[string]endpoint.Endpoint
	routers map[string]contracts.IRouter
}

var App *Application

//初始化成全局变量
func init() {
	App = &Application{
		Status:  true,
		config:  make(map[string]contracts.Iconfig),
		Env:     make(map[string]interface{}),
		handler: make(map[string]endpoint.Endpoint),
		routers: make(map[string]contracts.IRouter),
	}
}
