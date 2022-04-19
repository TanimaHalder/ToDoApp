package store

import (
	"ToDoApp/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertNewTodo(t *testing.T) {

	repo := &DB{}
	obj := types.Todo{
		ID:          1,
		Description: "test2",
		Completed:   false,
	}
	err := repo.InsertNewTodo(&obj)

	assert.NoError(t, err)
}
