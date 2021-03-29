package main

import "encoding/json"

func(app App) Respond(c Context, data interface{}, status int){
	c.Writer.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(c.Writer).Encode(data)

		if err != nil {
			app.logger.Println("Respond() - encoding of data failed - %v", data)
			return
		}
	}
}
