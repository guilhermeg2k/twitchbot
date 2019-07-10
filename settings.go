package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Settings struct json
type Settings struct {
	Host    string
	Port    uint
	User    string
	OAuth   string
	Channel string
}

func getSettings() (settings Settings) {
	settingsFile, err := os.Open("settings.json")
	checkError(err)
	fileData, err := ioutil.ReadAll(settingsFile)
	checkError(err)
	err = json.Unmarshal(fileData, &settings)
	checkError(err)
	return settings
}
