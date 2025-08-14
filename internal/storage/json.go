package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AaronBrownDev/go-todo-cli/internal/todo"
)

var todoList []todo.Todo
var jsonPath string

func init() {
	// Grabs home directory path. If fails, prints the error then exits program with code 1.
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Adds the designated json file path together with OS specific separators
	jsonPath = filepath.Join(homePath, ".todo", "data.json")

	// If the file does not exist, creates the path and data.json file
	if _, err = os.ReadFile(jsonPath); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(homePath, ".todo"), 0700)
		os.WriteFile(jsonPath, []byte("[]\n"), 0600)
	}

	// Reads the json file and returns the content as a slice of bytes. If fails, prints the error then exits program with code 1.
	content, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmarshals the slice of bytes into a slice of todo.Todo. If fails, prints the error then exits program with code 1.
	if err = json.Unmarshal(content, &todoList); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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
