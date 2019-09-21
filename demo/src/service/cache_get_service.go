package service

import (
	"encoding/json"
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

type CacheGetServioce struct {
	next contracts.IService
}

func (it *CacheGetServioce) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *CacheGetServioce) Handle(ctx contracts.Context) error {

	v := wego.Cache("aaaaa")
	d := make(map[string]interface{})
	err := json.Unmarshal(v, &d)
	if err != nil {
		//转换错误,一般为没找到
	} else {
		ctx.Log.Info(d["aaa"])
	}

	return it.next.Handle(ctx)
}
