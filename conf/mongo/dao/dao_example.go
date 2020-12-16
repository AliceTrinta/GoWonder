package dao

import (
	"GoWonder/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//An example of a CRUD for mongoDB.

//The Db variable will establish one or more connections with the cluster of servers defined by the url parameter.
var db *mgo.Database

//Constant that represents the collection at the database.
const (
	COLLECTION = "wonderland"
)

//Function to get all objects stored at the collection.
func (m *MongoDB) GetAll() ([]model.Wonderland, error) {
	var wonder []model.Wonderland
	err := db.C(COLLECTION).Find(bson.M{}).All(&wonder)
	return wonder, err
}

//Function to get an object specified by the ID.
func (m *MongoDB) GetByID(id string) (model.Wonderland, error) {
	var wonder model.Wonderland
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&wonder)
	return wonder, err
}

//Function to create an object in the collection.
func (m *MongoDB) Create(wonderland model.Wonderland) error {
	err := db.C(COLLECTION).Insert(&wonderland)
	return err
}

//Function to delete some object in the collection.
func (m *MongoDB) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

//Function to update some object in the collection.
func (m *MongoDB) Update(id string, wonderland model.Wonderland) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &wonderland)
	return err
}
