package main

import (
	"fmt"
)

func(app App) toDoListGetAll(c Context) {
	fmt.Println("getResource() - route =", c.Route, "method =", c.Method, "path =", c.Path)
}

func(app App) toDoListGetItem(c Context) {
	fmt.Println("getResource() - route =", c.Route, "method =", c.Method, "path =", c.Path)

	key, ok := c.Vars["id"]

	if ok {
		fmt.Fprintf(c.Response, "Getting resource %v", key)
	} else {
		fmt.Fprintf(c.Response, "There is no key called 'fred'!")
	}
}
