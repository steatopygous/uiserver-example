package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

type BrowserType int

const (
	Chrome BrowserType = iota
	Edge
	Firefox
	Safari
)

type WindowSize struct {
	Width int
	Height int
}

type WindowPosition struct {
	X int
	Y int
}

type Preferences struct {
	Browser  BrowserType
	Position WindowPosition
	Size     WindowSize
}

func LoadPreferences() Preferences {
	path, err := preferencesPath()

	if err != nil {
		return defaultPreferences()
	}

	reader, err := os.Open(path)

	if err != nil {
		return defaultPreferences()
	}

	decoder := json.NewDecoder(reader)

	preferences := Preferences{}

	err = decoder.Decode(&preferences)

	if err != nil {
		fmt.Println("LoadPreferences() - preferences file", path, "could not be parsed")
		return defaultPreferences()
	}

	return preferences
}

func(preferences *Preferences) Save() error {
	path, err := preferencesPath()

	if err != nil {
		return err
	}

	writer, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	encoder := json.NewEncoder(writer)

	err = encoder.Encode(preferences)

	if err != nil {
		fmt.Println("SavePreferences() - failed to write preferences to", path, "...", err)
		return err
	}

	return nil
}

func preferencesPath() (string, error) {
	user, err := user.Current()

	if err != nil {
		fmt.Println("Preferences.Load() - failed to get current user...", err)
		return "", err
	}

	return filepath.FromSlash(user.HomeDir + "/.uiserver-example.json"), nil
}

func defaultPreferences() Preferences {
	return Preferences{Chrome, WindowPosition{400, 100}, WindowSize{700, 900}}
}

func(preferences *Preferences) setWindowSize(size WindowSize) {
	preferences.Size = size

	preferences.Save()
}
