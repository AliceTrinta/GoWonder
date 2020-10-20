package dao

import (
	"GoWonder/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

const (
	COLLECTION = "wonderland"
)

func (m *MongoDB) GetAll() ([]model.Wonderland, error) {
	var wonder []model.Wonderland
	err := db.C(COLLECTION).Find(bson.M{}).All(&wonder)
	return wonder, err
}

func (m *MongoDB) GetByID(id string) (model.Wonderland, error) {
	var wonder model.Wonderland
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&wonder)
	return wonder, err
}

func (m *MongoDB) Create(wonderland model.Wonderland) error {
	err := db.C(COLLECTION).Insert(&wonderland)
	return err
}

func (m *MongoDB) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *MongoDB) Update(id string, wonderland model.Wonderland) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &wonderland)
	return err
}
