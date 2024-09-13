package main

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"log"
	"os"
)

func getModuleVersion(modulePath string) (string, error) {
	// Read the go.mod file
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", fmt.Errorf("error reading go.mod: %v", err)
	}

	// Parse the go.mod file
	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return "", fmt.Errorf("error parsing go.mod: %v", err)
	}

	// Find the module version
	for _, req := range modFile.Require {
		if req.Mod.Path == modulePath {
			return req.Mod.Version, nil
		}
	}

	return "", fmt.Errorf("module %s not found in go.mod", modulePath)
}

func main() {
	modulePath := "github.com/maddalax/mhtml/framework"
	version, err := getModuleVersion(modulePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	assetDir := fmt.Sprintf("%s/go/pkg/mod/%s@%s/assets", dirname, modulePath, version)
	files, _ := os.ReadDir(assetDir)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
