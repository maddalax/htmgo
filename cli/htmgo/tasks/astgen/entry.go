package astgen

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
	"slices"
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

const GeneratedDirName = "__htmgo"
const HttpModuleName = "net/http"
const ChiModuleName = "github.com/go-chi/chi/v5"
const ModuleName = "github.com/maddalax/htmgo/framework/h"

var PackageName = fmt.Sprintf("package %s", GeneratedDirName)
var GeneratedFileLine = fmt.Sprintf("// Package %s THIS FILE IS GENERATED. DO NOT EDIT.", GeneratedDirName)

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

	// Disconnect leading slashes
	slicedDir1 = strings.TrimPrefix(slicedDir1, string(filepath.Separator))
	slicedDir2 = strings.TrimPrefix(slicedDir2, string(filepath.Separator))

	// Return the longer one
	if len(slicedDir1) > len(slicedDir2) {
		return slicedDir1
	}
	return slicedDir2
}

func findPublicFuncsReturningHPartial(dir string, predicate func(partial Partial) bool) ([]Partial, error) {
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
											p := Partial{
												Package:  node.Name.Name,
												Import:   sliceCommonPrefix(cwd, strings.ReplaceAll(filepath.Dir(path), `\`, `/`)),
												FuncName: funcDecl.Name.Name,
											}
											if predicate(p) {
												partials = append(partials, p)
											}
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
												Import:   strings.ReplaceAll(filepath.Dir(path), `\`, `/`),
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
		path := r.URL.Path
	`

	if len(partials) == 0 {
		body = ""
	}

	moduleName := GetModuleName()
	for _, f := range partials {
		if f.FuncName == fName {
			continue
		}
		caller := fmt.Sprintf("%s.%s", f.Package, f.FuncName)
		path := fmt.Sprintf("/%s/%s.%s", moduleName, f.Import, f.FuncName)

		body += fmt.Sprintf(`
			if path == "%s" || path == "%s" {
				cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)
				return %s(cc)
			}
		`, f.FuncName, path, caller)
	}

	body += "return nil"

	f := Function{
		Name: fName,
		Parameters: []NameType{
			{Name: "r", Type: "*http.Request"},
		},
		Return: []ReturnType{
			{Type: "*h.Partial"},
		},
		Body: body,
	}

	builder.Append(builder.BuildFunction(f))

	registerFunction := fmt.Sprintf(`
		func RegisterPartials(router *chi.Mux) {
				router.Handle("/%s/partials*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					partial := GetPartialFromContext(r)
					if partial == nil {
						w.WriteHeader(404)	
						return
					}		
					h.PartialView(w, partial)	
			}))
		}
	`, moduleName)

	builder.AppendLine(registerFunction)
}

func writePartialsFile() {
	cwd := process.GetWorkingDir()
	partialPath := filepath.Join(cwd, "partials")
	partials, err := findPublicFuncsReturningHPartial(partialPath, func(partial Partial) bool {
		return partial.FuncName != "GetPartialFromContext"
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	builder := NewCodeBuilder(nil)
	builder.AppendLine(GeneratedFileLine)
	builder.AppendLine(PackageName)
	builder.AddImport(ModuleName)
	builder.AddImport(HttpModuleName)
	builder.AddImport(ChiModuleName)

	moduleName := GetModuleName()
	for _, partial := range partials {
		builder.AddImport(fmt.Sprintf(`%s/%s`, moduleName, partial.Import))
	}

	buildGetPartialFromContext(builder, partials)

	WriteFile(filepath.Join(GeneratedDirName, "partials-generated.go"), func(content *ast.File) string {
		return builder.String()
	})
}

func formatRoute(path string) string {
	path = strings.TrimSuffix(path, "index.go")
	path = strings.TrimSuffix(path, ".go")
	path = strings.TrimPrefix(path, "pages/")
	path = strings.TrimPrefix(path, "pages\\")
	path = strings.ReplaceAll(path, "$", ":")
	path = strings.ReplaceAll(path, "_", "/")
	path = strings.ReplaceAll(path, ".", "/")
	path = strings.ReplaceAll(path, "\\", "/")

	parts := strings.Split(path, "/")

	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			parts[i] = fmt.Sprintf("{%s}", part[1:])
		}
	}

	path = strings.Join(parts, "/")

	if path == "" {
		return "/"
	}
	if !strings.HasPrefix(path, "/") {
		path = filepath.Join("/", path)
	}
	if strings.HasSuffix(path, "/") {
		return strings.ReplaceAll(path[:len(path)-1], `\`, "/")
	}
	return strings.ReplaceAll(filepath.Clean(path), `\`, "/")
}

func writePagesFile() {

	builder := NewCodeBuilder(nil)
	builder.AppendLine(GeneratedFileLine)
	builder.AppendLine(PackageName)
	builder.AddImport(HttpModuleName)
	builder.AddImport(ChiModuleName)

	pages, _ := findPublicFuncsReturningHPage("pages")

	if len(pages) > 0 {
		builder.AddImport(ModuleName)
	}

	for _, page := range pages {
		if page.Import != "" {
			moduleName := GetModuleName()
			builder.AddImport(
				fmt.Sprintf(`%s/%s`, moduleName, page.Import),
			)
		}
	}

	fName := "RegisterPages"
	body := `
	`

	for _, page := range pages {
		call := fmt.Sprintf("%s.%s", page.Package, page.FuncName)

		body += fmt.Sprintf(
			`
			router.Get("%s", func(writer http.ResponseWriter, request *http.Request) {
				cc := request.Context().Value(h.RequestContextKey).(*h.RequestContext)
				h.HtmlView(writer, %s(cc))
			})
			`, formatRoute(page.Path), call,
		)
	}

	f := Function{
		Name: fName,
		Parameters: []NameType{
			{Name: "router", Type: "*chi.Mux"},
		},
		Body: body,
	}

	builder.Append(builder.BuildFunction(f))

	WriteFile(filepath.Join(GeneratedDirName, "pages-generated.go"), func(content *ast.File) string {
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

func GenAst(flags ...process.RunFlag) error {
	if GetModuleName() == "" {
		if slices.Contains(flags, process.ExitOnError) {
			os.Exit(1)
		}
		return fmt.Errorf("error getting module name")
	}
	writePartialsFile()
	writePagesFile()

	WriteFile("__htmgo/setup-generated.go", func(content *ast.File) string {

		return fmt.Sprintf(`
			// Package __htmgo THIS FILE IS GENERATED. DO NOT EDIT.
			package __htmgo

			import (
				"%s"
			)

			func Register(r *chi.Mux) {
				RegisterPartials(r)
				RegisterPages(r)
			}
		`, ChiModuleName)
	})

	return nil
}
