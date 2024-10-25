package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

const frameworkRepo = "github.com/maddalax/htmgo/framework"
const htmlToHtmgoRepo = "github.com/maddalax/htmgo/tools/html-to-htmgo"

var depsToUpdate = []string{
	frameworkRepo,
	htmlToHtmgoRepo,
}

const githubAPIURL = "https://api.github.com/repos/maddalax/htmgo/commits"

// Commit represents the structure of a commit object returned by the GitHub API.
type Commit struct {
	Sha string `json:"sha"`
}

func RunCommand(path string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = path // Set the working directory to the current path.
	return cmd.Run()
}

func main() {
	// Get the latest commit hash from the remote repository.
	latestCommitHash, err := getLatestCommitHash()
	if err != nil {
		fmt.Println("Error getting latest commit hash:", err)
		return
	}

	// Start from the current directory.
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	wg := sync.WaitGroup{}
	// Walk through directories.
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the directory contains a go.mod file.
		if info.IsDir() && fileExists(filepath.Join(path, "go.mod")) {
			goModPath := filepath.Join(path, "go.mod")
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, s := range depsToUpdate {
					updateDepToLatestVersion(s, goModPath, latestCommitHash)
				}
			}()
		}

		return nil
	})

	wg.Wait()

	if err != nil {
		fmt.Println("Error walking through directories:", err)
	}

}

// fileExists checks if a file exists at the given path.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func updateDepToLatestVersion(dep string, goModPath string, latestCommitHash string) {
	if containsDep(dep, goModPath) {
		dir := filepath.Dir(goModPath)
		// Run go get github.com/maddalax/htmgo/framework@<latestCommitHash>.
		fmt.Printf("Running 'go get' with latest commit hash in %s\n", dep)
		RunCommand(dir, "go", "get", fmt.Sprintf("%s@%s", dep, latestCommitHash))
		RunCommand(dir, "go", "mod", "tidy")
	}
}

// containsDep checks if 'github.com/maddalax/htmgo/framework' is in the go.mod file.
func containsDep(dep string, goModPath string) bool {
	file, err := os.Open(goModPath)
	if err != nil {
		fmt.Println("Error opening go.mod file:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), dep) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading go.mod file:", err)
	}

	return false
}

// getLatestCommitHash fetches the latest commit hash from the GitHub API for the framework repository.
func getLatestCommitHash() (string, error) {
	// Create an HTTP request to the GitHub API to get the latest commit.
	resp, err := http.Get(githubAPIURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest commit: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch latest commit: status code %d", resp.StatusCode)
	}

	// Parse the JSON response.
	var commits []Commit
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return "", fmt.Errorf("failed to decode JSON response: %v", err)
	}

	if len(commits) == 0 {
		return "", fmt.Errorf("no commits found in repository")
	}

	// Return the SHA of the latest commit.
	return commits[0].Sha, nil
}
