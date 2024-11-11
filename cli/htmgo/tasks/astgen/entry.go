package astgen

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/internal/dirutil"
	"github.com/maddalax/htmgo/cli/htmgo/tasks/process"
	"github.com/maddalax/htmgo/framework/h"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/mod/modfile"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
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
	Path     string
}

const GeneratedDirName = "__htmgo"
const HttpModuleName = "net/http"
const ChiModuleName = "github.com/go-chi/chi/v5"
const ModuleName = "github.com/maddalax/htmgo/framework/h"

var PackageName = fmt.Sprintf("package %s", GeneratedDirName)
var GeneratedFileLine = fmt.Sprintf("// Package %s THIS FILE IS GENERATED. DO NOT EDIT.", GeneratedDirName)

func toPascaleCase(input string) string {
	words := strings.Split(input, "_")
	for i := range words {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}
	return strings.Join(words, "")
}

func isValidGoVariableName(name string) bool {
	// Variable name must not be empty
	if name == "" {
		return false
	}
	// First character must be a letter or underscore
	if !unicode.IsLetter(rune(name[0])) && name[0] != '_' {
		return false
	}
	// Remaining characters must be letters, digits, or underscores
	for _, char := range name[1:] {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
			return false
		}
	}
	return true
}

func normalizePath(path string) string {
	return strings.ReplaceAll(path, `\`, "/")
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
		return normalizePath(slicedDir1)
	}
	return normalizePath(slicedDir2)
}

func hasOnlyReqContextParam(funcType *ast.FuncType) bool {
	if len(funcType.Params.List) != 1 {
		return false
	}
	if funcType.Params.List[0].Names == nil {
		return false
	}
	if len(funcType.Params.List[0].Names) != 1 {
		return false
	}
	t := funcType.Params.List[0].Type
	name, ok := t.(*ast.StarExpr)
	if !ok {
		return false
	}
	selectorExpr, ok := name.X.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	ident, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}
	return ident.Name == "h" && selectorExpr.Sel.Name == "RequestContext"
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
										if selectorExpr.Sel.Name == "Partial" && hasOnlyReqContextParam(funcDecl.Type) {
											p := Partial{
												Package:  node.Name.Name,
												Path:     normalizePath(sliceCommonPrefix(cwd, path)),
												Import:   sliceCommonPrefix(cwd, normalizePath(filepath.Dir(path))),
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
										if selectorExpr.Sel.Name == "Page" && hasOnlyReqContextParam(funcDecl.Type) {
											pages = append(pages, Page{
												Package:  node.Name.Name,
												Import:   normalizePath(filepath.Dir(path)),
												Path:     normalizePath(path),
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
	moduleName := GetModuleName()

	var routerHandlerMethod = func(path string, caller string) string {
		return fmt.Sprintf(`
			router.Handle("%s", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		     cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)					
         partial := %s(cc)
					if partial == nil {
						w.WriteHeader(404)	
						return
					}		
					h.PartialView(w, partial)	
			}))`, path, caller)
	}

	handlerMethods := make([]string, 0)

	for _, f := range partials {
		caller := fmt.Sprintf("%s.%s", f.Package, f.FuncName)
		path := fmt.Sprintf("/%s/%s.%s", moduleName, f.Import, f.FuncName)
		handlerMethods = append(handlerMethods, routerHandlerMethod(path, caller))
	}

	registerFunction := fmt.Sprintf(`
		func RegisterPartials(router *chi.Mux) {
				%s
		}
	`, strings.Join(handlerMethods, "\n"))

	builder.AppendLine(registerFunction)
}

func writePartialsFile() {
	config := dirutil.GetConfig()

	cwd := process.GetWorkingDir()
	partialPath := filepath.Join(cwd)
	partials, err := findPublicFuncsReturningHPartial(partialPath, func(partial Partial) bool {
		return partial.FuncName != "GetPartialFromContext"
	})

	partials = h.Filter(partials, func(partial Partial) bool {
		return !dirutil.IsGlobExclude(partial.Path, config.AutomaticPartialRoutingIgnore)
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	builder := NewCodeBuilder(nil)
	builder.AppendLine(GeneratedFileLine)
	builder.AppendLine(PackageName)
	builder.AddImport(ChiModuleName)

	if len(partials) > 0 {
		builder.AddImport(ModuleName)
		builder.AddImport(HttpModuleName)
	}

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
	config := dirutil.GetConfig()

	builder := NewCodeBuilder(nil)
	builder.AppendLine(GeneratedFileLine)
	builder.AppendLine(PackageName)
	builder.AddImport(HttpModuleName)
	builder.AddImport(ChiModuleName)

	pages, _ := findPublicFuncsReturningHPage("pages")

	pages = h.Filter(pages, func(page Page) bool {
		return !dirutil.IsGlobExclude(page.Path, config.AutomaticPageRoutingIgnore)
	})

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

func writeAssetsFile() {
	cwd := process.GetWorkingDir()
	config := dirutil.GetConfig()

	slog.Debug("writing assets file", slog.String("cwd", cwd), slog.String("config", config.PublicAssetPath))

	distAssets := filepath.Join(cwd, "assets", "dist")
	hasAssets := false

	builder := strings.Builder{}

	builder.WriteString(`package assets`)
	builder.WriteString("\n")

	filepath.WalkDir(distAssets, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasPrefix(d.Name(), ".") {
			return nil
		}

		path = strings.ReplaceAll(path, distAssets, "")
		httpUrl := normalizePath(fmt.Sprintf("%s%s", config.PublicAssetPath, path))

		path = normalizePath(path)
		path = strings.ReplaceAll(path, "/", "_")
		path = strings.ReplaceAll(path, "//", "_")

		name := strings.ReplaceAll(path, ".", "_")
		name = strings.ReplaceAll(name, "-", "_")

		name = toPascaleCase(name)

		if isValidGoVariableName(name) {
			builder.WriteString(fmt.Sprintf(`const %s = "%s"`, name, httpUrl))
			builder.WriteString("\n")
			hasAssets = true
		}

		return nil
	})

	builder.WriteString("\n")

	str := builder.String()

	if hasAssets {
		WriteFile(filepath.Join(GeneratedDirName, "assets", "assets-generated.go"), func(content *ast.File) string {
			return str
		})
	}

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
	writeAssetsFile()

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
