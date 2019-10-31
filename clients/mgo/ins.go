package mgo

import (
	"context"
	"github.com/9299381/wego/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

func GetMgo(database ...string) *mongo.Database {
	config := configs.LoadMongoConfig()
	var client *mongo.Client
	want, err := readpref.New(readpref.SecondaryMode) //表示只使用辅助节点
	if err != nil {
		panic(err)
	}
	wc := writeconcern.New(writeconcern.WMajority())
	readconcern.Majority()
	//链接mongo服务
	opt := options.Client().ApplyURI(config.Uri)
	opt.SetLocalThreshold(3 * time.Second)     //只使用与mongo操作耗时小于3秒的
	opt.SetMaxConnIdleTime(5 * time.Second)    //指定连接可以保持空闲的最大毫秒数
	opt.SetMaxPoolSize(200)                    //使用最大的连接数
	opt.SetReadPreference(want)                //表示只使用辅助节点
	opt.SetReadConcern(readconcern.Majority()) //指定查询应返回实例的最新数据确认为，已写入副本集中的大多数成员
	opt.SetWriteConcern(wc)                    //请求确认写操作传播到大多数mongod实例
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if client, err = mongo.Connect(ctx, opt); err != nil {
		panic(err)
	}
	//UseSession(client)
	//判断服务是否可用
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	var db *mongo.Database
	if database == nil {
		db = client.Database(config.Database)
	} else {
		db = client.Database(database[0])

	}
	return db
}
