package main

// toDoListPurgePost() handles purging of items that have been marked as done
func(app App) toDoListPurgePost() Handler {
	return func(c Context) {
		app.tdl.Purge()
		app.Respond(c, nil, 201)
	}
}

// toDoListPost() handles creation of new todo items
func(app App) toDoListPost() Handler {
	type request struct {
		text string
	}

	type response struct {
		id int
	}

	return func(c Context) {
		r := request{}

		err := app.Decode(c.Request, &r)

		if err != nil {
			app.logger.Println("toDoListPost() - failed to decode the body JSON - err =", err)
			app.Respond(c, nil, 400)
			return
		}

		id := app.tdl.Add(r.text)

		app.Respond(c, response{id}, 200)
	}
}
