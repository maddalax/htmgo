package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"strings"
	"text/template"
)

// ... (Field and Method structs remain the same)

type CodeBuilder struct {
	structTemplate        *template.Template
	methodTemplate        *template.Template
	toEventTemplate       *template.Template
	commandSwitchTemplate *template.Template
	validateTemplate      *template.Template
	funcTemplate          *template.Template
	interfaceTemplate     *template.Template
	result                string
	astu                  *AstUtil
}

type Assignment struct {
	Field    string
	Value    string
	Operator string
}

type Function struct {
	Name       string
	Parameters []NameType
	Return     []ReturnType
	Body       string
}

type Struct struct {
	Name   string
	Fields []NameType
}

type Validation struct {
	MethodCall string
}

type Method struct {
	StructReceiver string
	StructName     string
	MethodName     string
	Body           string
	Parameters     []NameType
	Return         []ReturnType
}

type NameType struct {
	Name string
	Type string
}

func (nt *NameType) ToAst() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{ast.NewIdent(nt.Name)}, // Create an identifier for the field name
		Type:  ast.NewIdent(nt.Type),               // Create an identifier for the field type
	}
}

type ReturnType struct {
	Type string
}

func (rt *ReturnType) ToAst() *ast.Field {
	return &ast.Field{
		Type: ast.NewIdent(rt.Type),
	}
}

type Case struct {
	CommandType string
	CaseBody    string
}

func (cb *CodeBuilder) Append(result string) *CodeBuilder {
	cb.result += "\n"
	cb.result += result
	cb.result += "\n"
	return cb
}

func (cb *CodeBuilder) String() string {
	return cb.result
}

func (cb *CodeBuilder) SyncFromAst() *CodeBuilder {
	cb.SetResult(cb.astu.String())
	return cb
}

func (cb *CodeBuilder) SyncToAst() *CodeBuilder {
	cb.astu.Sync(cb)
	return cb
}

func (cb *CodeBuilder) SetResult(result string) *CodeBuilder {
	cb.result = result
	return cb
}

func NewCodeBuilder(astu *AstUtil) *CodeBuilder {
	toEventTemplate, _ := template.New("toEvent").Parse(`
	func (c *{{.StructName}}) {{.MethodCall}}() *{{.EventType}} {
		var event *{{.EventType}} = &{{.EventType}}{}
		{{range .Fields}}
		event.{{.Name}} = c.{{.Name}}
		{{end}}
		return event
	}
	`)

	validationTemplate := template.Must(template.New("validation").Parse(`
	func (c *{{.StructName}}) {{.MethodCall}}() error {
		{{range .Validations}}if err := {{.MethodCall}}; err != nil {
			return err
		}
		{{end}}
		return nil
	}
	`))

	methodTemplate := template.Must(template.New("method").Parse(`
	func ({{.StructReceiver}} *{{.StructName}}) {{.MethodCall}}({{range $index, $param := .Parameters}}{{if $index}}, {{end}}{{$param.Name}} {{$param.Type}}{{end}}) ({{range $index, $ret := .Return}}{{if $index}}, {{end}}{{$ret.Type}}{{end}}) {
		{{if .Body}}
		{{.Body}}
		{{else}}
		return {{range $index, $ret := .Return}}{{if $index}}, {{end}}nil{{end}}
		{{end}}
	}   
	`))

	commandSwitchTemplate, _ := template.New("commandSwitch").Parse(`
	switch c := command.(type) {
	{{range .Cases}}case *{{.CommandType}}:
		{{.CaseBody}}
	{{end}}
	}
	`)

	structTemplate, _ := template.New("struct").Parse(`
	type {{.StructName}} struct {
		{{range .Fields}}
		{{.Name}} {{.Type}}
		{{end}}
	}
	`)

	interfaceTemplate := template.Must(template.New("interface").Parse(`
	type {{.Name}} interface {
		{{range .Methods}}
		{{.Name}}({{range $index, $param := .Parameters}}{{if $index}}, {{end}}{{$param.Name}} {{$param.Type}}{{end}}) {{range $index, $ret := .Return}}{{if $index}}, {{end}}{{$ret.Type}}{{end}}
		{{end}}
	}
	`))

	funcTemplateStr := `
	func {{ .Name }}({{ range $index, $param := .Parameters }}{{if $index}}, {{end}}{{$param.Name}} {{$param.Type}}{{end}}) ({{range $index, $ret := .Return }}{{if $index}}, {{end}}{{$ret.Type}}{{end}}) {
       {{.Body}}
	}
	`

	funcTemplate := template.Must(template.New("function").Funcs(template.FuncMap{
		"trimPrefix": strings.TrimPrefix,
	}).Parse(funcTemplateStr))

	builder := &CodeBuilder{
		funcTemplate:          funcTemplate,
		structTemplate:        structTemplate,
		interfaceTemplate:     interfaceTemplate,
		toEventTemplate:       toEventTemplate,
		methodTemplate:        methodTemplate,
		validateTemplate:      validationTemplate,
		commandSwitchTemplate: commandSwitchTemplate,
		astu:                  astu,
	}

	if astu != nil {
		builder.SyncFromAst()
	}

	return builder
}

