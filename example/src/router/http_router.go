package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/servers"
)

type HttpRouter struct {
	*servers.HttpCommServer
}

func (it *HttpRouter) Boot()  {
	it.HttpCommServer = servers.NewHttpCommServer()
}

func (it *HttpRouter) Register()  {

	it.Get("/exam/one", wego.Handler("one"))
	it.Get("/exam/two", wego.Handler("two"))
	it.Post("/exam/auth", wego.Handler("auth"))
	it.Get("/exam/sql", wego.Handler("sql"))
	it.Get("/exam/redis", wego.Handler("redis"))
	it.Post("/exam/post", wego.Handler("post"))
	it.Get("/exam/job", wego.Handler("job"))

	it.Get("/exam/cache_set", wego.Handler("cache_set"))
	it.Get("/exam/cache_get", wego.Handler("cache_get"))

}