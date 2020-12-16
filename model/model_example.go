package model

import "gopkg.in/mgo.v2/bson"

//Examples of models.

type Wonderland struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Characters []Character   `bson:"characters" json:"characters"`
}

type Character struct {
	Name string `bson:"name" json:"name"`
}
