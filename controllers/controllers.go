package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var prepend = "/api/v1/"

func SetupRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(prepend+"/", getPeople).Methods("GET")
	router.HandleFunc(prepend+"/person/{id}", getPerson).Methods("GET")
	router.HandleFunc(prepend+"/person", createPerson).Methods("POST")
	router.HandleFunc(prepend+"/person/{id}", updatePerson).Methods("PUT")
	router.HandleFunc(prepend+"/person/{id}", deletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
