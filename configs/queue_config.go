package configs

import (
	"strconv"
	"strings"
	"time"
)

type QueueConfig struct {
	Prefix      string
	Listen      []string
	Interval    time.Duration
	Concurrency int
}

func (it *QueueConfig) Load() *QueueConfig {
	listen := Env("QUEUE_LISTEN", "queue")
	interal, _ := strconv.Atoi(Env("QUEUE_INTERVAL", "1"))
	concurrency, _ := strconv.Atoi(Env("QUEUE_CONCURRENCY", "1"))
	config := &QueueConfig{
		Prefix:      Env("QUEUE_PREFIX", "wego"),
		Listen:      strings.Split(listen, ","),
		Interval:    time.Duration(interal) * time.Second,
		Concurrency: concurrency,
	}
	return config
}
