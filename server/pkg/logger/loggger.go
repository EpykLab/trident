package logger

import (
	"log"
	"os"
)

var (
	// InfoLogger : Exported Info Logger
	InfoLogger *log.Logger
	// ErrorLogger : Exported Error Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
