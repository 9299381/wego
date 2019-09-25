package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type HttpRouter struct {
	*servers.HttpCommServer
}

func (it *HttpRouter) Boot() {
	it.HttpCommServer = servers.NewHttpCommServer()
}

func (it *HttpRouter) Register() {

	it.Get("/demo/one", wego.Handler("one"))
	it.Get("/demo/two", wego.Handler("two"))
	it.Post("/demo/auth", wego.Handler("auth"))
	it.Get("/demo/sql", wego.Handler("sql"))
	it.Get("/demo/redis", wego.Handler("redis"))
	it.Post("/demo/post", wego.Handler("post"))
	it.Get("/demo/queue", wego.Handler("queue"))

	it.Get("/demo/cache_set", wego.Handler("cache_set"))
	it.Get("/demo/cache_get", wego.Handler("cache_get"))
	//验证validate
	it.Get("/demo/valid", wego.Handler("valid"))

	it.Get("/demo/consul", wego.Handler("consul"))
	it.Get("/demo/event", wego.Handler("event"))

}
