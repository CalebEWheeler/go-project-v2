package controllers

import (
	"log"
	"net/http"

	"github.com/CalebEWheeler/go-project-v2/database"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	pre := router.PathPrefix("/api/v1").Subrouter()
	pre.HandleFunc("/person", getPeople).Methods("GET")
	pre.HandleFunc("/person/{id}", getPerson).Methods("GET")
	pre.HandleFunc("/person", createPerson).Methods("POST")
	pre.HandleFunc("/person/{id}", updatePerson).Methods("PUT")
	pre.HandleFunc("/person/{id}", deletePerson).Methods("DELETE")
	// router.HandleFunc(prepend+"/person/1", getPerson).Methods("GET")

	database.InitDatabase()

	log.Fatal(http.ListenAndServe(":8080", router))
	return router
}
