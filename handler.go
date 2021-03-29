package main

type Handler interface {
	ServerHTTP(context Context)
}

type HandlerFunc func(context Context)

func(hf HandlerFunc) ServeHTTP(context Context) {
	hf(context)
}