package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

var errLog *os.File
var Logger *log.Logger

func init() {
	errLog, err := os.OpenFile(time.Now().Format("20060102")+"log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}

	Logger = log.New(errLog, "applog: ", log.Lshortfile|log.LstdFlags)
}
