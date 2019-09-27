package clients

import (
	"github.com/9299381/wego/clients/mysql_client"
	"github.com/9299381/wego/clients/redis_client"
	"github.com/go-xorm/xorm"
	"github.com/gomodule/redigo/redis"
)

func DB() *xorm.Engine {
	return mysql_client.Get()
}

func Redis() redis.Conn {
	return redis_client.Get()
}

// 为统一php模式而封装
// micro -> service,  service ->route
func Micro(micro string) *microService {
	return &microService{
		micro:  micro,
		params: map[string]interface{}{},
	}
}
