package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/pkg/browser"
)

// openUI() opens the preferred browser, pointing to the todo list server
func openUI(preferences Preferences) {
	switch *display {
	case "tab":
		go openDefaultBrowserTab()

	case "chrome":
		go openChromeWindow(preferences)

	default:
		go openDefaultBrowserTab()
	}
}

// openChromeWindow() attempts to open Chrome as an app, pointing to the todo list server
func openChromeWindow(preferences Preferences) {
	chrome := chromePath()
	app := fmt.Sprintf("--app=%s", serverUrl())

	position := "--window-position=200,100"
	size := fmt.Sprintf("--window-size=%d,%d", preferences.Size.Width, preferences.Size.Height)

	time.Sleep(100 * time.Millisecond)

	err := exec.Command(chrome, app, position, size).Run()

	if err != nil {
		fmt.Println("openChromeWindow() - Failed to open Chrome app window...", err)
		os.Exit(-1)
	}
}

// chromePath() returns the path to the Chrome executable, based on the OS
func chromePath() string {
	if runtime.GOOS == "windows" {
		return "chrome.exe"
	} else {
		return "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	}
}

// opeDefaultBrowserTab() attempts to open a tab in the default browser, pointing to the todo list server
func openDefaultBrowserTab() {
	time.Sleep(100 * time.Millisecond)

	err := browser.OpenURL(serverUrl())

	if err != nil {
		fmt.Println("openDefaultBrowserTab() - Failed to open a tab in the default browser...", err)
		os.Exit(-1)
	}
}

// serverUrl() returns the URL for the todo list server
func serverUrl() string {
	return fmt.Sprintf("http://localhost:%d/", *port)
}
