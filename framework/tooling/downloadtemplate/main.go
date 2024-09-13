package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func deleteAllExceptTemplate(outPath string, excludeDir string) {
	// List all files and directories in the root folder
	files, err := os.ReadDir(outPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Iterate through each item in the root folder
	for _, file := range files {
		// Skip the excluded directory
		if file.Name() == excludeDir {
			fmt.Printf("Skipping directory: %s\n", file.Name())
			continue
		}

		// Get full path
		fullPath := filepath.Join(outPath, file.Name())

		// Remove the file or directory
		fmt.Printf("Removing: %s\n", fullPath)
		err := os.RemoveAll(fullPath)
		if err != nil {
			fmt.Printf("Error removing %s: %v\n", fullPath, err)
		} else {
			fmt.Printf("Successfully removed %s\n", fullPath)
		}
	}
}

func main() {
	var outPath string
	fmt.Print("What should we call your new app? ")
	fmt.Scanln(&outPath) // Reads user input from the command line

	if outPath == "" {
		fmt.Println("Please provide a name for your app.")
		return
	}

	excludeDir := "starter-template"

	install := exec.Command("git", "clone", "https://github.com/maddalax/mhtml", "--depth=1", outPath)
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	err := install.Run()

	if err != nil {
		println("Error downloading template %v\n", err)
		return
	}

	deleteAllExceptTemplate(outPath, excludeDir)

	templateDir := filepath.Join(outPath, excludeDir)

	err = exec.Command("mv", templateDir+"/*", "./"+outPath).Run()

	if err != nil {
		println("Error moving files %v\n", err)
		return
	}

	err = exec.Command("rm", "-rf", templateDir).Run()

	if err != nil {
		println("Error removing directory %v\n", err)
		return
	}

	println("Template downloaded successfully!")
}
