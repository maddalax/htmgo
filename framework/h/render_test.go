package h

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
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
func sortHtmlAttributes(input string) string {
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
	expected := sortHtmlAttributes(expectedRaw)
	result := sortHtmlAttributes(Render(div))

	assert.Equal(t,
		expected,
		result)
}

func TestRawContent(t *testing.T) {
	t.Parallel()
	str := "<div>hello, world</div>"
	raw := UnsafeRaw(str)
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

func TestTagSelfClosing(t *testing.T) {
	t.Parallel()
	assert.Equal(t, `<input type="text"/>`, Render(
		Input("text"),
	))
	// assert the tag cannot have children
	assert.Equal(t, `<input type="text"/>`, Render(
		Input("text", Div()),
	))

	assert.Equal(t, `<div id="test"></div>`, Render(
		Div(Id("test")),
	))
	assert.Equal(t, `<div id="test"><div ></div></div>`, Render(
		Div(Id("test"), Div()),
	))
}

func TestCached(t *testing.T) {
	t.Parallel()
	count := 0
	page := Cached(time.Hour, func() *Element {
		count++
		return ComplexPage()
	})

	firstRender := sortHtmlAttributes(Render(page()))
	secondRender := sortHtmlAttributes(Render(page()))

	assert.Equal(t, firstRender, secondRender)
	assert.Equal(t, 1, count)
	assert.Equal(t, firstRender, sortHtmlAttributes(Render(ComplexPage())))
}

func TestCachedT(t *testing.T) {
	t.Parallel()
	count := 0
	page := CachedT(time.Hour, func(a string) *Element {
		count++
		return ComplexPage()
	})

	firstRender := sortHtmlAttributes(Render(page("a")))
	secondRender := sortHtmlAttributes(Render(page("a")))

	assert.Equal(t, firstRender, secondRender)
	assert.Equal(t, 1, count)
	assert.Equal(t, firstRender, sortHtmlAttributes(Render(ComplexPage())))
}

func TestCachedT2(t *testing.T) {
	t.Parallel()
	count := 0
	page := CachedT2(time.Hour, func(a string, b string) *Element {
		count++
		return ComplexPage()
	})

	firstRender := sortHtmlAttributes(Render(page("a", "b")))
	secondRender := sortHtmlAttributes(Render(page("a", "b")))

	assert.Equal(t, firstRender, secondRender)
	assert.Equal(t, 1, count)
	assert.Equal(t, firstRender, sortHtmlAttributes(Render(ComplexPage())))
}

func TestCachedT3(t *testing.T) {
	t.Parallel()
	count := 0
	page := CachedT3(time.Hour, func(a string, b string, c string) *Element {
		count++
		return ComplexPage()
	})

	firstRender := sortHtmlAttributes(Render(page("a", "b", "c")))
	secondRender := sortHtmlAttributes(Render(page("a", "b", "c")))

	assert.Equal(t, firstRender, secondRender)
	assert.Equal(t, 1, count)
	assert.Equal(t, firstRender, sortHtmlAttributes(Render(ComplexPage())))
}

func TestCachedT4(t *testing.T) {
	t.Parallel()
	count := 0
	page := CachedT4(time.Hour, func(a string, b string, c string, d string) *Element {
		count++
		return ComplexPage()
	})

	firstRender := sortHtmlAttributes(Render(page("a", "b", "c", "d")))
	secondRender := sortHtmlAttributes(Render(page("a", "b", "c", "d")))

	assert.Equal(t, firstRender, secondRender)
	assert.Equal(t, 1, count)
	assert.Equal(t, firstRender, sortHtmlAttributes(Render(ComplexPage())))
}

func TestCachedExpired(t *testing.T) {
	t.Parallel()
	count := 0
	page := Cached(time.Millisecond*3, func() *Element {
		count++
		return ComplexPage()
	})

	firstRender := sortHtmlAttributes(Render(page()))
	time.Sleep(time.Millisecond * 5)
	secondRender := sortHtmlAttributes(Render(page()))

	assert.Equal(t, firstRender, secondRender)
	assert.Equal(t, 2, count)
}

func TestCacheMultiple(t *testing.T) {
	t.Parallel()
	count := 0
	cachedItem := Cached(time.Hour, func() *Element {
		count++
		return Div(Text("hello"))
	})

	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Render(Div(
				cachedItem(),
				cachedItem(),
				cachedItem(),
			))
		}()
	}

	wg.Wait()

	assert.Equal(t, 1, count)
}

func TestCacheByKey(t *testing.T) {
	t.Parallel()
	renderCount := 0
	callCount := 0
	cachedItem := CachedPerKey(time.Hour, func() (any, GetElementFunc) {
		key := "key"
		if callCount == 3 {
			key = "key2"
		}
		if callCount == 4 {
			key = "key"
		}
		callCount++
		return key, func() *Element {
			renderCount++
			return Div(Text("hello"))
		}
	})

	Render(Div(
		cachedItem(),
		cachedItem(),
		cachedItem(),
		cachedItem(),
		cachedItem(),
	))

	assert.Equal(t, 5, callCount)
	assert.Equal(t, 2, renderCount)
}

