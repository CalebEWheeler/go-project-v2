package controllers

import (
	"fmt"
	"net/http"
)

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reached all people")
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got a person")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "created a person")
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updated a person")
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete a person")
}
