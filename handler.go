package main

import "github.com/steatopygous/uiserver"

type Handler = uiserver.Handler
type HandlerFunc func(context Context)

func(hf HandlerFunc) ServeHTTP(context Context) {
	hf(context)
}