func TestCacheByKeyT(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := CachedPerKeyT(time.Hour, func(key string) (any, GetElementFunc) {
		return key, func() *Element {
			renderCount++
			return Div(Text("hello"))
		}
	})

	Render(Div(
		cachedItem("one"),
		cachedItem("one"),
		cachedItem("two"),
		cachedItem("two"),
		cachedItem("three"),
	))

	assert.Equal(t, 3, renderCount)
}

func TestCacheByKeyT4(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := CachedPerKeyT4(time.Hour, func(arg1 string, arg2 string, arg3 string, arg4 string) (any, GetElementFunc) {
		return arg1, func() *Element {
			renderCount++
			return Div(Text("hello"))
		}
	})

	Render(Div(
		cachedItem("one", uuid.NewString(), uuid.NewString(), uuid.NewString()),
		cachedItem("one", uuid.NewString(), uuid.NewString(), uuid.NewString()),
		cachedItem("two", uuid.NewString(), uuid.NewString(), uuid.NewString()),
		cachedItem("two", uuid.NewString(), uuid.NewString(), uuid.NewString()),
		cachedItem("three", uuid.NewString(), uuid.NewString(), uuid.NewString()),
	))

	assert.Equal(t, 3, renderCount)
}

func TestCacheByKeyT3(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := CachedPerKeyT3(time.Hour, func(arg1 string, arg2 string, arg3 string) (any, GetElementFunc) {
		return arg1, func() *Element {
			renderCount++
			return Div(Text("hello"))
		}
	})

	Render(Div(
		cachedItem("one", uuid.NewString(), uuid.NewString()),
		cachedItem("one", uuid.NewString(), uuid.NewString()),
		cachedItem("two", uuid.NewString(), uuid.NewString()),
		cachedItem("two", uuid.NewString(), uuid.NewString()),
		cachedItem("three", uuid.NewString(), uuid.NewString()),
	))

	assert.Equal(t, 3, renderCount)
}

func TestCacheByKeyT2(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := CachedPerKeyT2(time.Hour, func(arg1 string, arg2 string) (any, GetElementFunc) {
		return arg1, func() *Element {
			renderCount++
			return Div(Text("hello"))
		}
	})

	Render(Div(
		cachedItem("one", uuid.NewString()),
		cachedItem("one", uuid.NewString()),
		cachedItem("two", uuid.NewString()),
		cachedItem("two", uuid.NewString()),
		cachedItem("three", uuid.NewString()),
	))

	assert.Equal(t, 3, renderCount)
}

func TestCacheByKeyConcurrent(t *testing.T) {
	t.Parallel()
	renderCount := 0
	callCount := 0
	cachedItem := CachedPerKey(time.Hour, func() (any, GetElementFunc) {
		key := "key"
		if callCount == 3 {
			key = "key2"
		}
		if callCount == 4 {
			key = "key"
		}
		callCount++
		return key, func() *Element {
			renderCount++
			return Div(Text("hello"))
		}
	})

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Render(Div(
				cachedItem(),
			))
		}()
	}

	wg.Wait()

	assert.Equal(t, 5, callCount)
	assert.Equal(t, 2, renderCount)
}

func TestCacheByKeyT1_2(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := CachedPerKeyT(time.Hour, func(key string) (any, GetElementFunc) {
		return key, func() *Element {
			renderCount++
			return Pf(key)
		}
	})

	assert.Equal(t, "<p >one</p>", Render(cachedItem("one")))
	assert.Equal(t, "<p >two</p>", Render(cachedItem("two")))
	assert.Equal(t, "<p >two</p>", Render(cachedItem("two")))
	assert.Equal(t, 2, renderCount)
}

func TestCacheByKeyT1Expired(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := CachedPerKeyT(time.Millisecond, func(key string) (any, GetElementFunc) {
		return key, func() *Element {
			renderCount++
			return Pf(key)
		}
	})

	assert.Equal(t, "<p >one</p>", Render(cachedItem("one")))
	assert.Equal(t, "<p >two</p>", Render(cachedItem("two")))
	time.Sleep(time.Millisecond * 2)
	assert.Equal(t, "<p >two</p>", Render(cachedItem("two")))
	assert.Equal(t, 3, renderCount)
}

