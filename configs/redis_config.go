package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"strconv"
	"time"
)

type RedisConfig struct {
	Config
	Uri         string
	Auth        string
	Db          int
	MaxActive   int
	MaxIdle     int
	IdleTimeout time.Duration
}

func (it *RedisConfig) Load() contracts.Iconfig {

	db, _ := strconv.Atoi(wego.Env("REDIS_DB", "0"))
	maxActive, _ := strconv.Atoi(wego.Env("REDIS_MAX_ACTIVE", "50"))
	maxIdle, _ := strconv.Atoi(wego.Env("REDIS_MAX_IDLE", "5"))
	timeout, _ := strconv.Atoi(wego.Env("REDIS_TIMEOUT", "10"))
	config := &RedisConfig{
		Uri:         wego.Env("REDIS_URI", "127.0.0.1:6937"),
		Auth:        wego.Env("REDIS_AUTH", "password"),
		Db:          db,
		MaxActive:   maxActive,
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(timeout) * time.Second,
	}
	return config
}

func (it *RedisConfig) Get(key string) string {
	return it.GetKey(it, key)
}
