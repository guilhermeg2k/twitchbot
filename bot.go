package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Bot twitch bot
type Bot struct {
	Hello          string
	Commands       []string
	CommandsReturn map[string]string
}

func (twitch *Twitch) handleCommand() {
	for {
		command := <-twitch.CommandsChannel
		fmt.Println(twitch.bot.CommandsReturn["#github"])
		for _, botCommand := range twitch.bot.Commands {
			if botCommand == command.Command {
				twitch.sendMenssage(fmt.Sprintf("@%s %s", command.Nick, twitch.bot.CommandsReturn[command.Command]))
			}
		}
	}
}
func InicializeBot() Bot {
	botFile, err := os.Open("bot.json")
	checkError(err)
	fileData, err := ioutil.ReadAll(botFile)
	checkError(err)
	var bot Bot
	err = json.Unmarshal(fileData, &bot)
	checkError(err)
	return bot
}
