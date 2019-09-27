package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/tools/idwork"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"strconv"
)

func NewConsulHttpRegister(service, host, port string) *consul.Registrar {
	//注销掉重复的
	serviceDeregister(service, host, port)
	check := api.AgentServiceCheck{
		HTTP:     "http://" + host + ":" + port + "/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status.",
	}
	p, _ := strconv.Atoi(port)
	reg := api.AgentServiceRegistration{
		ID:      service + "_" + idwork.ID(),
		Name:    service,
		Address: host,
		Port:    p,
		Tags:    []string{"http"},
		Check:   &check,
	}
	registy := consul.NewRegistrar(getConsullClient(), &reg, loggers.NewKitLog())
	registy.Register()
	return registy

}
func NewConsulGrpcRegister(service, host, port string) *consul.Registrar {
	//注销掉重复的
	serviceDeregister(service, host, port)
	p, _ := strconv.Atoi(port)
	check := api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%s:%d/%s", host, p, "health"),
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status.",
	}
	reg := api.AgentServiceRegistration{
		ID:      service + "_" + idwork.ID(),
		Name:    service,
		Address: host,
		Port:    p,
		Tags:    []string{"grpc"},
		Check:   &check,
	}
	registy := consul.NewRegistrar(getConsullClient(), &reg, loggers.NewKitLog())
	registy.Register()
	return registy

}

func GetConsulService(service string) (entity *api.ServiceEntry, err error) {
	//这里考虑可以从缓存中读取,10分钟过期,比如
	var entitys []*api.ServiceEntry
	c, _ := cache.Get("consul_entitys")
	if c != nil {
		entitys = []*api.ServiceEntry{}
		err = json.Unmarshal(c, &entitys)
		if err != nil {
			panic(err)
			return
		}
	} else {
		client := getConsullClient()
		entitys, _, err = client.Service(service, "", false, &api.QueryOptions{})
		if err != nil || len(entitys) == 0 {
			err = errors.New("9999::没有找到响应的服务")
			return
		}
		_ = cache.Set("consul_entitys", entitys, 60)
	}
	//随机取一个
	entity = entitys[rand.Int()%len(entitys)]
	return

}

func getConsullClient() consul.Client {
	var client consul.Client
	{
		config := api.DefaultConfig()
		config.Address = args.Registy
		consulClient, _ := api.NewClient(config)
		client = consul.NewClient(consulClient)
	}
	return client
}

func serviceDeregister(service, host, port string) {
	client := getConsullClient()
	entitys, _, err := client.Service(service, "", false, &api.QueryOptions{})
	if err == nil {
		for _, entity := range entitys {
			str1 := fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
			str2 := fmt.Sprintf("%s:%s", host, port)
			if str1 == str2 {
				r := &api.AgentServiceRegistration{
					ID:   entity.Service.ID,
					Name: service,
				}
				_ = client.Deregister(r)
			}
		}
	}
}
