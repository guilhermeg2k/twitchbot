package main

import (
	"fmt"
	"log"
)

func main() {
	logFile := getLogFile()
	log.SetOutput(logFile)
	log.Println("Bot started ...")
	fmt.Println("Bot started ...")
	SocketClient(getSettings())
}
