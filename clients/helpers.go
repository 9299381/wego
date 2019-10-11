package clients

import (
	"github.com/9299381/wego/clients/mysql"
	"github.com/9299381/wego/clients/redis"
	"github.com/go-xorm/xorm"
	redigo "github.com/gomodule/redigo/redis"
)

func DB() *xorm.Engine {
	return mysql.GetDB()
}

func Redis() redigo.Conn {
	return redis.GetRedis()
}

func RedisPool() *redigo.Pool {
	return redis.GetRedisPool()
}
