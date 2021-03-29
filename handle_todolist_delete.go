package main

import (
	"strconv"
)

func(app App) toDoListDelete(c Context) {
	idString, ok := c.Vars["id"]

	if ok {
		id, err := strconv.Atoi(idString)

		if err != nil {
			// Return an error

			app.logger.Println("toDoListDelete() - badly formatted todo ID =", idString)
		}

		app.tdl.Delete(id)

		// Return 201
	} else {
		// Return an error

		app.logger.Println("toDoListDelete() - missing id route parameter")
	}
}
