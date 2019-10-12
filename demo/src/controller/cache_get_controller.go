package controller

import (
	"encoding/json"
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/contracts"
)

type CacheGetController struct {
}

func (it *CacheGetController) Handle(ctx contracts.Context) (interface{}, error) {

	v, _ := cache.Get("aaaaa")
	d := make(map[string]interface{})
	err := json.Unmarshal(v, &d)
	if err != nil {
		return nil, err
	} else {
		ctx.Log.Info(d["aaa"])
	}

	return nil, nil
}

func (it *CacheGetController) Valid(ctx contracts.Context) error {
	return nil
}
