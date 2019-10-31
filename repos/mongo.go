package repos

import (
	"github.com/9299381/wego/clients/mgo"
	"github.com/9299381/wego/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

////////// mongodb 操作
func Mongo(database ...string) *mongo.Database {
	config := configs.LoadMongoConfig()
	if database == nil {
		return mgo.GetMongo().Database(config.Database)

	} else {
		return mgo.GetMongo().Database(database[0])
	}
}
