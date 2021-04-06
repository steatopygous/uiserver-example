package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/steatopygous/uiserver"
)

//go:embed ui/public/*
var ui embed.FS

type App struct {
	tdl         *ToDoList
	logger      *log.Logger
	preferences Preferences
}

var help = flag.Bool("help", false, "display available command-line flags")
var path = flag.String("json", "./todos.json", "path to the todos JSON file")
var port = flag.Int("port", 1234, "port the server will run on")
var display = flag.String("display", "tab", "tab | chrome")

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	todos()
}

// todos() starts the server and opens the UI
func todos() {
	server := uiserver.New(ui)

	tdl := NewToDoList(*path)

	logger := createLogger()
	preferences := Load()

	app := App{&tdl, logger, preferences}

	app.createRoutes(&server)

	openUI(preferences)

	listenOn := fmt.Sprintf(":%d", *port)

	err := server.Run(listenOn)

	if err != nil {
		fmt.Println("uiServer.Run() returned an error...", err)
	}

	fmt.Println("Hmmm... we shouldn't get to here!")
}

// createLogger() creates a simple logger for errors
func createLogger() *log.Logger {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

type Context = uiserver.Context // An alias to make writing REST handlers a little more succinct

