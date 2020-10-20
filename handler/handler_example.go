package handler

import (
	"GoWonder/conf/mongo/dao"
	"GoWonder/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var con = dao.MongoDB{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

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

func GetAll(w http.ResponseWriter, r *http.Request) {
	wonder, err := con.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, wonder)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wonder, err := con.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	respondWithJson(w, http.StatusOK, wonder)
}

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

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := con.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}