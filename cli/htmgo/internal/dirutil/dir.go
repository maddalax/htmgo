package dirutil

import (
	"errors"
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

func HasFileFromRoot(file string) bool {
	cwd := process.GetWorkingDir()
	path := filepath.Join(cwd, file)
	_, err := os.Stat(path)
	return err == nil
}

func CopyDir(srcDir, dstDir string, predicate func(path string, exists bool) bool) error {
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
			exists := true
			if _, err := os.Stat(dstPath); errors.Is(err, os.ErrNotExist) {
				exists = false
			}
			if predicate(srcPath, exists) {
				err := CopyFile(srcPath, dstPath)
				if err != nil {
					return err
				}
			}
			// If it's a file, copy the file.
		}
		return nil
	})
}

func CopyFile(src, dst string) error {
	slog.Debug("copying file", slog.String("src", src), slog.String("dst", dst))
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
