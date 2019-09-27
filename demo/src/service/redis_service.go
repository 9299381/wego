package service

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/contracts"
	"github.com/gomodule/redigo/redis"
)

type RedisService struct {
	next contracts.IService
}

func (it *RedisService) Next(srv contracts.IService) contracts.IService {
	it.next = srv
	return it
}
func (it *RedisService) Handle(ctx contracts.Context) error {
	client := clients.Redis() //从pool中获取一个链接
	defer client.Close()      //延时释放链接,本方法执行完毕时释放
	_, _ = client.Do("SET", "go_key", "value")
	res, _ := redis.String(client.Do("GET", "go_key"))
	exists, _ := redis.Bool(client.Do("EXISTS", "foo"))
	if exists {
		ctx.Log.Info("foo 存在")
	} else {
		_, _ = client.Do("SET", "foo", "value")
		ctx.Log.Info("foo 不存在")

	}
	ctx.Log.Info("redis-go_key 的值:", res)
	return it.next.Handle(ctx)
}
