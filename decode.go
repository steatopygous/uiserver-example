package main

import (
	"encoding/json"
	"net/http"
)

// Decode() decodes the body of the request into the provided interface
func(app App) Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}
