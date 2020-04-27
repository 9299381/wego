package controller

import (
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/contracts"
	"github.com/tidwall/gjson"
)

type CacheGetController struct {
}

func (s *CacheGetController) Handle(ctx contracts.Context) (interface{}, error) {
	//GetByte 方式
	jsonBytes, _ := cache.GetByte("key")
	// Get方式
	obj := make(map[string]interface{})
	_ = cache.Get("key", &obj)

	//使用gjson 更方便 直接从json字符串中取值
	return map[string]interface{}{
		"ccc":  gjson.Get(string(jsonBytes), "ccc.a").Str,
		"aaa":  obj["aaa"],
		"gaaa": gjson.Get(string(jsonBytes), "aaa").String(),
	}, nil
}
