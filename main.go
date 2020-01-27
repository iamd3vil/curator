package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	initConfig()
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs

		fmt.Println("quitting.....")
		done <- true
	}()

	h, err := NewHub()
	if err != nil {
		log.Fatalf("error while initilizing hub: %v", err)
	}
	defer h.watcher.Close()

	// Handle events
	go h.HandleEvents()

	err = h.AddWatchedDir(cfg.InputDir)
	if err != nil {
		log.Fatalf("error while monitoring: %s: %v", cfg.InputDir, err)
	}

	<-done
}
