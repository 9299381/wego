package configs

type MongoConfig struct {
	Uri      string
	Database string
	MaxPool  int
}

func LoadMongoConfig() *MongoConfig {
	config := &MongoConfig{
		Uri:      EnvString("mongo.uri", "mongodb://127.0.0.1:27017"),
		Database: EnvString("mongo.database", "base"),
		MaxPool:  EnvInt("mongo.max_pool", 200),
	}

	return config
}
