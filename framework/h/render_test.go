package h

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"sort"
	"strings"
	"testing"
)

// Sort attributes of a node by attribute name
func sortAttributes(node *html.Node) {
	if node.Type == html.ElementNode && len(node.Attr) > 1 {
		sort.SliceStable(node.Attr, func(i, j int) bool {
			return node.Attr[i].Key < node.Attr[j].Key
		})
	}
}

// Traverse and sort attributes in the entire HTML tree
func traverseAndSortAttributes(node *html.Node) {
	sortAttributes(node)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverseAndSortAttributes(child)
	}
}

// Parse HTML, sort attributes, and render back to a string
func parseSortAndRenderHTML(input string) string {
	// Parse the HTML string into a node tree
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return ""
	}

	// Traverse and sort attributes for each node
	traverseAndSortAttributes(doc)

	// Use a buffer to capture the rendered HTML
	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return ""
	}

	// Return the rendered HTML as a string
	return buf.String()
}

func TestRender(t *testing.T) {
	t.Parallel()
	div := Div(
		Id("my-div"),
		Attribute("data-attr-2", "value"),
		Attributes(&AttributeMap{
			"data-attr-3": "value",
			"data-attr-4": "value",
		}),
		HxBeforeRequest(
			SetText("before request"),
		),
		HxAfterRequest(
			SetText("after request"),
		),
		Children(
			Div(Text("hello, world")),
		),
		Text("hello, child"),
	)

	div.attributes["data-attr-1"] = "value"

	expectedRaw := `<div data-attr-1="value" id="my-div" data-attr-2="value" data-attr-3="value" data-attr-4="value" hx-on::before-request="this.innerText = 'before request';" hx-on::after-request="this.innerText = 'after request';"><div >hello, world</div>hello, child</div>`
	expected := parseSortAndRenderHTML(expectedRaw)
	result := parseSortAndRenderHTML(Render(div))

	assert.Equal(t,
		expected,
		result)
}

func TestRawContent(t *testing.T) {
	t.Parallel()
	str := "<div>hello, world</div>"
	raw := Raw(str)
	assert.Equal(t, str, Render(raw))
}

func TestConditional(t *testing.T) {
	t.Parallel()
	result := Render(
		Div(
			Ternary(true, Text("true"), Text("false")),
		),
	)
	assert.Equal(t, "<div >true</div>", result)

	result = Render(
		Div(
			If(false, Text("true")),
		),
	)
	assert.Equal(t, "<div ></div>", result)
}

func BenchmarkMailTo(b *testing.B) {
	b.ReportAllocs()
	ctx := RenderContext{
		builder: &strings.Builder{},
	}
	for i := 0; i < b.N; i++ {
		MailTo().Render(&ctx)
		ctx.builder.Reset()
	}
}

func BenchmarkComplexPage(b *testing.B) {
	b.Skip()
	b.ReportAllocs()
	ctx := RenderContext{
		builder: &strings.Builder{},
	}
	page := ComplexPage()
	for i := 0; i < b.N; i++ {
		page.Render(&ctx)
		ctx.builder.Reset()
	}
}

