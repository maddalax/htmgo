package astgen

import (
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

func WriteFile(path string, cb func(content *ast.File) string) {
	currentDir := process.GetWorkingDir()

	path = filepath.Join(currentDir, path)

	slog.Debug("astgen.WriteFile", slog.String("path", path))

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
			log.Printf("Failed to format source: %v\n", err.Error())
			data := string(bytes)
			println(data)
			return
		}
	}

	// Define the file path where you want to save the buffer
	process.Run("git add "+path, process.Silent)

	// Save the buffer to a file
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		PanicF("Failed to write buffer to file: %v", err)
	}
}
