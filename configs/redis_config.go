package configs

import (
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

func LoadRedisConfig() *RedisConfig {
	config := &RedisConfig{
		Uri:         EnvString("redis.uri", "127.0.0.1:6937"),
		Auth:        EnvString("redis.auth", "password"),
		Db:          EnvInt("redis.db", 0),
		MaxActive:   EnvInt("redis.max_active", 50),
		MaxIdle:     EnvInt("redis.max_idle", 5),
		IdleTimeout: time.Duration(EnvInt("redis.timeout", 10)) * time.Second,
	}
	return config
}