func (cb *CodeBuilder) PrependLine(line string) {
	cb.result = line + "\n" + cb.result
}

func (cb *CodeBuilder) PrependLineF(format string, args ...interface{}) {
	cb.PrependLine(fmt.Sprintf(format, args...))
}

func (cb *CodeBuilder) AppendLine(line string) {
	cb.result += line + "\n"
}

func (cb *CodeBuilder) PrependLineIfNotExist(format string) {
	if !strings.Contains(cb.result, format) {
		cb.PrependLine(format)
	}
}

func (cb *CodeBuilder) AppendLineIfNotExist(line string) {
	if !strings.Contains(cb.result, line) {
		cb.result += line + "\n"
	}
}

func (cb *CodeBuilder) HasString(str string) bool {
	return strings.Contains(cb.result, str)
}

func (cb *CodeBuilder) HasStringF(str string, args ...interface{}) bool {
	return strings.Contains(cb.result, fmt.Sprintf(str, args...))
}

func (cb *CodeBuilder) AppendLineF(format string, args ...interface{}) {
	cb.AppendLine(fmt.Sprintf(format, args...))
}

func (cb *CodeBuilder) ExecuteTemplate(template *template.Template, data map[string]any) string {
	var buf bytes.Buffer
	err := template.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return removeAllEmptyLines(buf.String())
}

func (cb *CodeBuilder) AddImport(imp string) *CodeBuilder {
	line := fmt.Sprintf(`import "%s"`, imp)
	if !strings.Contains(cb.result, line) {
		cb.AppendLine(line)
	}
	return cb
}

func (cb *CodeBuilder) BuildAssignment(a Assignment) string {
	return a.Field + " " + a.Operator + " " + a.Value
}

func (cb *CodeBuilder) BuildInterface(name string, methods []Function) string {
	return cb.ExecuteTemplate(cb.interfaceTemplate, map[string]any{
		"Name":    name,
		"Methods": methods,
	})
}

func (cb *CodeBuilder) BuildNewCommandMethod(_struct Struct) string {
	return cb.BuildFunction(Function{
		Name:       fmt.Sprintf("New%s", _struct.Name),
		Parameters: _struct.Fields,
		Return: []ReturnType{
			{Type: fmt.Sprintf("*%s", _struct.Name)},
		},
		Body: cb.Template(`
			result := {{.ReturnType}}{} 
    		{{range .Fields}}result.{{.Name}} = {{.Name}}
    		{{end}}{{if .Body}}
    		{{.Body}}{{end}}
    		return result
		`, map[string]any{
			"Fields":     _struct.Fields,
			"ReturnType": "&" + _struct.Name,
		}),
	})
}

func (cb *CodeBuilder) BuildValidation(structName string, methodName string, validations []Validation) string {
	return cb.ExecuteTemplate(cb.validateTemplate, map[string]any{
		"StructName":  structName,
		"MethodCall":  methodName,
		"Validations": validations,
	})
}

func (cb *CodeBuilder) BuildMethod(method Method) string {
	data := map[string]any{
		"StructReceiver": method.StructReceiver,
		"StructName":     method.StructName,
		"MethodCall":     method.MethodName,
		"Parameters":     method.Parameters,
		"Return":         method.Return,
	}
	if method.Body != "" {
		data["Body"] = method.Body
	}
	return cb.ExecuteTemplate(cb.methodTemplate, data)
}

func (cb *CodeBuilder) Template(templateStr string, data map[string]any) string {
	t, err := template.New("temp").Parse(templateStr)
	if err != nil {
		panic(err)
	}
	return cb.ExecuteTemplate(t, data)
}

func (cb *CodeBuilder) BuildStruct(s Struct) string {
	return cb.ExecuteTemplate(cb.structTemplate, map[string]any{
		"StructName": s.Name,
		"Fields":     s.Fields,
	})
}

func (cb *CodeBuilder) BuildFunction(f Function) string {
	return cb.ExecuteTemplate(cb.funcTemplate, map[string]any{
		"Name":       f.Name,
		"Parameters": f.Parameters,
		"Return":     f.Return,
		"Body":       f.Body,
	})
}

func (cb *CodeBuilder) BuildCommandSwitch(cases []Case) (string, error) {
	var buf bytes.Buffer
	err := cb.commandSwitchTemplate.Execute(&buf, map[string]interface{}{
		"Cases": cases,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func removeAllEmptyLines(input string) string {
	// Split the input into lines
	lines := strings.Split(input, "\n")

	// Filter out empty lines
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	// Join the non-empty lines back into a string
	output := strings.Join(nonEmptyLines, "\n")

	return output
}

func Map[T any, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

func FlatMap[T any, U any](slice []T, transform func(T) []U) []U {
	var result []U
	for _, v := range slice {
		result = append(result, transform(v)...)
	}
	return result
}
