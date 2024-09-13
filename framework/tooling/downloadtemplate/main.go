package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	cwd, _ := os.Getwd()
	var outPath string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What should we call your new app? Enter name: ")
	outPath, _ = reader.ReadString('\n')
	outPath = strings.ReplaceAll(outPath, "\n", "")
	outPath = strings.ReplaceAll(outPath, " ", "-")
	outPath = strings.ToLower(outPath)

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

	newDir := filepath.Join(cwd, outPath)

	mvCmd := exec.Command("cp", "-vaR", "starter-template/.", ".")
	mvCmd.Dir = newDir
	err = mvCmd.Run()

	if err != nil {
		println("Error moving files %v\n", err)
		return
	}

	rmCmd := exec.Command("rm", "-rf", "starter-template")
	rmCmd.Dir = newDir
	err = rmCmd.Run()

	if err != nil {
		println("Error removing starter-template %v\n", err)
		return
	}

	println("Template downloaded successfully!")
}
