package astgen

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/tasks/process"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/mod/modfile"
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

type Partial struct {
	FuncName string
	Package  string
	Import   string
}

func sliceCommonPrefix(dir1, dir2 string) string {
	// Use filepath.Clean to normalize the paths
	dir1 = filepath.Clean(dir1)
	dir2 = filepath.Clean(dir2)

	// Find the common prefix
	commonPrefix := dir1
	if len(dir1) > len(dir2) {
		commonPrefix = dir2
	}

	for !strings.HasPrefix(dir1, commonPrefix) {
		commonPrefix = filepath.Dir(commonPrefix)
	}

	// Slice off the common prefix
	slicedDir1 := strings.TrimPrefix(dir1, commonPrefix)
	slicedDir2 := strings.TrimPrefix(dir2, commonPrefix)

	// Remove leading slashes
	slicedDir1 = strings.TrimPrefix(slicedDir1, string(filepath.Separator))
	slicedDir2 = strings.TrimPrefix(slicedDir2, string(filepath.Separator))

	// Return the longer one
	if len(slicedDir1) > len(slicedDir2) {
		return slicedDir1
	}
	return slicedDir2
}

func findPublicFuncsReturningHPartial(dir string) ([]Partial, error) {
	var partials []Partial
	cwd := process.GetWorkingDir()

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
				// Only consider exported (public) partials.
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
											partials = append(partials, Partial{
												Package:  node.Name.Name,
												Import:   sliceCommonPrefix(cwd, filepath.Dir(path)),
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

	return partials, nil
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
												Import:   filepath.Dir(path),
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

func buildGetPartialFromContext(builder *CodeBuilder, partials []Partial) {
	fName := "GetPartialFromContext"

	body := `
		path := ctx.Request().URL.Path
	`

	moduleName := GetModuleName()
	for _, f := range partials {
		if f.FuncName == fName {
			continue
		}
		caller := fmt.Sprintf("%s.%s", f.Package, f.FuncName)
		path := fmt.Sprintf("/%s/%s.%s", moduleName, f.Import, f.FuncName)

		body += fmt.Sprintf(`
			if path == "%s" || path == "%s" {
				cc := ctx.(*h.RequestContext)
				return %s(cc)
			}
		`, f.FuncName, path, caller)
	}

	body += "return nil"

	f := Function{
		Name: fName,
		Parameters: []NameType{
			{Name: "ctx", Type: "echo.Context"},
		},
		Return: []ReturnType{
			{Type: "*h.Partial"},
		},
		Body: body,
	}

	builder.Append(builder.BuildFunction(f))

	registerFunction := fmt.Sprintf(`
		func RegisterPartials(f *echo.Echo) {
			f.Any("%s/partials*", func(ctx echo.Context) error {
			partial := GetPartialFromContext(ctx)
			if partial == nil {
				return ctx.NoContent(404)
			}
			return h.PartialView(ctx, partial)
			})
		}
	`, moduleName)

	builder.AppendLine(registerFunction)
}

func writePartialsFile() {
	cwd := process.GetWorkingDir()
	partialPath := filepath.Join(cwd, "partials")
	partials, err := findPublicFuncsReturningHPartial(partialPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(partials) == 0 {
		return
	}

	builder := NewCodeBuilder(nil)
	builder.AppendLine(`// Package partials THIS FILE IS GENERATED. DO NOT EDIT.`)
	builder.AppendLine("package load")
	builder.AddImport("github.com/maddalax/htmgo/framework/h")
	builder.AddImport("github.com/labstack/echo/v4")

	moduleName := GetModuleName()
	for _, partial := range partials {
		if partial.Import == "partials/load" {
			continue
		}
		builder.AddImport(fmt.Sprintf(`%s/%s`, moduleName, partial.Import))
	}

	buildGetPartialFromContext(builder, partials)

	WriteFile(filepath.Join("partials", "load", "generated.go"), func(content *ast.File) string {
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
	builder.AddImport("github.com/labstack/echo/v4")
	builder.AddImport("github.com/maddalax/htmgo/framework/h")

	pages, _ := findPublicFuncsReturningHPage("pages")

	if len(pages) == 0 {
		return
	}

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
			f.GET("%s", func(ctx echo.Context) error {
				cc := ctx.(*h.RequestContext)
				return h.HtmlView(ctx, %s(cc))
			})
		`, formatRoute(page.Path), call)
	}

	f := Function{
		Name: fName,
		Parameters: []NameType{
			{Name: "f", Type: "*echo.Echo"},
		},
		Body: body,
	}

	builder.Append(builder.BuildFunction(f))

	WriteFile("pages/generated.go", func(content *ast.File) string {
		return builder.String()
	})
}

func GetModuleName() string {
	wd := process.GetWorkingDir()
	modPath := filepath.Join(wd, "go.mod")
	goModBytes, err := os.ReadFile(modPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading go.mod: %v\n", err)
		return ""
	}
	modName := modfile.ModulePath(goModBytes)
	return modName
}

func GenAst(exitOnError bool) error {
	if GetModuleName() == "" {
		if exitOnError {
			os.Exit(1)
		}
		return fmt.Errorf("error getting module name")
	}
	writePartialsFile()
	writePagesFile()
	return nil
}
