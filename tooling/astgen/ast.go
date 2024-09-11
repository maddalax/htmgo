package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
)

type AstUtil struct {
	file *ast.File
}

func NewAstUtil(file *ast.File) *AstUtil {
	return &AstUtil{file: file}
}

func (a *AstUtil) Sync(builder *CodeBuilder) {
	fset := token.NewFileSet()
	// Parse the code string into an AST
	f, err := parser.ParseFile(fset, "", builder.String(), 0)
	if err != nil {
		panic(err)
	}
	a.file = f
}

func (a *AstUtil) HasMethod(methodName string) bool {
	return a.GetMethodByName(methodName) != nil
}

func (a *AstUtil) HasStruct(name string) bool {
	for _, decl := range a.file.Decls {
		_, ok := decl.(*ast.GenDecl)
		if ok {
			spec, ok := decl.(*ast.GenDecl).Specs[0].(*ast.TypeSpec)
			if ok {
				if spec.Name.Name == name {
					return true
				}
			}
		}
	}
	return false
}

func (a *AstUtil) GetMethodByName(name string) *ast.FuncDecl {
	for _, decl := range a.file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == name {
			return funcDecl
		}
	}
	return nil
}

func (a *AstUtil) GetMethodAsString(methodName string) string {
	fdecl := a.GetMethodByName(methodName)

	if fdecl == nil {
		return ""
	}

	var buf bytes.Buffer
	fset := token.NewFileSet() // Create a FileSet to manage source file positions

	// Use printer.Fprint to format the AST node into the buffer
	err := printer.Fprint(&buf, fset, fdecl)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

type UpdateSignature struct {
	MethodName string
	NewName    string
	NewParams  []NameType
	NewResults []ReturnType
}

func (a *AstUtil) UpdateMethodSignature(req UpdateSignature) {

	newParams := Map(req.NewParams, func(nt NameType) *ast.Field {
		return nt.ToAst()
	})

	newResults := Map(req.NewResults, func(rt ReturnType) *ast.Field {
		return rt.ToAst()
	})

	for _, decl := range a.file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == req.MethodName {
			// Method found, update its signature
			funcDecl.Name = ast.NewIdent(req.NewName)
			funcDecl.Type.Params = &ast.FieldList{List: newParams}
			funcDecl.Type.Results = &ast.FieldList{List: newResults}
		}
	}
}

func (a *AstUtil) SetPackageName(name string) {
	a.file.Name = ast.NewIdent(name)
}

func (a *AstUtil) AddImport(path string) {
	if !astutil.UsesImport(a.file, path) {
		astutil.AddImport(token.NewFileSet(), a.file, path)
	}
}

func (a *AstUtil) DeleteMethod(methodName string) {
	for i, decl := range a.file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == methodName {
			// Method found, remove it from the slice
			a.file.Decls = append(a.file.Decls[:i], a.file.Decls[i+1:]...)
			return
		}
	}
}

func (a *AstUtil) String() string {
	var buf bytes.Buffer
	fset := token.NewFileSet()
	err := printer.Fprint(&buf, fset, a.file)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
