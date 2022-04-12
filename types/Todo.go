package types

import (
	"encoding/json"
	"io"
)

type Todo struct {
	ID          int64  `json:"ID"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t *Todo) DecodeFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

func (t *Todo) EncodeToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(t)

}

var ToDoList = []Todo{
	{
		ID:          1,
		Description: "grocery shooping",
		Completed:   true,
	},
	{
		ID:          2,
		Description: "kitchen cleanning",
		Completed:   true,
	},
	{
		ID:          3,
		Description: "laundry cleaning",
		Completed:   false,
	},
}
