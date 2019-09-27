package service

import (
	"github.com/9299381/wego/cache"
	"github.com/9299381/wego/contracts"
)

type CacheSetServioce struct {
	next contracts.IService
}

func (it *CacheSetServioce) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}

func (it *CacheSetServioce) Handle(ctx contracts.Context) error {

	v := make(map[string]interface{})
	v["aaa"] = "bbb"
	v["ccc"] = "ddd"
	_ = cache.Set("aaaaa", v, 60)

	return it.next.Handle(ctx)
}
