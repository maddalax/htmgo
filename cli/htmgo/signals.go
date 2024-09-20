package main

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"os"
	"os/signal"
	"syscall"
)

func RegisterSignals() chan bool {
	// Create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	// Register the channel to receive interrupt and terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)
	// Run a goroutine to handle signals
	go func() {
		// Block until a signal is received
		sig := <-sigs
		fmt.Println()
		fmt.Println("Received signal:", sig)
		// Perform cleanup
		fmt.Println("Cleaning up...")
		process.KillAll()
		// Signal that cleanup is done
		done <- true
	}()

	return done
}
