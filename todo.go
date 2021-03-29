package main

type ToDo struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
