package htmltogo

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"slices"
	"strings"
)

func Indent(input string) string {
	fset := token.NewFileSet()
	// Parse the code string into an AST
	f, err := parser.ParseFile(fset, "", input, parser.ParseComments)

	if err != nil {
		return input
	}

	htmgoComponentTypes := []string{
		"h.Element",
		"h.Page",
		"h.Partial",
		"h.Ren",
	}

	for _, decl := range f.Decls {
		switch c := decl.(type) {
		case *ast.FuncDecl:

			if c.Type.Results == nil || len(c.Type.Results.List) == 0 {
				continue
			}

			returnType := c.Type.Results.List[0].Type

			isHtmgoComponent := false
			if v, ok := returnType.(*ast.StarExpr); ok {
				if x, ok := v.X.(*ast.SelectorExpr); ok {
					name := x.X.(*ast.Ident).Name
					str := name + "." + x.Sel.Name
					isHtmgoComponent = slices.Contains(htmgoComponentTypes, str)
				}
			}

			if !isHtmgoComponent {
				continue
			}

			var isHTag = func(n ast.Expr) bool {
				switch argc := n.(type) {
				// If the first argument is another node, add an indent
				case *ast.CallExpr:
					if v, ok := argc.Fun.(*ast.SelectorExpr); ok {
						if v2, ok := v.X.(*ast.Ident); ok {
							if v2.Name == "h" || v2.Name == "js" {
								return true
							}
						}
					}
				}
				return false
			}

			var indent = func(children []ast.Expr) []ast.Expr {
				children = append(children, ast.NewIdent("INDENTME"))
				return children
			}

			astutil.Apply(c.Body, nil, func(cursor *astutil.Cursor) bool {
				switch n := cursor.Node().(type) {
				case *ast.CallExpr:
					newChildren := make([]ast.Expr, 0)

					hasAnyHElements := false

					for _, arg := range n.Args {
						if isHTag(arg) {
							hasAnyHElements = true
							break
						}
					}

					for i, arg := range n.Args {

						if len(n.Args) == 1 && isHTag(arg) {
							newChildren = indent(newChildren)
							newChildren = append(newChildren, arg)
							newChildren = indent(newChildren)
							continue
						}

						if !hasAnyHElements {
							newChildren = append(newChildren, arg)
							continue
						}

						if len(n.Args) > 1 {
							if i == 0 {
								newChildren = indent(newChildren)
							}
						}
						newChildren = append(newChildren, arg)
						if len(n.Args) > 1 {
							newChildren = indent(newChildren)
						}
					}
					n.Args = newChildren
					return true
				}
				return true
			})
		}
	}

	// Convert the AST node to a string
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, f); err != nil {
		fmt.Println("Error printing AST:", err)
		return input
	}

	// Output the formatted code
	indented := strings.ReplaceAll(buf.String(), "INDENTME,", "\n\t\t")
	indented = strings.ReplaceAll(indented, ", INDENTME", ", \n\t\t")

	formatted, err := format.Source([]byte(indented))

	if err != nil {
		return input
	}

	return string(formatted)
}
