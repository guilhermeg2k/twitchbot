package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//Twitch connection
type Twitch struct {
	Conn            net.Conn
	S               Settings
	bot             Bot
	CommandsChannel chan Command
}

var cliChannel chan string

//SocketClient inicialize
func SocketClient(s Settings) {
	var userInput string
	bot := InicializeBot()
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	conn, err := net.Dial("tcp", addr)
	twitch := Twitch{Conn: conn, S: s, bot: bot, CommandsChannel: make(chan Command)}
	checkError(err)
	go twitch.handle()
	twitch.authenticate()
	twitch.joinChannel()
	//twitch.sendMenssage(bot.Hello)
	stdIn := bufio.NewReader(os.Stdin)
	for {
		userInput, err = stdIn.ReadString('\n')
		checkError(err)
		handdleCliCommand(userInput, &twitch)
	}
}

func (twitch *Twitch) handle() {
	go twitch.handleCommand()
	var err error
	var buff []byte
	var n int
	var data string
	for {
		buff = make([]byte, 1024)
		n, err = twitch.Conn.Read(buff)
		if err != nil {
			checkError(err)
		}
		data = string(buff[0:n])
		fmt.Println(data)
		if strings.Contains(data, "PRIVMSG") {
			msg := getCommandFromData(data)
			for _, command := range twitch.bot.Commands {
				if _, ok := command[msg.Msg]; ok {
					fmt.Println(msg.Msg)
					twitch.CommandsChannel <- msg
					break
				}
			}
			if strings.Contains(data, "PING") {
				twitch.pong()
			}
		}
	}
}

func (twitch *Twitch) pong() {
	twitch.Conn.Write([]byte("PONG :tmi.twitch.tv\r\n"))
}
func (twitch *Twitch) authenticate() {
	twitch.Conn.Write([]byte(fmt.Sprintf("PASS %s\r\n", twitch.S.OAuth)))
	twitch.Conn.Write([]byte(fmt.Sprintf("NICK %s\r\n", twitch.S.User)))
}

func (twitch *Twitch) joinChannel() {
	twitch.Conn.Write([]byte(fmt.Sprintf("JOIN #%s\r\n", twitch.S.Channel)))
}

func (twitch *Twitch) sendMenssage(msg string) {
	twitch.Conn.Write([]byte(fmt.Sprintf("PRIVMSG #%s :%s\r\n", twitch.S.Channel, msg)))
}

func getCommandFromData(data string) Command {
	var msg Command
	dataSplited := strings.Split(data, " ")
	aux := strings.Split(dataSplited[0], "!")
	aux = strings.Split(aux[0], ":")
	msg.Nick = aux[1]
	aux = strings.Split(dataSplited[3], ":")
	aux = strings.Split(aux[1], "\r")
	msg.Msg = aux[0]
	return msg
}
