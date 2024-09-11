package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	Path     string
	FuncName string
	Package  string
	Import   string
}

func findPublicFuncsReturningHPartial(dir string) ([]string, error) {
	var functions []string

	// Walk through the directory to find all Go files.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process Go files.
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Parse the Go file.
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
		if err != nil {
			return err
		}

		// Inspect the AST for function declarations.
		ast.Inspect(node, func(n ast.Node) bool {
			// Check if the node is a function declaration.
			if funcDecl, ok := n.(*ast.FuncDecl); ok {
				// Only consider exported (public) functions.
				if funcDecl.Name.IsExported() {
					// Check the return type.
					if funcDecl.Type.Results != nil {
						for _, result := range funcDecl.Type.Results.List {
							// Check if the return type is *h.Partial.
							if starExpr, ok := result.Type.(*ast.StarExpr); ok {
								if selectorExpr, ok := starExpr.X.(*ast.SelectorExpr); ok {
									// Check if the package name is 'h' and type is 'Partial'.
									if ident, ok := selectorExpr.X.(*ast.Ident); ok && ident.Name == "h" {
										if selectorExpr.Sel.Name == "Partial" {
											functions = append(functions, funcDecl.Name.Name)
											break
										}
									}
								}
							}
						}
					}
				}
			}
			return true
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return functions, nil
}

func findPublicFuncsReturningHPage(dir string) ([]Page, error) {
	var pages = make([]Page, 0)

	// Walk through the directory to find all Go files.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process Go files.
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Parse the Go file.
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
		if err != nil {
			return err
		}

		// Inspect the AST for function declarations.
		ast.Inspect(node, func(n ast.Node) bool {
			// Check if the node is a function declaration.
			if funcDecl, ok := n.(*ast.FuncDecl); ok {
				// Only consider exported (public) functions.
				if funcDecl.Name.IsExported() {
					// Check the return type.
					if funcDecl.Type.Results != nil {
						for _, result := range funcDecl.Type.Results.List {
							// Check if the return type is *h.Partial.
							if starExpr, ok := result.Type.(*ast.StarExpr); ok {
								if selectorExpr, ok := starExpr.X.(*ast.SelectorExpr); ok {
									// Check if the package name is 'h' and type is 'Partial'.
									if ident, ok := selectorExpr.X.(*ast.Ident); ok && ident.Name == "h" {
										if selectorExpr.Sel.Name == "Page" {
											pages = append(pages, Page{
												Package:  node.Name.Name,
												Import:   fmt.Sprintf("mhtml/%s", filepath.Dir(path)),
												Path:     path,
												FuncName: funcDecl.Name.Name,
											})
											break
										}
									}
								}
							}
						}
					}
				}
			}
			return true
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return pages, nil
}

func buildGetPartialFromContext(builder *CodeBuilder, funcs []string) {
	fName := "GetPartialFromContext"

	body := `
		path := ctx.Path()
	`

	for _, f := range funcs {
		if f == fName {
			continue
		}
		body += fmt.Sprintf(`
			if path == "%s" || path == "/mhtml/partials.%s" {
				return %s(ctx)
			}
		`, f, f, f)
	}

	body += "return nil"

	f := Function{
		Name: fName,
		Parameters: []NameType{
			{Name: "ctx", Type: "*fiber.Ctx"},
		},
		Return: []ReturnType{
			{Type: "*h.Partial"},
		},
		Body: body,
	}

	builder.Append(builder.BuildFunction(f))
}

func writePartialsFile() {
	cwd, _ := os.Getwd()
	partialPath := filepath.Join(cwd, "partials")
	funcs, err := findPublicFuncsReturningHPartial(partialPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	builder := NewCodeBuilder(nil)
	builder.AppendLine(`// Package partials THIS FILE IS GENERATED. DO NOT EDIT.`)
	builder.AppendLine("package partials")
	builder.AddImport("mhtml/h")
	builder.AddImport("github.com/gofiber/fiber/v2")

	buildGetPartialFromContext(builder, funcs)

	WriteFile(filepath.Join("partials", "generated.go"), func(content *ast.File) string {
		return builder.String()
	})
}

func formatRoute(path string) string {
	path = strings.TrimSuffix(path, "index.go")
	path = strings.TrimSuffix(path, ".go")
	path = strings.TrimPrefix(path, "pages/")
	path = strings.ReplaceAll(path, "$", ":")
	path = strings.ReplaceAll(path, "_", "/")
	path = strings.ReplaceAll(path, ".", "/")
	if path == "" {
		return "/"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if strings.HasSuffix(path, "/") {
		return path[:len(path)-1]
	}
	return path
}

func writePagesFile() {
	builder := NewCodeBuilder(nil)
	builder.AppendLine(`// Package pages THIS FILE IS GENERATED. DO NOT EDIT.`)
	builder.AppendLine("package pages")
	builder.AddImport("github.com/gofiber/fiber/v2")
	builder.AddImport("mhtml/h")

	pages, _ := findPublicFuncsReturningHPage("pages")

	for _, page := range pages {
		if page.Import != "" && page.Package != "pages" {
			builder.AddImport(page.Import)
		}
	}

	fName := "RegisterPages"
	body := `
	`

	for _, page := range pages {
		call := fmt.Sprintf("%s.%s", page.Package, page.FuncName)
		if page.Package == "pages" {
			call = page.FuncName
		}

		body += fmt.Sprintf(`
			f.Get("%s", func(ctx *fiber.Ctx) error {
				return h.HtmlView(ctx, %s(ctx))
			})
		`, formatRoute(page.Path), call)
	}

	f := Function{
		Name: fName,
		Parameters: []NameType{
			{Name: "f", Type: "*fiber.App"},
		},
		Body: body,
	}

	builder.Append(builder.BuildFunction(f))

	WriteFile("pages/generated.go", func(content *ast.File) string {
		return builder.String()
	})
}

func main() {
	writePartialsFile()
	writePagesFile()
}
