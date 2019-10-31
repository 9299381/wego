package redis

import (
	"fmt"
	"github.com/9299381/wego/configs"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

/**
MaxActive 最大连接数，即最多的tcp连接数，一般建议往大的配置，
但不要超过操作系统文件句柄个数（centos下可以ulimit -n查看）。
MaxIdle 最大空闲连接数，即会有这么多个连接提前等待着，但过了超时时间也会关闭。
IdleTimeout 空闲连接超时时间，
但应该设置比redis服务器超时时间短。否则服务端超时了，客户端保持着连接也没用。
Wait 这是个很有用的配置。如果超过最大连接，是报错，还是等待
*/

var pool *redis.Pool
var once sync.Once

func GetRedisPool() *redis.Pool {
	once.Do(func() {
		pool = newRedisPool()
	})
	return pool
}

func newRedisPool() *redis.Pool {
	conf := configs.LoadRedisConfig()
	timeout := conf.IdleTimeout
	pool = &redis.Pool{
		MaxActive:   conf.MaxActive,
		MaxIdle:     conf.MaxIdle,
		IdleTimeout: timeout,
		Wait:        true,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial(
				"tcp",
				conf.Uri,
				redis.DialPassword(conf.Auth),
				redis.DialDatabase(conf.Db),
				redis.DialConnectTimeout(timeout),
				redis.DialReadTimeout(timeout),
				redis.DialWriteTimeout(timeout),
			)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			return
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return pool

}
