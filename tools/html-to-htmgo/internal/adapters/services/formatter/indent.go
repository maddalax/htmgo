package formatter

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"strings"
)

func Indent(input string) string {
	fset := token.NewFileSet()
	// Parse the code string into an AST
	f, err := parser.ParseFile(fset, "", input, 0)

	if err != nil {
		return input
	}

	component := f.Decls[1].(*ast.FuncDecl)

	astutil.Apply(component.Body, nil, func(cursor *astutil.Cursor) bool {
		switch n := cursor.Node().(type) {
		case *ast.CallExpr:
			newChildren := make([]ast.Expr, 0)
			for i, arg := range n.Args {
				if i == 0 {
					switch arg.(type) {
					// If the first argument is another node, add an indent
					case *ast.CallExpr:
						newChildren = append(newChildren, ast.NewIdent("INDENT"))
					}
				}
				newChildren = append(newChildren, arg)
				newChildren = append(newChildren, ast.NewIdent("INDENT"))
			}
			n.Args = newChildren
			return true
		}
		return true
	})

	// Convert the AST node to a string
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, component); err != nil {
		fmt.Println("Error printing AST:", err)
		return input
	}

	// Output the formatted code
	indented := strings.ReplaceAll(buf.String(), "INDENT,", "\n\t\t")
	indented = strings.ReplaceAll(indented, ", INDENT", ", \n\t\t")

	return indented
}
