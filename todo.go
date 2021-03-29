package main

type ToDo struct {
  Id int
  Text string
  Done bool
}

func(todo ToDo) MarkAsDone() {
  todo.Done = true
}

func(todo ToDo) MarkAsPending() {
  todo.Done = false
}


