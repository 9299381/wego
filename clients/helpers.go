package clients

import (
	"github.com/9299381/wego/clients/mysql"
	"github.com/9299381/wego/clients/redis"
	"github.com/go-xorm/xorm"
	redigo "github.com/gomodule/redigo/redis"
)

func DB() *xorm.Engine {
	return mysql.Get()
}

func Redis() redigo.Conn {
	return redis.Get()
}

func RedisPool() *redigo.Pool {
	return redis.Pool
}

// 为统一php模式而封装
// micro -> service,  service ->route
func Micro(micro string) *microService {
	return &microService{
		micro:  micro,
		params: map[string]interface{}{},
	}
}
