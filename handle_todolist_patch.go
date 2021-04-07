package main

import (
	"strconv"
)

// toDoListPatch() updates the done status of the specified item, if it exists.
func(app App) toDoListPatch() Handler {
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

		tdl := app.tdl

		item, ok := tdl.Items[id]

		if !ok {
			app.logger.Println("toDoListPatch() - no item exists with requested ID", idString)
			app.Respond(c, nil, 404)
			return
		}

		// Make a copy of the item and decode the request over the top of it, which
		// will update any fields provided, but leave everything else alone.

		newItem := *item

		err = app.Decode(c.Request, &newItem)

		if err != nil {
			app.logger.Println("toDoListPatch() - Failed to decode the request JSON")
			app.Respond(c, nil, 400)
		}

		// In case the caller attempts to change the ID

		newItem.Id = item.Id

		// Replace the existing item with the updated one and save the list

		tdl.Items[id] = &newItem
		tdl.Save()

		app.Respond(c, nil, 204)
	}
}

