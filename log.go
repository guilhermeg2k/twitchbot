package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func getLogFile() io.Writer {
	logFileName := fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02 15:04:05"))
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkError(err)
	return file
}
