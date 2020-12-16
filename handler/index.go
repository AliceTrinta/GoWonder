package handler

import "net/http"

//The index function will be called by the index page.
func Index(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Welcome to Wonder")
	w.Write(msg)
}
