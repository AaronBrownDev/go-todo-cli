package storage

import "github.com/AaronBrownDev/go-todo-cli/internal/todo"

type Storage interface {
	Save(todo *todo.Todo) error
	Get(id string) (*todo.Todo, error)
	// TODO: List
	Update(id string, todo *todo.Todo) error
	Delete(id string) error
}
