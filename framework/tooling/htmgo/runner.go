package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

//go:embed Taskfile.yml
var taskFile string

func main() {
	commandMap := make(map[string]*flag.FlagSet)
	commands := []string{"template", "run", "build", "setup", "css", "css-watch", "ast-watch", "watch", "go-watch", "watch"}

	for _, command := range commands {
		commandMap[command] = flag.NewFlagSet(command, flag.ExitOnError)
	}

	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("Usage: htmgo [%s]", strings.Join(commands, " | ")))
		os.Exit(1)
	}

	c := commandMap[os.Args[1]]

	if c == nil {
		fmt.Println(fmt.Sprintf("Usage: htmgo [%s]", strings.Join(commands, " | ")))
		os.Exit(1)
		return
	}

	err := c.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	// Install the latest version of Task
	install := exec.Command("go", "install", "github.com/go-task/task/v3/cmd/task@latest")

	err = install.Run()
	if err != nil {
		fmt.Printf("Error installing task: %v\n", err)
		return
	}

	temp, err := os.CreateTemp("", "Taskfile.yml")

	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		return
	}

	os.WriteFile(temp.Name(), []byte(taskFile), 0644)

	taskName := os.Args[1]

	if taskName == "watch" {
		tasks := []string{"css-watch", "ast-watch", "go-watch"}
		wg := sync.WaitGroup{}
		for _, task := range tasks {
			wg.Add(1)
			go func() {
				cmd := exec.Command("task", "-t", temp.Name(), task)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Printf("Error running task command: %v\n", err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	} else {
		// Define the command and arguments
		cmd := exec.Command("task", "-t", temp.Name(), os.Args[1])
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
}
