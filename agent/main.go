package main

import (
	"github.com/Epyklab/trident/agent/logparser"
	"github.com/Epyklab/trident/agent/logstream"
	"github.com/Epyklab/trident/agent/malhandler"
	"github.com/Epyklab/trident/agent/utils"
	"github.com/Epyklab/trident/agent/webserver"
)

func main() {

	config, _ := utils.ParseConfig()

	logFiles := config.Files

	lineChan := make(chan string)

	for _, logFile := range logFiles {
		go logparser.TailFile(logFile, lineChan)
	}

	go malhandler.WatchDirectory("/opt/trident/uploads", lineChan)

	go logstream.ShipLogs(lineChan)

	go webserver.Router(lineChan)

	select {}

}
