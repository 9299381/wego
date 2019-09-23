package wego

import (
	"encoding/json"
	"fmt"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/queues"
	"github.com/9299381/wego/tools/snowflake"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-xorm/xorm"
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

//快捷入口

func ID() string {
	serverId, _ := strconv.Atoi(Env("SERVER_ID", "512"))
	return snowflake.GetID(int64(serverId))
}
func DB() *xorm.Engine {
	return App.MySql
}

func Redis() redis.Conn {
	return App.Redis.Get()
}

func Queue(name string, router string, params map[string]interface{}) error {
	conn := Redis()
	defer conn.Close()
	job := &queues.Job{
		Queue: name,
		Payload: queues.Payload{
			Route:  router,
			Params: params,
		},
	}
	prefix := Env("QUEUE_PREFIX", "wego")
	return queues.Enqueue(conn, job, prefix)

}

func Env(key string, value ...interface{}) string {
	ret, exist := App.Env[key]
	if exist {
		return ret.(string)
	} else {
		return value[0].(string)
	}
}
func Config(key string, conf ...contracts.Iconfig) contracts.Iconfig {
	if conf == nil {
		//这里是get
		obj, exist := App.config[key]
		if exist {
			return obj
		}
	} else {
		//这里是set
		App.config[key] = conf[0].Load()
	}
	return nil
}

func Cache(key string, value ...interface{}) []byte {
	if value == nil {
		//这里是get
		k := []byte(key)
		ret, err := App.Cache.Get(k)
		if err != nil {
			return nil
		} else {
			return ret
		}
	} else {
		k := []byte(key)
		val := value[0]
		v, err := json.Marshal(val)
		if err != nil {
			return nil
		}
		var expire int
		if value[1] != nil {
			expire = value[1].(int)
		} else {
			expire = 60
		}
		_ = App.Cache.Set(k, v, expire)
	}
	return nil
}

func Provider(p contracts.IProvider) {
	p.Boot()
	p.Register()
}

func Handler(name string, endpoint ...endpoint.Endpoint) endpoint.Endpoint {
	if endpoint == nil {
		ret, exist := App.handler[name]
		if exist {
			return ret
		}
	} else {
		App.handler[name] = endpoint[0]
	}
	return nil
}

func Router(name string, server contracts.IRouter) {
	server.Boot()
	server.Load()
	server.Register()
	App.routers[name] = server
}

//启动server
func Start() {

	servers := strings.Split(args.Server, ",")
	routers := make(map[string]contracts.IRouter)

	for _, s := range servers {
		if ss, exist := App.routers[strings.Trim(s, " ")]; exist == true {
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
		App.Logger.Info(<-errChan)
	}
}
