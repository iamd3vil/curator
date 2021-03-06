package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// Hub holds all application context
type Hub struct {
	watcher *fsnotify.Watcher
}

// AddWatchedDir adds the given path to monitor
func (h *Hub) AddWatchedDir(path string) error {
	return h.watcher.Add(path)
}

// HandleEvents handles events coming from the watcher
func (h *Hub) HandleEvents() {
	for {
		select {
		case event, ok := <-h.watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				if isEbook(event.Name) {
					err := convertFileToMobi(event.Name, cfg.app.OutputDir)
					if err != nil {
						log.Printf("error while conversion:%s:%v", event.Name, err)
					}
				}
			}
		case err, ok := <-h.watcher.Errors:
			if !ok {
				return
			}
			log.Println("error watching input dir:", err)
		}
	}
}

// NewHub returns a new hub instance
func NewHub() (*Hub, error) {
	h := &Hub{}

	// Start a watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	h.watcher = watcher

	return h, nil
}
