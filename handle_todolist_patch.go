package main

import (
	"strconv"
)

// toDoListPatch() updates the done status of the specified item, if it exists.
func(app App) toDoListPatch() Handler {
	type request struct {
		done bool
	}

	return func(c Context) {
		idString, ok := c.Vars["id"]

		if !ok {
			app.logger.Println("toDoListPatch() - missing id route parameter")
			app.Respond(c, nil, 400)
			return
		}

		id, err := strconv.Atoi(idString)

		if err != nil {
			app.logger.Println("toDoListPatch() - badly formatted todo ID =", idString)
			app.Respond(c, nil, 400)
			return
		}

		item, ok := app.tdl.Items[id]

		if !ok {
			app.Respond(c, nil, 404)
			return
		}

		r := request{}

		err = app.Decode(c.Request, &r)

		if err != nil {
			app.logger.Println("toDoListPatch() - Failed to decode the request JSON")
			app.Respond(c, nil, 400)
		}

		item.Done = r.done

		app.Respond(c, nil, 201)
	}
}

