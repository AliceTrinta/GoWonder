package main

import (
	"GoWonder/conf/mongo/dao"
	"GoWonder/controller"
	"fmt"
	"log"
	"net/http"
	"time"
)

//The Init function will create the stabilish a connection with the database.
func init() {

	//Connect to your MongoDatabase
	dao.MongoConnect()
	fmt.Println("Mongo connected")
}

func main() {
	handler := controller.Begin()
	svr := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      handler,
	}
	log.Fatal(svr.ListenAndServe())
}
