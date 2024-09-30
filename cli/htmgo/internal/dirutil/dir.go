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

func CreateHtmgoDir() {
	if !HasFileFromRoot("__htmgo") {
		CreateDirFromRoot("__htmgo")
	}
}

func CreateDirFromRoot(dir string) error {
	cwd := process.GetWorkingDir()
	path := filepath.Join(cwd, dir)
	return os.MkdirAll(path, 0700)
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

func MoveFile(src, dst string) error {
	slog.Debug("moving file", slog.String("src", src), slog.String("dst", dst))
	// Copy the file.
	err := CopyFile(src, dst)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}
	// Disconnect the source file.
	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("failed to remove source file: %v", err)
	}
	return nil
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

func DeleteDir(dir string) error {
	return os.RemoveAll(dir)
}
