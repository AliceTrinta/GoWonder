package dao

import (
	"log"

	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
)

//The MongoDB will store the database information address.
type MongoDB struct {
	Server   string
	Database string
}

//The ReadMongoConfig function will read the configuration file that contains the database information address.
func readMongoConfig() (m *MongoDB) {
	if _, err := toml.DecodeFile("C:\\Users\\ATG20\\go\\src\\GoWonder\\conf\\mongo\\dao\\mongo.toml", &m); err != nil {
		log.Fatal(err)
	}
	return
}

//The MongoConnect function will establish the connection with the database.
func MongoConnect() {
	var m = readMongoConfig()
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
