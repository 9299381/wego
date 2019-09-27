package configs

import (
	"strconv"
	"time"
)

type RedisConfig struct {
	Uri         string
	Auth        string
	Db          int
	MaxActive   int
	MaxIdle     int
	IdleTimeout time.Duration
}

func (it *RedisConfig) Load() *RedisConfig {

	db, _ := strconv.Atoi(Env("REDIS_DB", "0"))
	maxActive, _ := strconv.Atoi(Env("REDIS_MAX_ACTIVE", "50"))
	maxIdle, _ := strconv.Atoi(Env("REDIS_MAX_IDLE", "5"))
	timeout, _ := strconv.Atoi(Env("REDIS_TIMEOUT", "10"))
	config := &RedisConfig{
		Uri:         Env("REDIS_URI", "127.0.0.1:6937"),
		Auth:        Env("REDIS_AUTH", "password"),
		Db:          db,
		MaxActive:   maxActive,
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(timeout) * time.Second,
	}
	return config
}
