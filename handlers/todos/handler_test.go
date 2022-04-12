package todos

import (
	"ToDoApp/types"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewToDo(T *testing.T) {

	rr := httptest.NewRecorder()
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"210\"}")

	}
	body := []byte(`{"ID":1,"Description":"Testing1" , "Completed":false}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/createTodo", bytes.NewReader(body))

	if err != nil {

		T.Fatal(err)
	}
	handler(rr, req)
	status := rr.Code

	if status != http.StatusOK {
		T.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var t types.Todo
	err = t.DecodeFromJSON(rr.Body)

	if err != nil {
		T.Errorf("Error in decoding body")
	}
	var s td
	err = s.db.InsertNewTodo(&t)
	fmt.Println(t.ID)

	if err != nil {
		T.Errorf("error in Insert to DB")
	}
}

func TestDeleteToDo(T *testing.T) {

	rr := httptest.NewRecorder()
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"210\"}")

	}
	body := []byte(`{"ID":1,"Description":"Testing1" , "Completed":false}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/createTodo", bytes.NewReader(body))

	if err != nil {

		T.Fatal(err)
	}
	handler(rr, req)
	status := rr.Code

	if status != http.StatusOK {
		T.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
