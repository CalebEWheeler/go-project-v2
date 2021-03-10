package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var prepend = "/api/v1"

type PersonResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func TestGetPeople(t *testing.T) {
	req, err := http.NewRequest("GET", prepend+"/person", nil)
	if err != nil {
		t.Fatal(err)
	}
	respRec := httptest.NewRecorder()
	fmt.Println(req.URL)
	handler := http.HandlerFunc(getPeople)
	handler.ServeHTTP(respRec, req)
	if status := respRec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//Check if the response body is what we expect. Use strings.TrimRight(respRec, "\n") to match the formatting of respRec.Body.String() to expected.
	expected := `[{"id":1,"name":"Caleb","age":26,"created_at":"2021-03-04T15:05:12Z","updated_at":"2021-03-04T15:16:37Z"},{"id":2,"name":"Brad","age":25,"created_at":"2021-03-04T15:16:05Z","updated_at":"2021-03-04T15:16:05Z"}]`
	if strings.TrimRight(respRec.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", respRec.Body.String(), expected)
	}
}

func TestGetPerson(t *testing.T) {
	IDs := []string{"1", "2"}
	for _, id := range IDs {
		req, err := http.NewRequest("GET", prepend+"/person/"+id, nil)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("TestGetPerson(): %v\n", req.URL)
		respRec := httptest.NewRecorder()
		handler := http.HandlerFunc(getPerson)
		handler.ServeHTTP(respRec, req)

		if status := respRec.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if id == "1" {
			expected := `{"id":1,"name":"Caleb","age":26,"created_at":"2021-03-04T15:05:12Z","updated_at":"2021-03-04T15:16:37Z"}`
			if strings.TrimRight(respRec.Body.String(), "\n") != expected {
				t.Errorf("handler returned unexpected body: got %v want %v", respRec.Body.String(), expected)
			}
		}

		if id == "2" {
			expected := `{"id":2,"name":"Brad","age":25,"created_at":"2021-03-04T15:16:05Z","updated_at":"2021-03-04T15:16:05Z"}`
			if strings.TrimRight(respRec.Body.String(), "\n") != expected {
				t.Errorf("handler returned unexpected body: got %v want %v", respRec.Body.String(), expected)
			}
		}
	}
}

func TestCreatePerson(t *testing.T) {
	var jsonStr = []byte(`{"name":"Betsy","age":"29"}`)

	req, err := http.NewRequest("POST", prepend+"person", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(createPerson)
	handler.ServeHTTP(respRec, req)
	if status := respRec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := "New person was created"
	if respRec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", respRec.Body.String(), expected)
	}

}
