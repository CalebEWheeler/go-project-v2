package controllers

import (
	"fmt"
	"net/http"

	"github.com/CalebEWheeler/go-project-v2/database"
)

type Person struct {
	ID        string `json:"person_id"`
	Name      string `json:"person_name"`
	Age       string `json:"person_age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var err error

func getPeople(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Reached all people")
	w.Header().Set("Content-Type", "application/json")

	// var people []Person

	result, err := database.Connect().Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "you've reached me.")
	defer result.Close()

	// for result.Next() {
	// 	var person Person
	// 	err := result.Scan(person.ID, person.Name, person.Age)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	people = append(people, person)
	// }

	// json.NewEncoder(w).Encode(people)
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
