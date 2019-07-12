package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//Bot .
type Bot struct {
	Hello        string
	BlockedMsg   string
	Commands     []map[string]*Command
	TimedMsgs    []map[string]*TimedMsg
	BlockedWords []string
}

//TimedMsg .
type TimedMsg struct {
	Msg      string
	Delay    uint
	LastTime time.Time
}

//Command .
type Command struct {
	Anwser  string
	Delay   int
	Nick    string
	Msg     string
	LasTime time.Time
}

func (twitch *Twitch) timedMessages() {
	for {
		for _, msg := range twitch.bot.TimedMsgs {
			for _, key := range msg {
				if time.Since(key.LastTime) > time.Second*time.Duration(key.Delay) {
					key.LastTime = time.Now()
					twitch.sendMessage(key.Msg)
				}
			}
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func (twitch *Twitch) wordFilter() {
	for {
		command := <-twitch.WordFilterChannel
		twitch.timeOut(command.Nick, 1)
		twitch.sendMessage(fmt.Sprintf("@%s %s", command.Nick, twitch.bot.BlockedMsg))
	}
}

func (twitch *Twitch) handleCommand() {
	for {
		command := <-twitch.CommandsChannel
		for _, _command := range twitch.bot.Commands {
			if _, ok := _command[command.Msg]; ok {
				if time.Since(_command[command.Msg].LasTime) > time.Second*time.Duration(_command[command.Msg].Delay) {
					_command[command.Msg].LasTime = time.Now()
					twitch.sendMessage(fmt.Sprintf("@%s %s", command.Nick, _command[command.Msg].Anwser))
				}
			}
		}
	}
}

//InicializeBot .
func InicializeBot() (bot Bot) {
	botFile, err := os.Open("bot.json")
	checkError(err)
	fileData, err := ioutil.ReadAll(botFile)
	checkError(err)
	err = json.Unmarshal(fileData, &bot)
	checkError(err)
	return
}

func (twitch *Twitch) timeOut(nick string, time uint) {
	twitch.sendMessage(fmt.Sprintf("/timeout %s %d", nick, time))
}
