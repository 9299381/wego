package controller

import (
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/contracts"
)

type CacheSetController struct {
}

func (s *CacheSetController) Handle(ctx contracts.Context) (interface{}, error) {

	v := make(map[string]interface{})
	v["aaa"] = "bbb"

	dd := make(map[string]interface{})
	dd["a"] = "111"
	dd["b"] = "222"

	v["ccc"] = dd
	_ = cache.Set("key", v, 60)

	return nil, nil
}
