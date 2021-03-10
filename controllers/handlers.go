package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/CalebEWheeler/go-project-v2/database"
	"github.com/gorilla/mux"
)

type Person struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var tblName = "person"
var err error

func getPeople(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var people []Person

	result, err := database.Connect().Query("SELECT * FROM " + tblName)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var person Person
		err := result.Scan(&person.ID, &person.Name, &person.Age, &person.CreatedAt, &person.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		people = append(people, person)
	}
	json.NewEncoder(res).Encode(people)
}

func getPerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	fmt.Printf("getPerson() URL: %v", req.URL)

	result, err := database.Connect().Query("SELECT * FROM " + tblName + " WHERE id=" + params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var person Person

	for result.Next() {
		err := result.Scan(&person.ID, &person.Name, &person.Age, &person.CreatedAt, &person.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(res).Encode(person)
}

func createPerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	stmt, err := database.Connect().Prepare("INSERT INTO " + tblName + "(name, age) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	age := keyVal["age"]

	_, err = stmt.Exec(name, age)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(res, "New person was created")
}

func updatePerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	stmt, err := database.Connect().Prepare("UPDATE " + tblName + " SET name = ?, age = ? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newName := keyVal["name"]
	newAge := keyVal["age"]

	_, err = stmt.Exec(newName, newAge, params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(res, "Person with ID = %s was updated", params["id"])
}

func deletePerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	stmt, err := database.Connect().Prepare("DELETE FROM " + tblName + " WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(res, "Person with ID = %s was deleted", params["id"])
}
