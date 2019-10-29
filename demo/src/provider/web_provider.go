package provider

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/demo/src/web_controller"
	"github.com/9299381/wego/filters"
)

type WebProvider struct {
}

func (s *WebProvider) Boot() {
}

func (s *WebProvider) Register() {
	wego.Handler("page_one", filters.New(&web_controller.PageOne{}))
}
