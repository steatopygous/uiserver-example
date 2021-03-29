package main

import (
	"fmt"
	"strconv"
)

func(app App) toDoListGetAll() Handler {
	return func(c Context) {
		app.Respond(c, app.tdl.Items, 200)
	}
}

func(app App) toDoListGetItem() Handler {
	return func(c Context) {
		idString, ok := c.Vars["id"]

		if !ok {
			app.logger.Println(fmt.Sprintf("toDoListGetItem() - no 'id' field was provided in the URL"))
			app.Respond(c, nil, 400)
			return
		}

		id, err := strconv.Atoi(idString)

		if err != nil {
			app.logger.Println(fmt.Sprintf("toDoListGetItem() - the 'id' field provided was not an integer"))
			app.Respond(c, nil, 400)
			return
		}

		item, ok := app.tdl.Items[id]

		if ok {
			app.Respond(c, item, 200)
		} else {
			app.Respond(c, nil, 404)
		}
	}
}
