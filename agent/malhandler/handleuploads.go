package malhandler

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	logger "github.com/Epyklab/trident/agent/utils/logging"
	"github.com/fsnotify/fsnotify"
)

func WatchDirectory(path string, lineChan chan<- string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}

	defer func() {
		if err := watcher.Close(); err != nil {
			log.Fatalf("Failed to close watcher: %v", err)
		}
	}()

	go watchFiles(watcher, lineChan)

	log.Printf("Adding directory to watcher: %s", path)
	if err = watcher.Add(path); err != nil {
		log.Fatalf("Failed to add directory to watcher: %v", err)
	}

	// Add this loop to keep the function alive.
	// You can check if there is a graceful way to stop this loop later,
	// for example checking for a context cancel, os signal interrupt,
	// or a quit signal from another channel.
	for {
		time.Sleep(1 * time.Second)
	}
}

func watchFiles(watcher *fsnotify.Watcher, lineChan chan<- string) {
	log.Println("started watching for file changes")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				log.Println("Failed to read from events channel")
				return
			}
			log.Printf("received event: %s\n", event)

			if event.Op&fsnotify.Write == fsnotify.Write { // fsnotify seems to not trigger on CREATE, just write.
				log.Printf("detected new file: %s\n", event.Name)
				time.Sleep(5 * time.Second) // ensures enough time to finish writes
				processFile(event.Name, lineChan)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				log.Println("Failed to read from errors channel")
				return
			}
			log.Printf("received error: %v\n", err)
		}
	}
}

func processFile(filename string, lineChan chan<- string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("unable to read file: %s, reason: %v\n", filename, err)
	} else {
		encodedContent := base64.StdEncoding.EncodeToString(content)

		event := logger.Entry(
			encodedContent,
			"fileupload")

		// TODO: hash file before deleting?
		fmt.Print(event)
		lineChan <- string(event)
		if err = os.Remove(filename); err != nil {
			log.Printf("unable to delete file: %s, reason: %v\n", filename, err)
		}
	}
}
