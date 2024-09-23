package dirutil

import (
	"errors"
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
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

func DeleteAllExcept(path string, keepDir string) error {
	// Get the info of the path
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		// We only care about directories
		return nil
	}

	// If path is the keepDir, do nothing
	if path == keepDir {
		return nil
	}

	// Check if keepDir is under path (path is an ancestor of keepDir)
	relToKeepDir, err := filepath.Rel(path, keepDir)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(relToKeepDir, "..") {
		// keepDir is under path; recurse into subdirectories
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			entryPath := filepath.Join(path, entry.Name())
			err := DeleteAllExcept(entryPath, keepDir)
			if err != nil {
				return err
			}
		}
	} else {
		// Check if path is under keepDir (path is a descendant of keepDir)
		relFromKeepDir, err := filepath.Rel(keepDir, path)
		if err != nil {
			return err
		}

		if !strings.HasPrefix(relFromKeepDir, "..") {
			// path is under keepDir; do nothing
			return nil
		} else {
			// path is neither keepDir nor related; delete it
			slog.Debug("Deleting directory:", slog.String("path", path))
			err = os.RemoveAll(path)
			if err != nil {
				return err
			}
			// Skip further traversal since the directory is deleted
			return nil
		}
	}

	return nil
}
