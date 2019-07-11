package main

import (
	"strings"
)

func handdleCliCommand(command string, twitch *Twitch) {
	command = strings.Replace(command, "\n", "", 1)
	switch command {
	case "reload":
		twitch.bot = InicializeBot()
	case "reloadSettings":
		twitch.S = getSettings()
	}
}
