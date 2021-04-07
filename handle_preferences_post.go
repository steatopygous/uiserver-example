package main

// preferencesPatch() handles updates to the preferences
func(app App) preferencesPatch() Handler {
	type request struct {
		Width int `json:"width"`
		Height int `json:"height"`
	}

	return func(c Context) {
		r := request{}

		err := app.Decode(c.Request, &r)

		if err != nil {
			app.logger.Println("preferencesPatch() - failed to decode the body JSON - err =", err)
			app.Respond(c, nil, 400)
			return
		}

		app.preferences.setWindowSize(WindowSize{r.Width, r.Height})

		app.Respond(c, nil, 204)
	}
}
