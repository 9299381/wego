package controller

import (
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/contracts"
)

type CacheSetController struct {
	*contracts.Controller
}

func (s *CacheSetController) Handle(ctx contracts.Context) (interface{}, error) {

	v := make(map[string]interface{})
	v["aaa"] = "bbb"
	v["ccc"] = "ddd"
	_ = cache.Set("aaaaa", v, 60)

	return nil, nil
}
