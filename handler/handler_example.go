package handler

import (
	"GoWonder/conf/mongo/dao"
	"GoWonder/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//Here, we have an example of a handler to deal with the http requests.

//The con variable is an instance of a databse connection.
var con = dao.MongoDB{}

//Some way to deal with errors.
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//This function will deal with a request to create an object in the database.
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var wonder model.Wonderland
	if err := json.NewDecoder(r.Body).Decode(&wonder); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	wonder.ID = bson.NewObjectId()
	if err := con.Create(wonder); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, wonder)
}

//This function will deal with a request to get all the objects stores in the database.
func GetAll(w http.ResponseWriter, r *http.Request) {
	wonder, err := con.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, wonder)
}

//This function will deal with a request to get an specific object at the dabase.
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wonder, err := con.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	respondWithJson(w, http.StatusOK, wonder)
}

//This function will deal with a request to do a update at an object at the database.
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var wonder model.Wonderland
	if err := json.NewDecoder(r.Body).Decode(&wonder); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := con.Update(params["id"], wonder); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": " success!"})
}

//This function will deal with a request to delet an object at the database.
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := con.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
