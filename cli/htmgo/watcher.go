package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/cli/htmgo/internal"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/module"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var ignoredDirs = []string{".git", ".idea", "node_modules", "vendor"}

func startWatcher(cb func(version string, file []*fsnotify.Event)) {
	events := make([]*fsnotify.Event, 0)
	debouncer := internal.NewDebouncer(500 * time.Millisecond)

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
				slog.Debug("event:", slog.String("name", event.Name), slog.String("op", event.Op.String()))
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Remove) || event.Has(fsnotify.Rename) {
					events = append(events, &event)
					debouncer.Do(func() {
						seen := make(map[string]bool)
						dedupe := make([]*fsnotify.Event, 0)
						for _, e := range events {
							if _, ok := seen[e.Name]; !ok {
								seen[e.Name] = true
								dedupe = append(dedupe, e)
							}
						}
						cb(uuid.NewString()[0:6], dedupe)
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

	frameworkPath := module.GetDependencyPath("github.com/maddalax/htmgo/framework")

	if !strings.HasPrefix(frameworkPath, "github.com/") {
		assetPath := filepath.Join(frameworkPath, "assets", "dist")
		slog.Debug("Watching directory:", slog.String("path", assetPath))
		watcher.Add(assetPath)
	}

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
