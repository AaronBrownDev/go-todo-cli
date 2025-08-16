package storage

import "github.com/AaronBrownDev/go-todo-cli/internal/todo"

type FilterFunc func(todo *todo.Todo) bool

type Storage interface {
	Save(todo *todo.Todo) error
	Get(id string) (*todo.Todo, error)
	List(filters ...FilterFunc) ([]*todo.Todo, error)
	Update(id string, todo *todo.Todo) error
	Delete(id string) error
}
