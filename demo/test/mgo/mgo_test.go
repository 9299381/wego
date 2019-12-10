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
	mongo.Col(PEOPLE, func(c *mgo.Collection) {
		_ = c.Insert(person)
	})
}
func TestMgoFind(t *testing.T) {
	var persons []Person
	filter := bson.M{}
	mongo.Col(PEOPLE, func(c *mgo.Collection) {
		_ = c.Find(filter).All(&persons)
	})
	fmt.Print(persons)
}
func TestMgoFineOne(t *testing.T) {
	var p0 = &Person{}
	var p1 = &Person{}
	var p2 = &Person{}
	id := bson.ObjectIdHex("5de71ec1d4a40398def8c0df")
	filter1 := bson.D{{"_id", id}}
	filter2 := bson.M{"_id": id}
	mongo.Col(PEOPLE, func(c *mgo.Collection) {
		_ = c.FindId(id).One(p0)
		_ = c.Find(filter1).One(p1)
		_ = c.Find(filter2).One(p2)

	})
	fmt.Println(p0.Name)
	fmt.Println(p1.Name)
	fmt.Println(p2.Name)

}
func TestMgoUpdate(t *testing.T) {
	filter := bson.M{
		"_id": bson.ObjectIdHex("5de71ec1d4a40398def8c0df")}
	update := bson.M{"$set": bson.M{"name": "中文"}}
	mongo.Col(PEOPLE, func(c *mgo.Collection) {
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
func TestCount(t *testing.T) {
	filter := bson.M{}
	var count int
	mongo.Col(PEOPLE, func(c *mgo.Collection) {
		count, _ = c.Find(filter).Count()
	})
	fmt.Println(count)
}
