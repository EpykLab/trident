// logger/logger.go
package logger

import (
	"log"
	"os"
)

var (
	// Log is an exported Logger
	Log *log.Logger
)

func init() {
	// Open the log file
	file, err := os.OpenFile("/var/log/trident.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	// Create a new logger
	Log = log.New(file, "TRIDENT: ", log.Ldate|log.Ltime|log.Lshortfile)
}
