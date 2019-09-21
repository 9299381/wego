package provider

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/demo/src/service"
	"github.com/9299381/wego/filters"
	"github.com/9299381/wego/services"
)

type ExamProvider struct {
}

func (it *ExamProvider) Boot() {
}

func (it *ExamProvider) Register() {

	wego.Handler("one", filters.Limit(
		services.Chain(
			&service.OneService{},
			&service.TwoService{},
		)),
	)

	wego.Handler("auth", filters.Chain(
		&filters.ResponseEndpoint{},
		&filters.JwtEndpoint{},
		&filters.LimitEndpoint{},
		&filters.CommEndpoint{
			Service: services.Chain(
				&service.AuthService{},
			)}))

	wego.Handler("two", filters.New(services.Chain(&service.TwoService{})))
	wego.Handler("post", filters.New(services.Chain(&service.PostService{})))
	wego.Handler("sql", filters.New(services.Chain(&service.SqlService{})))
	wego.Handler("redis", filters.New(services.Chain(&service.RedisService{})))
	wego.Handler("job", filters.New(services.Chain(&service.TestJob{})))
	wego.Handler("cache_set", filters.New(services.Chain(&service.CacheSetServioce{})))
	wego.Handler("cache_get", filters.New(services.Chain(&service.CacheGetServioce{})))

	wego.Handler("valid", filters.New(services.Chain(&service.ValidService{})))

}
