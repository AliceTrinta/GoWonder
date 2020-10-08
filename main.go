package main

import (
	"GoWonder/controller"
	"log"
	"net/http"
	"time"
)

func main()  {
	handler := controller.Begin()
	svr := &http.Server{
		Addr: "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
		Handler: handler,
	}
	log.Fatal(svr.ListenAndServe())
}