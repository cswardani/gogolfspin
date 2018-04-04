package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	
	"config"
	"dao"
	"model"
)

var config = Config{}
var dao = MoviesDAO{}

// AllFieldsEndPoint endpoint
func AllFieldsEndPoint(w http.ResponseWriter, r *http.Request) {
	fields, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, fields)
}

// FindFieldEndpoint endpoint
func FindFieldEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fields, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Field ID")
		return
	}
	respondWithJson(w, http.StatusOK, fields)
}

// CreateFieldEndPoint endpoint
func CreateFieldEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var field Field
	if err := json.NewDecoder(r.Body).Decode(&field); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	field.ID = bson.NewObjectId()
	if err := dao.Insert(field); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, field)
}

// UpdateFieldEndPoint endpoint
func UpdateFieldEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var field Field
	if err := json.NewDecoder(r.Body).Decode(&field); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(field); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteFieldEndPoint endpoint
func DeleteFieldEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var field Field
	if err := json.NewDecoder(r.Body).Decode(&field); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(field); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fields", AllFieldsEndPoint).Methods("GET")
	r.HandleFunc("/fields", CreateFieldEndPoint).Methods("POST")
	r.HandleFunc("/fields", UpdateFieldEndPoint).Methods("PUT")
	r.HandleFunc("/fields", DeleteFieldEndPoint).Methods("DELETE")
	r.HandleFunc("/fields/{id}", FindFieldEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
