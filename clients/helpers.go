package clients

import (
	"github.com/9299381/wego/clients/mysql"
	"github.com/9299381/wego/clients/redis"
	redigo "github.com/gomodule/redigo/redis"
	"xorm.io/xorm"
)

func DB() *xorm.Engine {
	return mysql.GetDB()
}

func Redis() redigo.Conn {
	return redis.GetRedisPool().Get()
}

func RedisPool() *redigo.Pool {
	return redis.GetRedisPool()
}
