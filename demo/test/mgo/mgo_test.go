package mgo

import (
	"context"
	"fmt"
	"github.com/9299381/wego/repos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"testing"
	"time"
)

func getCollection() *mongo.Collection {
	return repos.Mongo().Collection("numbers")
}

func TestInsert(t *testing.T) {
	collection := getCollection()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	res, _ := collection.InsertOne(ctx, bson.M{"name": "cos", "value": 3.14159})
	id := res.InsertedID
	fmt.Println(id)

}

func TestFind(t *testing.T) {
	collection := getCollection()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.M{"name": "cos"})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
