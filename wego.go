package wego

import (
	"fmt"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/loggers"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd/consul"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Application struct {
	Status bool
	//必须初始化
	Service  map[string]contracts.IService
	Consul   map[string]*consul.Registrar
	Handlers map[string]endpoint.Endpoint
	Routers  map[string]contracts.IRouter
}

var App *Application

//初始化成全局变量
func init() {
	App = &Application{
		Status:   true,
		Service:  make(map[string]contracts.IService),
		Consul:   make(map[string]*consul.Registrar),
		Handlers: make(map[string]endpoint.Endpoint),
		Routers:  make(map[string]contracts.IRouter),
	}
}

func Provider(p contracts.IProvider) {
	p.Boot()
	p.Register()
}

func Handler(name string, endpoint ...endpoint.Endpoint) endpoint.Endpoint {
	if endpoint == nil {
		ret, exist := App.Handlers[name]
		if exist {
			return ret
		}
	} else {
		App.Handlers[name] = endpoint[0]
	}
	return nil
}

func Service(name string, service ...contracts.IService) contracts.IService {
	if service == nil {
		ret, exist := App.Service[name]
		if exist {
			return ret
		}
	} else {
		App.Service[name] = service[0]
	}
	return nil
}

func Router(name string, server contracts.IRouter) {
	server.Boot()
	server.Load()
	server.Register()
	App.Routers[name] = server
}

//启动server
func Start() {
	servers := strings.Split(args.Server, ",")
	routers := make(map[string]contracts.IRouter)

	for _, s := range servers {
		if ss, exist := App.Routers[strings.Trim(s, " ")]; exist == true {
			routers[s] = ss
		}
	}
	errChans := make(map[string]chan error)
	for key, router := range routers {
		errChans[key] = make(chan error)
		go func(errChan chan error, server contracts.IRouter) {
			errChan <- server.Start()
		}(errChans[key], router)
		go func(errChan chan error) {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			errChan <- fmt.Errorf("%s", <-c)
		}(errChans[key])
	}
	for _, errChan := range errChans {
		loggers.GetLog().Info(<-errChan)
	}
	//关闭各种路由服务
	for _, server := range routers {
		server.Close()
	}

}
