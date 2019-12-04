package mongo

import (
	"github.com/9299381/wego/configs"
	"gopkg.in/mgo.v2"
	"log"
)

var session *mgo.Session

func Session() *mgo.Session {
	if session == nil {
		var err error
		conf := configs.LoadMongoConfig()
		info := &mgo.DialInfo{
			Addrs:     conf.Address,
			Timeout:   conf.MaxIdleTime,
			Username:  conf.Username,
			Password:  conf.Password,
			Database:  conf.Database,
			PoolLimit: conf.MinPoolSize,
		}
		session, err = mgo.DialWithInfo(info)
		if err != nil {
			log.Fatalf("MongoCreateSession: %s\n", err)
		}
		session.SetMode(mgo.Monotonic, true)
	}
	return session.Clone()
}
func Coll(collection string, f func(*mgo.Collection)) {
	Table(configs.LoadMongoConfig().Database, collection, f)
}
func Table(database string, collection string, f func(*mgo.Collection)) {
	session := Session()
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			log.Fatalf("MongoSessionError: %v", err)
		}
	}()
	c := session.DB(database).C(collection)
	f(c)
}
