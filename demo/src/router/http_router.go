package router

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/servers"
)

type HttpRouter struct {
	*servers.HttpCommServer
}

func (s *HttpRouter) Boot() {
	s.HttpCommServer = servers.NewHttpCommServer()
}

func (s *HttpRouter) Register() {

	s.Get("/demo/one", wego.Handler("one"))
	s.Get("/demo/two", wego.Handler("two"))
	s.Post("/demo/auth", wego.Handler("auth"))
	s.Get("/demo/sql", wego.Handler("sql"))
	s.Get("/demo/redis", wego.Handler("redis"))
	s.Post("/demo/post", wego.Handler("post"))
	s.Get("/demo/queue", wego.Handler("queue"))

	s.Get("/demo/cache_set", wego.Handler("cache_set"))
	s.Get("/demo/cache_get", wego.Handler("cache_get"))
	//验证validate
	s.Post("/demo/valid", wego.Handler("valid"))

	s.Get("/demo/consul", wego.Handler("consul"))
	s.Get("/demo/event", wego.Handler("event"))
	s.Get("/demo/publish", wego.Handler("publish"))

	if args.Mode != "prod" {
		s.HandleSwagger()
	}
}
