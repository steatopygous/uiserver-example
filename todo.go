package main

type ToDo struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func (todo ToDo) MarkAsDone() {
	todo.Done = true
}

func (todo ToDo) MarkAsPending() {
	todo.Done = false
}
