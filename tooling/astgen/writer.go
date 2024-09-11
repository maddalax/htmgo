package main

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func WriteFile(path string, cb func(content *ast.File) string) {
	currentDir, err := os.Getwd()

	path = filepath.Join(currentDir, path)

	dir := filepath.Dir(path)

	os.MkdirAll(dir, 0755)

	bytes, err := os.ReadFile(path)

	if err != nil {
		_, err = os.Create(path)
		bytes = make([]byte, 0)
		if err != nil {
			PanicF("Failed to create file: %v", err)
		}
	}

	// Create a FileSet to manage source file positions
	fset := token.NewFileSet()

	// Parse the file into an AST
	f, _ := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if f == nil {
		f = &ast.File{
			Name:  ast.NewIdent("replacemeplz"), // Set the package name
			Decls: []ast.Decl{},                 // No declarations initially
		}
	}

	bytes = []byte(cb(f))
	formatEnabled := true

	if formatEnabled {
		bytes, err = format.Source(bytes)

		if err != nil {
			log.Printf("Failed to format source: %v\n\n%s", err)
			data := string(bytes)
			println(data)
			return
		}
	}
	// Define the file path where you want to save the buffer

	cmd := exec.Command("git", "add", path)
	err = cmd.Run()

	if err != nil {
		log.Printf("Failed to run git add: %v\n", err)
	}

	// Save the buffer to a file
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		PanicF("Failed to write buffer to file: %v", err)
	}
}
