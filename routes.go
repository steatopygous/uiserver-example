package main

import (
	"log"

	"github.com/steatopygous/uiserver"
)

type App struct {
	tdl ToDoList
	logger *log.Logger
}

type Handler = uiserver.Handler

func createRoutes(server uiserver.UIServer, toDoListPath string, logger *log.Logger) {
	tdl := NewToDoList(toDoListPath)
	app := App{tdl, logger}

	server.Get("/api/todos", app.toDoListGetAll())

	//server.Post("/api/todos/purge", app.toDoListPurgePost())
	server.Post("/api/todos", app.toDoListPost())

	server.Get("/api/todos/{id}", app.toDoListGetItem())
	server.Patch("/api/todos/{id}", app.toDoListPatch())
	server.Delete("/api/todos/{id}", app.toDoListDelete())
}


