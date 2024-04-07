// copyright epyklab.com 2024

// Handles reading generic log files such as /var/log/auth.log

package logparser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	UUID uuid.UUID `json:"uuid"`
	Time time.Time `json:"time"`
	Path string    `json:"path"`
	Line string    `json:"line"`
}

func UUID() uuid.UUID {
	u := uuid.New()

	return u
}

func nowTime() time.Time {
	t := time.Now().UTC()
	fmt.Print(t)
	return t
}

func marshallToJSON(line string, file string) []byte {
	var uuid = UUID()
	var time = nowTime()

	e := &Entry{UUID: uuid, Time: time, Path: file, Line: line}

	entry, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}

	return entry
}

func TailFile(filePath string, lineChan chan string) (*Entry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Seek(0, io.SeekEnd) // Start reading from the end of the file

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			panic(err)
		}
		if len(line) > 0 {
			entry := marshallToJSON(line, filePath)
			lineChan <- string(entry)

		} else {
			time.Sleep(1 * time.Second) // Sleep if no new line is found
		}
	}
}
