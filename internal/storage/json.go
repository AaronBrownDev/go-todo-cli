package storage

import (
	"github.com/AaronBrownDev/go-todo-cli/internal/todo"
)

type jsonRepository struct {
	// TODO
}

func NewJsonRepository() Storage {
	return &jsonRepository{}
}

// Save implements Storage.
func (j *jsonRepository) Save(todo *todo.Todo) error {
	panic("unimplemented")
}

// Get implements Storage.
func (j *jsonRepository) Get(id string) (*todo.Todo, error) {
	panic("unimplemented")
}

// Update implements Storage.
func (j *jsonRepository) Update(id string, todo *todo.Todo) error {
	panic("unimplemented")
}

// Delete implements Storage.
func (j *jsonRepository) Delete(id string) error {
	panic("unimplemented")
}
