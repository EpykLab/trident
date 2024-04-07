package main

import (
	"github.com/Epyklab/trident/logparser"
	"github.com/Epyklab/trident/logstream"
	"github.com/Epyklab/trident/utils"
	"github.com/Epyklab/trident/webserver"
)

func main() {

	config, _ := utils.ParseConfig()

	logFiles := config.Files

	lineChan := make(chan string)

	for _, logFile := range logFiles {
		go logparser.TailFile(logFile, lineChan)
	}

	go logstream.ShipLogs(lineChan)

	go webserver.Router()

	select {}

}
