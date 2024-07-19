package logger

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

type LFormat struct {
	UUID     uuid.UUID `json:"uuid"`
	Datetime time.Time `json:"datetime"`
	Source   string    `json:"source"`
	Message  string    `json:"message"`
	// TODO: possible to add in optional values?
}

func getUUID() uuid.UUID {
	u := uuid.New()

	return u
}

func nowTime() time.Time {
	t := time.Now().UTC()

	return t
}

func Entry(message string, source string) []byte {

	var uuid = getUUID()
	var time = nowTime()

	entry := &LFormat{
		UUID:     uuid,
		Datetime: time,
		Source:   source,
		Message:  message,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
