package handler

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Welcome to Wonder")
	w.Write(msg)
}
