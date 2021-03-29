package main

import (
	"encoding/json"
	"net/http"
)

func(app App) Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
