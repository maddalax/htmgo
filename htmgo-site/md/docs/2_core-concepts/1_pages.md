## Pages

Pages are the entry point of an htmgo application. 

A simple page may look like:

```go
// route will be automatically registered based on the file name
func HelloHtmgoPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		h.Html(
			h.HxExtension(h.BaseExtensions()),
			h.Head(
				h.Link("/public/main.css", "stylesheet"),
				h.Script("/public/htmgo.js"),
			),
			h.Body(
				h.Pf("Hello, htmgo!"),
			),
		),
	)
}
```

htmgo uses [std http](https://pkg.go.dev/net/http) with chi router as its web server, ***h.RequestContext** is a thin wrapper around ***http.Request**. A page
must return *h.Page, and accept *h.RequestContext as a parameter

<br>

**Auto Registration**

htmgo uses file based routing. This means that we will automatically generate and register your routes with chi based on the files you have in the 'pages' directory.

For example, if you have a directory structure such as:

```bash
pages
  index.go
  users.go
  users.$id //id parameter can be accessed in your page with ctx.Param("id")
```

it will get registered into chi router as follows:

```bash
/
/users
/users/:id
```

You may put any functions you like in your pages file, auto registration will **ONLY** register functions that return ***h.Page**

<br>

**Tips:**

Generally it is it recommended to abstract common parts of your page into its own component and re-use it, such as script tags, including styling, etc.

Example:

```go
func RootPage(children ...h.Ren) *h.Element {
	return h.Html(
		h.HxExtension(h.BaseExtensions()),
		h.Head(
			h.Meta("viewport", "width=device-width, initial-scale=1"),
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
			h.Style(`
				html {
					scroll-behavior: smooth;
				}
			`),
		),
		h.Body(
			h.Class("bg-stone-50 min-h-screen overflow-x-hidden"),
			partials.NavBar(false),
			h.Fragment(children...),
		),
	)
}
```

```go
func UserPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		base.RootPage(
			h.Div(
				h.Pf("User ID: %s", ctx.Param("id")),
			),
		))
}
```

