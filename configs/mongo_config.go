package configs

import "time"

type MongoConfig struct {
	Address     []string
	Database    string
	Username    string
	Password    string
	MaxIdleTime time.Duration
	MinPoolSize int
}

/**
每次有连接被重置到空闲池时，打一个时间戳。
在轮询 goroutine 中每隔一段时间 review 空闲连接的空闲时长，
当时长大于maxIdleTimeMS时，就释放连接，将空闲池的大小控制在minPoolSize。
若maxIdleTimeMS不设置或为 0，则默认为不进行释放
*/
func LoadMongoConfig() *MongoConfig {
	idle := EnvInt("mongo.max_idle_time", 30)
	config := &MongoConfig{
		Address:     EnvStringSlice("mongo.uri", []string{"127.0.0.1:27017"}),
		Database:    EnvString("mongo.database", "base"),
		Username:    EnvString("mongo.username", ""),
		Password:    EnvString("mongo.password", ""),
		MinPoolSize: EnvInt("mongo.min_pool_size", 10),
		MaxIdleTime: time.Duration(idle) * time.Second,
	}
	return config
}
