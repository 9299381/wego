package mgo

import (
	"fmt"
	"github.com/9299381/wego/clients/mongo"
	"github.com/9299381/wego/tools"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type Person struct {
	Id    bson.ObjectId `bson:"_id"`
	Name  string        `json:"name"`
	Phone string        `json:"phone"`
}

const (
	PEOPLE string = "people"
)

func TestMgoInsert(t *testing.T) {
	person := Person{
		Id:    bson.NewObjectId(),
		Name:  tools.RandString(6, "a"),
		Phone: tools.RandString(10, "0"),
	}
	mongo.Coll(PEOPLE, func(c *mgo.Collection) {
		_ = c.Insert(person)
	})
}
func TestMgoFind(t *testing.T) {
	var persons []Person
	filter := bson.M{}
	mongo.Coll(PEOPLE, func(c *mgo.Collection) {
		_ = c.Find(filter).All(&persons)
	})
	fmt.Print(persons)
}
func TestMgoFineOne(t *testing.T) {
	var person = &Person{}
	filter := bson.M{"_id": bson.ObjectIdHex("5de71ec1d4a40398def8c0df")}
	mongo.Coll(PEOPLE, func(c *mgo.Collection) {
		_ = c.Find(filter).One(person)
	})
	fmt.Print(person.Name)
}
func TestMgoUpdate(t *testing.T) {
	filter := bson.M{
		"_id": bson.ObjectIdHex("5de71ec1d4a40398def8c0df")}
	update := bson.M{"$set": bson.M{"name": "中文"}}
	mongo.Coll(PEOPLE, func(c *mgo.Collection) {
		_ = c.Update(filter, update)
	})
}
func TestDb(t *testing.T) {
	person := Person{
		Id:    bson.NewObjectId(),
		Name:  tools.RandString(6, "a"),
		Phone: tools.RandString(10, "0"),
	}
	mongo.Table("test", PEOPLE, func(c *mgo.Collection) {
		_ = c.Insert(person)
	})

	var persons []Person
	filter := bson.M{}
	mongo.Table("test", PEOPLE, func(c *mgo.Collection) {
		_ = c.Find(filter).All(&persons)
	})
	fmt.Print(persons)
}
