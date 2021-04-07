package main

import (
	"strconv"
)

// toDoListDelete() handles deletion of a todo item
func(app App) toDoListDelete() Handler {
	return func(c Context) {
		idString, ok := c.Vars["id"]

		if !ok {
			app.logger.Println("toDoListDelete() - missing id route parameter")
			app.Respond(c, nil, 400)
			return
		}

		id, err := strconv.Atoi(idString)

		if err != nil {
			app.logger.Println("toDoListDelete() - badly formatted todo ID =", idString)
			app.Respond(c, nil, 400)
			return
		}

		app.tdl.Delete(id)
		app.Respond(c, nil, 204)
	}
}

