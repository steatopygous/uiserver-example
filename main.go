package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/pkg/browser"
	"github.com/steatopygous/uiserver"
)

//go:embed ui/*
var ui embed.FS

func main() {
	server := uiserver.New(ui)
	path := os.Args[1]
	logger := createLogger()

	createRoutes(server, path, logger)

	//go openDefaultBrowser()
	go openChromeWindow()

	err := server.Run(":1234")

	if err != nil {
		fmt.Println("uiServer.Run() returned an error...", err)
	}

	fmt.Println("Hmmm... we shouldn't get to here!")
}

func createLogger() *log.Logger {
	file, err := os.OpenFile("logs.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "", log.Ldate | log.Ltime | log.Lshortfile)
}


type Context = uiserver.Context // An alias to make writing REST handlers a little more succinct


func openChromeWindow() {
	chrome := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	app := "--app=http://localhost:1234/"
	position := "--window-position=500,50"
	size := "--window-size=900,1200"

	time.Sleep(100 * time.Millisecond)

	exec.Command(chrome, app, position, size).Start()
}

func openDefaultBrowser() {
	time.Sleep(100 * time.Millisecond)

	browser.OpenURL("http://localhost:1234/")
}