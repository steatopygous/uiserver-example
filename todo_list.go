package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// ToDoList holds the list of todo items, along with the next available ID and the path
// to the JSON file where the list is stored externally.
type ToDoList struct {
	path   string
	NextId int
	Items  map[int]*ToDo
}

// NewToDoList() creates a new ToDoList and loads its content from the JSON file, if it exists
func NewToDoList(path string) ToDoList {
	var items map[int]*ToDo

	tdl := ToDoList{path, 0, items}

	tdl.Load()

	return tdl
}

// Load() sets the list to the content of the JSON file
func (tdl ToDoList) Load() {
	if _, err := os.Stat(tdl.path); os.IsNotExist(err) {
		// The file doesn't exist yet.  Create it with the empty list.

		tdl.Save()
	} else {
		file, err := os.Open(tdl.path)

		if err != nil {
			panic(fmt.Sprintf("failed to open todo list file %s - %v", tdl.path, err))
		}

		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&tdl); err != nil {
			panic(fmt.Sprintf("failed to decode JSON for ToDoList"))
		}

		file.Close()
	}
}

// Save() dumps the current content of the list to the associated JSON file
func (tdl ToDoList) Save() {
	file, _ := os.Create(tdl.path)

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(&tdl); err != nil {
		panic(fmt.Sprintf("failed to encode JSON for ToDoList"))
	}

	file.Close()
}

// Add() creates a new todo item with the provided text and adds it to the list
func (tdl ToDoList) Add(text string) int {
	id := tdl.NextId

	tdl.NextId++

	todo := &ToDo{id, text, false}

	tdl.Items[id] = todo

	tdl.Save()

	return id
}

// Delete() removes the todo item with the given ID
func (tdl ToDoList) Delete(id int) error {
	_, ok := tdl.Items[id]

	if ok {
		delete(tdl.Items, id)
		tdl.Save()

		return nil
	} else {
		return errors.New(fmt.Sprintf("there is no item with ID %d", id))
	}
}

// Purge() removes all todo items that are marked as done
func (tdl ToDoList) Purge() {
	var done  []int

	for id, todo := range tdl.Items {
		if todo.Done {
			done = append(done, id)
		}
	}

	for _, id := range done {
		delete(tdl.Items, id)
	}

	tdl.Save()
}
