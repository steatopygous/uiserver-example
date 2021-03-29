package main

import "fmt"

func(app App) toDoListPatch(c Context) {
	key, ok := c.Vars["id"]

	if ok {
		fmt.Fprintf(c.Response, "Getting resource %v", key)
	} else {
		fmt.Fprintf(c.Response, "There is no key called 'fred'!")
	}
}

