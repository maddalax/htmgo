package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

//go:embed Taskfile.yml
var taskFile string

func main() {
	taskFlag := flag.String("task", "", "Specify the task to run (e.g., build, setup). Type -task list to see the list of tasks to run.")

	// Parse the command-line flags
	flag.Parse()

	temp, err := os.CreateTemp("", "Taskfile.yml")

	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		return
	}

	os.WriteFile(temp.Name(), []byte(taskFile), 0644)

	if *taskFlag == "" {
		fmt.Println("Please specify a task to run using the -task flag")
		return
	}

	if *taskFlag == "list" {
		*taskFlag = "--list"
	}

	// Define the command and arguments
	cmd := exec.Command("task", "-t", temp.Name(), *taskFlag)
	// Set the standard output and error to be the same as the Go program
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running task command: %v\n", err)
		return
	}
}
