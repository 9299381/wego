package clients

import (
	"errors"
	"fmt"
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"os"
	"strconv"
)

func NewConsulHttpRegister(name, host, port string) *consul.Registrar {

	check := api.AgentServiceCheck{
		HTTP:     "http://" + host + ":" + port + "/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status.",
	}
	p, _ := strconv.Atoi(port)
	reg := api.AgentServiceRegistration{
		ID:      name + "_" + wego.ID(),
		Name:    name,
		Address: host,
		Port:    p,
		Tags:    []string{"http"},
		Check:   &check,
	}
	registy := consul.NewRegistrar(GetConsullClient(), &reg, getLogger())
	registy.Register()
	return registy

}
func NewConsulGrpcRegister(name, host, port string) *consul.Registrar {

	p, _ := strconv.Atoi(port)
	check := api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%s:%d/%s", host, p, "health"),
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status.",
	}
	reg := api.AgentServiceRegistration{
		ID:      name + "_" + wego.ID(),
		Name:    name,
		Address: host,
		Port:    p,
		Tags:    []string{"grpc"},
		Check:   &check,
	}
	registy := consul.NewRegistrar(GetConsullClient(), &reg, getLogger())
	registy.Register()
	return registy

}

func GetConsullClient() consul.Client {
	var client consul.Client
	{
		config := api.DefaultConfig()
		config.Address = args.Registy
		consulClient, err := api.NewClient(config)
		if err != nil {
			fmt.Println("create consul client error:", err)
			os.Exit(1)
		}
		client = consul.NewClient(consulClient)
	}
	return client
}

func getLogger() log.Logger {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	return logger
}

func GetConsulService(service string) (*api.ServiceEntry, error) {
	//这里考虑可以从缓存中读取,10分钟过期,比如
	client := GetConsullClient()
	entitys, _, err := client.Service(service, "", false, &api.QueryOptions{})
	if err != nil || len(entitys) == 0 {
		return nil, errors.New("9999::没有找到响应的服务")
	}
	entity := entitys[rand.Int()%len(entitys)]
	return entity, nil

}
