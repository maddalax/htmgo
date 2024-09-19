package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var ignoredDirs = []string{".git", ".idea", "node_modules", "vendor"}

func startWatcher(cb func(file []*fsnotify.Event)) {
	events := make([]*fsnotify.Event, 0)
	debouncer := NewDebouncer(100 * time.Millisecond)

	defer func() {
		if r := recover(); r != nil {
			slog.Debug("Recovered from fatal error:", slog.String("error", r.(error).Error()))
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
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Remove) || event.Has(fsnotify.Rename) {
					events = append(events, &event)
					debouncer.Do(func() {
						cb(events)
						events = make([]*fsnotify.Event, 0)
					})
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				slog.Error("error:", err.Error())
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
				slog.Error("Error adding directory to watcher:", err)
			} else {
				slog.Debug("Watching directory:", slog.String("path", path))
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
