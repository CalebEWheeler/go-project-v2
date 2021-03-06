package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// var prepend = "/api/v1"

func TestGetPeople(t *testing.T) {
	req, err := http.NewRequest("GET", prepend+"/person", nil)
	if err != nil {
		t.Fatal(err)
	}
	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(getPeople)
	handler.ServeHTTP(respRec, req)
	if status := respRec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//Check if the response body is what we expect.
	expected := `[{"id":1,"name":"Caleb","age":26,"created_at":"2021-03-04T15:05:12Z","updated_at":"2021-03-04T15:16:37Z"},{"id":2,"name":"Brad","age":25,"created_at":"2021-03-04T15:16:05Z","updated_at":"2021-03-04T15:16:05Z"}]`
	if strings.TrimRight(respRec.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", respRec.Body.String(), expected)
	}
}