func TestCacheByKeyT1Expired_2(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := CachedPerKeyT(time.Millisecond*5, func(key string) (any, GetElementFunc) {
		return key, func() *Element {
			renderCount++
			return Pf(key)
		}
	})

	assert.Equal(t, "<p >one</p>", Render(cachedItem("one")))
	time.Sleep(time.Millisecond * 3)
	assert.Equal(t, "<p >two</p>", Render(cachedItem("two")))
	assert.Equal(t, "<p >two</p>", Render(cachedItem("two")))
	assert.Equal(t, "<p >two</p>", Render(cachedItem("two")))
	time.Sleep(time.Millisecond * 3)
	assert.Equal(t, "<p >one</p>", Render(cachedItem("one")))
	assert.Equal(t, "<p >two</p>", Render(cachedItem("two")))
	assert.Equal(t, 3, renderCount)
}

func TestClearExpiredCached(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := Cached(time.Millisecond*3, func() *Element {
		renderCount++
		return Pf("hello")
	})

	Render(cachedItem())
	Render(cachedItem())
	node := cachedItem().meta.(*CachedNode)
	assert.Equal(t, 1, renderCount)
	assert.NotEmpty(t, node.html)

	time.Sleep(time.Millisecond * 3)
	node.ClearExpired()

	assert.Empty(t, node.html)
}

func TestClearExpiredCacheByKey(t *testing.T) {
	t.Parallel()
	renderCount := 0
	cachedItem := CachedPerKeyT(time.Millisecond, func(key int) (any, GetElementFunc) {
		return key, func() *Element {
			renderCount++
			return Pf(strconv.Itoa(key))
		}
	})

	for i := 0; i < 100; i++ {
		Render(cachedItem(i))
	}

	node := cachedItem(0).meta.(*ByKeyEntry).parent.meta.(*CachedNode)
	assert.Equal(t, 100, len(node.byKeyExpiration))
	assert.Equal(t, 100, len(node.byKeyCache))

	time.Sleep(time.Millisecond * 2)

	Render(cachedItem(0))
	node.ClearExpired()

	assert.Equal(t, 1, len(node.byKeyExpiration))
	assert.Equal(t, 1, len(node.byKeyCache))

	node.ClearCache()

	assert.Equal(t, 0, len(node.byKeyExpiration))
	assert.Equal(t, 0, len(node.byKeyCache))
}

func TestBackgroundCleaner(t *testing.T) {
	t.Parallel()
	cachedItem := CachedPerKeyT(time.Second*2, func(key int) (any, GetElementFunc) {
		return key, func() *Element {
			return Pf(strconv.Itoa(key))
		}
	})
	for i := 0; i < 100; i++ {
		Render(cachedItem(i))
	}

	node := cachedItem(0).meta.(*ByKeyEntry).parent.meta.(*CachedNode)
	assert.Equal(t, 100, len(node.byKeyExpiration))
	assert.Equal(t, 100, len(node.byKeyCache))

	time.Sleep(time.Second * 3)

	assert.Equal(t, 0, len(node.byKeyExpiration))
	assert.Equal(t, 0, len(node.byKeyCache))
}

func TestEscapeHtml(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "&lt;script&gt;alert(1)&lt;/script&gt;", Render(Text("<script>alert(1)</script>")))
	assert.Equal(t, "<p >&lt;script&gt;alert(1)&lt;/script&gt;</p>", Render(Pf("<script>alert(1)</script>")))

}

func BenchmarkCacheByKey(b *testing.B) {
	b.ReportAllocs()
	page := CachedPerKeyT(time.Second*3, func(userId string) (any, GetElementFunc) {
		return userId, func() *Element {
			return MailTo(userId)
		}
	})

	for i := 0; i < 5000; i++ {
		userId := uuid.NewString()
		Render(page(userId))
	}

	Render(page(uuid.NewString()))
}

func BenchmarkMailToStatic(b *testing.B) {
	b.ReportAllocs()
	ctx := RenderContext{
		builder: &strings.Builder{},
	}
	page := MailTo("myemail")
	for i := 0; i < b.N; i++ {
		page.Render(&ctx)
		ctx.builder.Reset()
	}
}

func BenchmarkMailToDynamic(b *testing.B) {
	b.ReportAllocs()
	ctx := RenderContext{
		builder: &strings.Builder{},
	}
	for i := 0; i < b.N; i++ {
		MailTo(uuid.NewString()).Render(&ctx)
		ctx.builder.Reset()
	}
}

func BenchmarkCachedComplexPage(b *testing.B) {
	b.ReportAllocs()
	ctx := RenderContext{
		builder: &strings.Builder{},
	}
	for i := 0; i < b.N; i++ {
		CachedComplexPage().Render(&ctx)
		ctx.builder.Reset()
	}
}

func BenchmarkComplexPage(b *testing.B) {
	b.ReportAllocs()
	ctx := RenderContext{
		builder: &strings.Builder{},
	}
	for i := 0; i < b.N; i++ {
		ComplexPage().Render(&ctx)
		ctx.builder.Reset()
	}
}

var CachedComplexPage = Cached(time.Hour, func() *Element {
	return ComplexPage()
})

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

func MailTo(email string) *Element {
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
					Href(email),
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
