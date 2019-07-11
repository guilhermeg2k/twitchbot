package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//Bot twitch bot
type Bot struct {
	Hello    string
	Commands []map[string]*Command
}

//Command user Command
type Command struct {
	Awnser  string
	Delay   int
	Nick    string
	Msg     string
	LasTime time.Time
}

func (twitch *Twitch) handleCommand() {
	for {
		command := <-twitch.CommandsChannel
		for _, _command := range twitch.bot.Commands {
			if _, ok := _command[command.Msg]; ok {
				//fmt.Println(twitch.bot)
				if time.Since(_command[command.Msg].LasTime) > time.Second*time.Duration(_command[command.Msg].Delay) {
					_command[command.Msg].LasTime = time.Now()
					twitch.sendMenssage(fmt.Sprintf("@%s %s", command.Nick, _command[command.Msg].Awnser))
				}
			}
		}
	}
}

//InicializeBot get a instance of struct Bot
func InicializeBot() (bot Bot) {
	botFile, err := os.Open("bot.json")
	checkError(err)
	fileData, err := ioutil.ReadAll(botFile)
	checkError(err)
	err = json.Unmarshal(fileData, &bot)
	checkError(err)
	return
}
