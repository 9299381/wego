package configs

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"strconv"
	"strings"
	"time"
)

type QueueConfig struct {
	Config
	Prefix string
	Listen []string
	Interval time.Duration
	Concurrency    int
}

func (it *QueueConfig) Load() contracts.Iconfig {
	listen := wego.Env("QUEUE_LISTEN","queue")
	interal,_:= strconv.Atoi(wego.Env("QUEUE_INTERVAL","1"))
	concurrency,_ :=strconv.Atoi(wego.Env("QUEUE_CONCURRENCY","1"))
	config := &QueueConfig{
		Prefix:      wego.Env("QUEUE_PREFIX","wego"),
		Listen:      strings.Split(listen, ","),
		Interval:    time.Duration(interal)*time.Second,
		Concurrency: concurrency,
	}
	return config
}

func (it *QueueConfig)Get(key string) string {
	return it.GetKey(it,key)
}