func ComplexPage() *Element {
	return Html(
		Head(
			Meta("title", "Complex Page"),
			Meta(
				"charset",
				"UTF-8",
			),
			Meta(
				"viewport",
				"width=device-width, initial-scale=1.0",
			),
			Link(
				"stylesheet",
				"https://example.com/styles.css",
			),
		),
		Body(
			Header(
				Class("bg-gray-800 text-white py-4"),
				Div(
					Class("container mx-auto"),
					H1(Class("text-3xl font-bold"), Text("Welcome to the Complex Page")),
					Nav(
						Ul(
							Class("flex space-x-4"),
							Li(A(Href("#"), Text("Home"))),
							Li(A(Href("#"), Text("About"))),
							Li(A(Href("#"), Text("Services"))),
							Li(A(Href("#"), Text("Contact"))),
						),
					),
				),
			),
			Main(
				Class("container mx-auto mt-10"),
				Section(
					Class("grid grid-cols-3 gap-4"),
					Article(
						Class("col-span-2"),
						H2(Class("text-2xl font-semibold mb-4"), Text("Featured Article")),
						Img(Src("https://example.com/featured.jpg"), Alt("Featured Image")),
						P(Class("mt-2 text-lg"), Text("This is a large article to test rendering performance.")),
					),
					Aside(
						Class("bg-gray-100 p-4"),
						H3(Class("text-xl font-bold"), Text("Related Links")),
						Ul(
							Li(A(Href("#"), Text("Related Link 1"))),
							Li(A(Href("#"), Text("Related Link 2"))),
							Li(A(Href("#"), Text("Related Link 3"))),
						),
					),
				),
				Section(
					Class("my-8"),
					H2(Class("text-2xl font-semibold mb-4"), Text("User Registration Form")),
					Form(
						Post("/register", "click"),
						Div(
							Class("grid grid-cols-2 gap-4"),
							Label(For("first_name"), Text("First Name")),
							Input("text", Id("first_name"), Name("first_name"), Class("border p-2 w-full")),
							Label(For("last_name"), Text("Last Name")),
							Input("text", Id("last_name"), Name("last_name"), Class("border p-2 w-full")),
							Label(For("email"), Text("Email")),
							Input("email", Id("email"), Name("email"), Class("border p-2 w-full")),
							Label(For("password"), Text("Password")),
							Input("password", Id("password"), Name("password"), Class("border p-2 w-full")),
						),
						Button(
							Type("submit"),
							Class("bg-blue-500 text-white py-2 px-4 mt-4"),
							Text("Register"),
						),
					),
				),
				Section(
					Class("my-8"),
					H2(Class("text-2xl font-semibold mb-4"), Text("Data Table")),
					Table(
						Class("table-auto w-full border-collapse border"),
						THead(
							Tr(
								Th(Class("border px-4 py-2"), Text("ID")),
								Th(Class("border px-4 py-2"), Text("Name")),
								Th(Class("border px-4 py-2"), Text("Age")),
								Th(Class("border px-4 py-2"), Text("Occupation")),
							),
						),
						TBody(
							Tr(
								Td(Class("border px-4 py-2"), Text("1")),
								Td(Class("border px-4 py-2"), Text("John Doe")),
								Td(Class("border px-4 py-2"), Text("28")),
								Td(Class("border px-4 py-2"), Text("Engineer")),
							),
							Tr(
								Td(Class("border px-4 py-2"), Text("2")),
								Td(Class("border px-4 py-2"), Text("Jane Smith")),
								Td(Class("border px-4 py-2"), Text("34")),
								Td(Class("border px-4 py-2"), Text("Designer")),
							),
							Tr(
								Td(Class("border px-4 py-2"), Text("3")),
								Td(Class("border px-4 py-2"), Text("Alice Johnson")),
								Td(Class("border px-4 py-2"), Text("45")),
								Td(Class("border px-4 py-2"), Text("Manager")),
							),
						),
					),
				),
			),
			Footer(
				Class("bg-gray-800 text-white py-4 mt-10"),
				Div(
					Class("container mx-auto text-center"),
					Text("© 2024 Complex Page Inc. All rights reserved."),
				),
			),
		),
	)
}

func MailTo() *Element {
	return Div(
		H1(
			Text("Contact Us"),
		),
		Div(
			Style("font-family: 'sans-serif'"),
			Id("test"),
			Attribute("data-contents", `something with "quotes" and a <tag>`),
			Div(
				Text("email:"),
				A(
					Href("mailto:"+"test@htmgo.dev"),
					Text("Email me"),
				),
			),
		),
		Hr(
			Attribute("noshade", ""),
		),
		Hr(
			Attribute("optionA", ""),
			Attribute("optionB", ""),
			Attribute("optionC", "other"),
		),
		Hr(
			Attribute("noshade", ""),
		),
	)
}
