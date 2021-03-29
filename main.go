package main

import (
	"embed"
	"log"
	"os"
	"time"

	"github.com/steatopygous/uiserver"

	"github.com/pkg/browser"
)

//go:embed ui/*
var ui embed.FS

func main() {
	server := uiserver.New(ui)
	path := os.Args[1]
	logger := createLogger()

	createRoutes(server, path, logger)

	go openBrowser()

	server.Run(":1234")
}

func createLogger() *log.Logger {
	file, err := os.OpenFile("logs.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "", log.Ldate | log.Ltime | log.Lshortfile)
}


type Context = uiserver.Context // An alias to make writing REST handlers a little more succinct


func openBrowser() {
	time.Sleep(100 * time.Millisecond)

	browser.OpenURL("http://localhost:1234/")
}

