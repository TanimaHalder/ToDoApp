package todos

import (
	"ToDoApp/store"
	"ToDoApp/types"
	"fmt"

	"net/http"
)

type td struct {
	db   *store.DB
	Todo interface {
		NewToDo(rw http.ResponseWriter, r *http.Request)
		DeleteToDo(rw http.ResponseWriter, r *http.Request)
		GetToDo(rw http.ResponseWriter, r *http.Request)
	}
}

func New(db *store.DB) *td {
	return &td{db: db}
}
func (s td) NewToDo(rw http.ResponseWriter, r *http.Request) {

	var t types.Todo
	if err := t.DecodeFromJSON(r.Body); err != nil {
		http.Error(rw, "Failed to decode", http.StatusBadRequest)
		return
	}

	//types.StudentsList = append(types.StudentsList, student)

	if err := s.db.InsertNewTodo(&t); err != nil {
		http.Error(rw, "Failed to insert", http.StatusBadRequest)
		return
	}

	fmt.Printf("%#v", s.db.GetAllTodo())
	rw.WriteHeader(http.StatusCreated)
	t.EncodeToJSON(rw)
	return

}

func (s td) DeleteToDo(rw http.ResponseWriter, r *http.Request) {

	var t types.Todo
	/*params := mux.Vars(r)
	id, err := strconv.Atoi(params["ID"])*/

	if err := t.DecodeFromJSON(r.Body); err != nil {
		http.Error(rw, "Failed to decode", http.StatusBadRequest)
		return
	}

	s.db.Delete(&t, int(t.ID))

	fmt.Printf("%#v", s.db.GetAllTodo())
	rw.WriteHeader(http.StatusOK)
	t.EncodeToJSON(rw)
	return

}

func (s td) GetToDo(rw http.ResponseWriter, r *http.Request) {

	var t types.Todo

	if err := t.DecodeFromJSON(r.Body); err != nil {
		http.Error(rw, "Failed to decode", http.StatusBadRequest)
		return
	}

	t = s.db.GetOne(&t, int(t.ID))

	fmt.Printf("%#v", s.db.GetOne(&t, int(t.ID)))
	rw.WriteHeader(http.StatusOK)
	t.EncodeToJSON(rw)
	return

}
