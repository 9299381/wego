package controller

import (
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/contracts"
)

type CacheSetController struct {
}

func (it *CacheSetController) Handle(ctx contracts.Context) (interface{}, error) {

	v := make(map[string]interface{})
	v["aaa"] = "bbb"
	v["ccc"] = "ddd"
	_ = cache.Set("aaaaa", v, 60)

	return nil, nil
}

func (it *CacheSetController) Valid(ctx contracts.Context) error {
	return nil
}
