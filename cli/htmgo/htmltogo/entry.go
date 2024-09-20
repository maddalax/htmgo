package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/dave/jennifer/jen"
	"golang.org/x/net/html"
)

func main() {
	// Example HTML input
	htmlData := `
		<body><nav class="flex gap-4 items-center p-4 text-slate-600 "><a href="/" class="cursor-pointer hover:text-blue-400 ">Home</a><a class="cursor-pointer hover:text-blue-400 " href="/news">News</a><a href="/patients" class="cursor-pointer hover:text-blue-400 ">Patients</a></nav><div id="active-modal"></div><div class="flex flex-col gap-2 bg-white h-full "><div class="flex flex-col p-4 w-full "><div><div class="flex justify-between items-center "><p class="text-lg font-bold ">Manage Patients</p><button hx-target="#active-modal" type="button" id="add-patient" class="flex gap-1 items-center border p-4 rounded cursor-hover bg-blue-700 text-white rounded p-2 h-12 " hx-get="htmgo/partials/patient.AddPatientSheet">Add Patient</button></div><div hx-get="htmgo/partials/patient.List" hx-trigger="load, patient-added from:body" class=""><div class="mt-8" id="patient-list"><div class="flex flex-col gap-2 rounded p-4 bg-red-100 "><p>Name: Sydne</p><p>Reason for visit: arm hurts</p></div></div></div></div></div></div><div hx-get="/livereload" hx-trigger="every 200ms" class=""></div></body>
	`

	// Parse the HTML
	doc, err := html.Parse(bytes.NewReader([]byte(htmlData)))
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Jennifer file
	f := jen.NewFile("main")

	// Generate Jennifer code for the parsed HTML tree
	generatedCode := processNode(doc.FirstChild)

	// Add the generated code to the file
	f.Func().Id("Render").Params().Block(generatedCode...)

	// Render the generated code
	var buf bytes.Buffer
	err = f.Render(&buf)
	if err != nil {
		log.Fatal(err)
	}

	//// Format the generated code
	//formattedCode, err := format.Source(buf.Bytes())
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Output the formatted code
	fmt.Println(string(buf.Bytes()))
}

// Recursively process the HTML nodes and generate Jennifer code
func processNode(n *html.Node) []jen.Code {
	var code []jen.Code

	// Only process element nodes
	if n.Type == html.ElementNode {
		// Create a dynamic method call based on the tag name
		tagMethod := strings.Title(n.Data) // Capitalize the first letter of the tag

		// Add dynamic method call for the tag (e.g., h.Div(), h.Button(), etc.)
		code = append(code, jen.Id("h").Dot(tagMethod).Call(mergeArgs(n)...))
	}

	return code
}

// Merge attributes and children into a single slice for Call()
func mergeArgs(n *html.Node) []jen.Code {
	// Process attributes
	attrs := processAttributes(n.Attr)

	// Process children
	children := processChildren(n)

	// Combine attributes and children into one slice
	return append(attrs, children...)
}

// Process child nodes of a given HTML node
func processChildren(n *html.Node) []jen.Code {
	var children []jen.Code

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, processNode(c)...)
	}

	return children
}

func FormatFieldName(name string) string {
	split := strings.Split(name, "_")
	if strings.Contains(name, "-") {
		split = strings.Split(name, "-")
	}
	parts := make([]string, 0)
	for _, s := range split {
		parts = append(parts, PascalCase(s))
	}
	return strings.Join(parts, "")
}

func PascalCase(s string) string {
	if s == "" {
		return s
	}
	// Convert the first rune (character) to uppercase and concatenate with the rest of the string
	return strings.ToUpper(string(s[0])) + s[1:]
}

// Process the attributes of an HTML node and return Jennifer code
func processAttributes(attrs []html.Attribute) []jen.Code {
	var args []jen.Code
	for _, attr := range attrs {
		// Dynamically handle all attributes
		attrMethod := FormatFieldName(attr.Key) // E.g., convert "data-role" to "DataRole"
		args = append(args, jen.Id("h").Dot(attrMethod).Call(jen.Lit(attr.Val)))
	}
	return args
}
