package main

import (
	"github.com/steatopygous/uiserver"
)

type Handler = uiserver.Handler

// createRoutes() adds all of the routes + REST verb combinations to the server
func(app App) createRoutes(server *uiserver.UIServer) {
	server.Get("/api/todos", app.toDoListGetAll())

	server.Post("/api/todos/purge", app.toDoListPurgePost())
	server.Post("/api/todos", app.toDoListPost())

	server.Get("/api/todos/{id}", app.toDoListGetItem())
	server.Patch("/api/todos/{id}", app.toDoListPatch())
	server.Delete("/api/todos/{id}", app.toDoListDelete())

	server.Patch("/api/preferences", app.preferencesPatch())
}


