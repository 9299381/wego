package controller

import (
	"encoding/json"
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/contracts"
	"github.com/tidwall/gjson"
)

type CacheGetController struct {
}

func (s *CacheGetController) Handle(ctx contracts.Context) (interface{}, error) {

	v, _ := cache.Get("aaaaa")
	d := make(map[string]interface{})
	err := json.Unmarshal(v, &d)
	if err != nil {
		return nil, err
	} else {
		ctx.Log.Info(d["aaa"])
	}

	//使用gjson 更方便
	return map[string]interface{}{
		"ccc":  gjson.Get(string(v), "ccc").Str,
		"aaa":  d["aaa"],
		"gaaa": gjson.Get(string(v), "aaa").String(),
	}, nil
}
