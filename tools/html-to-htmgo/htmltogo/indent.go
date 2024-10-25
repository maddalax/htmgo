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
	f, err := parser.ParseFile(fset, "", input, 0)

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
		switch decl.(type) {
		case *ast.FuncDecl:
			component := decl.(*ast.FuncDecl)
			returnType := component.Type.Results.List[0].Type

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

			astutil.Apply(component.Body, nil, func(cursor *astutil.Cursor) bool {
				switch n := cursor.Node().(type) {
				case *ast.CallExpr:
					newChildren := make([]ast.Expr, 0)
					for i, arg := range n.Args {
						if i == 0 {
							switch arg.(type) {
							// If the first argument is another node, add an indent
							case *ast.CallExpr:
								newChildren = append(newChildren, ast.NewIdent("INDENTME"))
							}
						}
						newChildren = append(newChildren, arg)
						newChildren = append(newChildren, ast.NewIdent("INDENTME"))
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
