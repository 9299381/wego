package provider

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/demo/src/controller"
	"github.com/9299381/wego/filters"
)

type DemoProvider struct {
}

func (s *DemoProvider) Boot() {
}

func (s *DemoProvider) Register() {
	// 这种controller 可以重复再event,subscribe,queue中使用,因此注册到handler中
	//限速
	wego.Handler("one", filters.Limit(new(controller.OneController)))
	wego.Handler("two", filters.New(&controller.TwoController{}))
	wego.Handler("auth", filters.Chain(
		&filters.ResponseEndpoint{},
		&filters.JwtEndpoint{},
		&filters.LimitEndpoint{},
		&filters.CommEndpoint{
			Controller: &controller.AuthController{},
		}))

	wego.Handler("post", filters.New(&controller.PostController{}))
	wego.Handler("sql", filters.New(&controller.SqlController{}))
	wego.Handler("redis", filters.New(&controller.RedisController{}))
	wego.Handler("queue", filters.New(&controller.QueueController{}))
	wego.Handler("queue2", filters.New(&controller.Queue2Controller{}))

	wego.Handler("cache_set", filters.New(&controller.CacheSetController{}))
	wego.Handler("cache_get", filters.New(&controller.CacheGetController{}))
	//
	wego.Handler("valid", filters.New(&controller.ValidController{}))
	//
	wego.Handler("consul", filters.New(&controller.ConsulController{}))

	wego.Handler("event", filters.New(&controller.EventController{}))
	//
	wego.Handler("publish", filters.New(&controller.PublishController{}))
	wego.Handler("sleep", filters.New(&controller.SleepController{}))
	wego.Handler("mqtt_event", filters.New(&controller.MqttEventController{}))
}
