package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type ToDoList struct {
  Path string
  NextId int
  Items map[int]*ToDo
}

func NewToDoList(path string) ToDoList {
  var items map[int]*ToDo

  tdl := ToDoList{path, 0, items}

  tdl.Load()

  return tdl
}

func(tdl ToDoList) Load() {
  if _, err := os.Stat(tdl.Path); os.IsNotExist(err) {
    // The file doesn't exist yet.  Create it and leave the list empty.

    tdl.Save()
  } else {
    file, err := os.Open(tdl.Path)

    if err != nil {
      panic(fmt.Sprintf("failed to open todo list file %s - %v", tdl.Path, err))
    }

    decoder := json.NewDecoder(file)

    if err := decoder.Decode(&tdl); err != nil {
      panic(fmt.Sprintf("failed to decode JSON for ToDoList"))
    }

    file.Close()
  }
}

func(tdl ToDoList) Save() {
  file, _ := os.Create(tdl.Path)

  encoder := json.NewEncoder(file)

  if err := encoder.Encode(&tdl); err != nil {
    panic(fmt.Sprintf("failed to encode JSON for ToDoList"))
  }

  file.Close()
}

func(tdl ToDoList) Add(text string) {
  id := tdl.NextId

  tdl.NextId++

  todo := &ToDo{id, text, false}

  tdl.Items[id] = todo
}

func(tdl ToDoList) Delete(id int) {
  delete(tdl.Items, id)
}

