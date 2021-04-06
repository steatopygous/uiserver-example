package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/pkg/browser"
)

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

func openChromeWindow(preferences Preferences) {
	chrome := chromePath()
	app := fmt.Sprintf("--app=%s", serverUrl())

	position := fmt.Sprintf("--window-position=%d,%d", preferences.Position.X, preferences.Position.Y)
	size := fmt.Sprintf("--window-size=%d,%d", preferences.Size.Width, preferences.Size.Height)

	time.Sleep(100 * time.Millisecond)

	err := exec.Command(chrome, app, position, size).Run()

	if err != nil {
		fmt.Println("openChromeWindow() - Failed to open Chrome app window...", err)
		os.Exit(-1)
	}
}

func chromePath() string {
	if runtime.GOOS == "windows" {
		return "chrome.exe"
	} else {
		return "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	}
}

func openDefaultBrowserTab() {
	time.Sleep(100 * time.Millisecond)

	err := browser.OpenURL(serverUrl())

	if err != nil {
		fmt.Println("openDefaultBrowserTab() - Failed to open a tab on the default browser...", err)
		os.Exit(-1)
	}
}

func serverUrl() string {
	return fmt.Sprintf("http://localhost:%d/", *port)
}
