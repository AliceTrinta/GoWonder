package dao

import (
	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
	"log"
)

type MongoDB struct {
	Server string
	Database string
}

func readMongoConfig() (m *MongoDB) {
	if _, err := toml.DecodeFile("C:\\Users\\ATG20\\go\\src\\GoWonder\\conf\\mongo\\dao\\mongo.toml", &m); err != nil {
		log.Fatal(err)
	}
	return
}

func MongoConnect() {
	var m = readMongoConfig()
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
