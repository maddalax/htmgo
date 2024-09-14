package main

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"io"
	"log"
	"os"
	"path/filepath"
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

func copyFile(src, dst string) error {
	// Open the source file for reading.
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()
	// Create the destination file.
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dstFile.Close()
	// Copy the content from srcFile to dstFile.
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents: %v", err)
	}
	return nil
}

// copyDir copies a directory recursively from src to dst.
func copyDir(srcDir, dstDir string) error {
	// Walk the source directory tree.
	return filepath.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Construct the corresponding destination path.
		relPath, err := filepath.Rel(srcDir, srcPath)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)
		if info.IsDir() {
			// If it's a directory, create the corresponding directory in the destination.
			err := os.MkdirAll(dstPath, 0700)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		} else {
			// If it's a file, copy the file.
			err := copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	modulePath := "github.com/maddalax/htmgo/framework"
	version, err := getModuleVersion(modulePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	assetDir := fmt.Sprintf("%s/go/pkg/mod/%s@%s/assets/dist", dirname, modulePath, version)
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal("failed to get cwd")
	}

	destDir := fmt.Sprintf("%s/assets/dist", cwd)

	err = copyDir(assetDir, destDir)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("successfully copied assets to %s\n", destDir)
}
