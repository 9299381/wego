package wego

import (
	"github.com/9299381/wego/contracts"
	"github.com/coocood/freecache"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-xorm/xorm"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Status bool
	//定制加载
	Logger *logrus.Logger
	MySql  *xorm.Engine
	Redis  *redis.Pool
	Cache  *freecache.Cache
	Event  *contracts.EventStack
	//必须初始化
	Consul  map[string]*consul.Registrar
	Env     map[string]string
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
		Env:     make(map[string]string),
		Consul:  make(map[string]*consul.Registrar),
		handler: make(map[string]endpoint.Endpoint),
		routers: make(map[string]contracts.IRouter),
	}
}
