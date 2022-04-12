package router

import (
	"ToDoApp/handlers/todos"
	"ToDoApp/store"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *store.DB) *mux.Router {
	r := mux.NewRouter()

	todoH := todos.New(db)
	r.HandleFunc("/createTodo", todoH.NewToDo).Methods(http.MethodPost)
	r.HandleFunc("/deleteTodo", todoH.DeleteToDo).Methods(http.MethodDelete)
	r.HandleFunc("/getTodo", todoH.GetToDo).Methods(http.MethodPost)
	return r
}
