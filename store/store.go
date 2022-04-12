package store

import (
	"ToDoApp/types"
)

type DB struct {
}

type dbInterface interface {
	InsertNewTodo(s *types.Todo) (err error)
	GetAllTodo() []types.Todo
	Delete(s *types.Todo, id int) (err error)
	GetOne(s *types.Todo, id int) (t types.Todo)
}

func New() *DB {
	return &DB{}
}

func NewMockDatabase(db *DB) dbInterface {
	return &DB{}
}

func (d DB) InsertNewTodo(s *types.Todo) error {

	types.ToDoList = append(types.ToDoList, *s)

	return nil
}

func (d DB) GetAllTodo() []types.Todo {

	return types.ToDoList
}

func (d DB) Delete(s *types.Todo, id int) error {

	types.ToDoList = append(types.ToDoList[:id-1], types.ToDoList[id:]...)

	return nil
}

func (d DB) GetOne(s *types.Todo, id int) types.Todo {

	return types.ToDoList[id-1]
}
