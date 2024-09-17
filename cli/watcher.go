package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)

var ignoredDirs = []string{".git", ".idea", "node_modules", "vendor"}

func startWatcher(cb func(file []*fsnotify.Event)) {
	//debouncer := NewDebouncer(time.Millisecond * 100)
	events := make([]*fsnotify.Event, 0)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from fatal error:", r)
			// You can log the error here or take other corrective actions
		}
	}()
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					events = append(events, &event)
					go cb(events)
					events = make([]*fsnotify.Event, 0)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	rootDir := "."
	// Walk through the root directory and add all subdirectories to the watcher
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Ignore directories in the ignoredDirs list
		for _, ignoredDir := range ignoredDirs {
			if ignoredDir == info.Name() {
				return filepath.SkipDir
			}
		}
		// Only watch directories
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				log.Println("Error adding directory to watcher:", err)
			} else {
				log.Println("Watching directory:", path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	done := RegisterSignals()
	<-done
	println("process exited")
}
