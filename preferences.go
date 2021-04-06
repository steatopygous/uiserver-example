package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// WindowSize represents the width and height of the window (used when running Chrome as an app)
type WindowSize struct {
	Width int
	Height int
}

// Preferences represents the user's preferences, to be used next time the app is loaded
type Preferences struct {
	Size     WindowSize
}


// Load() loads the user's preferences
func Load() Preferences {
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
		fmt.Println("Load() - preferences file", path, "could not be parsed")
		return defaultPreferences()
	}

	return preferences
}


// Save() saves the user's preferences to disk, in their home directory
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


// preferencesPath() returns the path to the preferences file
func preferencesPath() (string, error) {
	user, err := user.Current()

	if err != nil {
		fmt.Println("Preferences.Load() - failed to get current user...", err)
		return "", err
	}

	return filepath.FromSlash(user.HomeDir + "/.uiserver-example.json"), nil
}


// defaultPreferences() returns a set of default preferences
func defaultPreferences() Preferences {
	return Preferences{WindowSize{700, 900}}
}


// setWindowSize() updates the window size in the preferences and saves them to disk
func(preferences *Preferences) setWindowSize(size WindowSize) {
	preferences.Size = size

	preferences.Save()
}
