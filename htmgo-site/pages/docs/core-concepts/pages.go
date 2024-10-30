package core_concepts

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

var ExcludeRootSnippet = `automatic_page_routing_ignore: ["pages/root.go"]`

var AbstractedRootPageUsageSnippet = `func UserPage(ctx *h.RequestContext) *h.Page {
	return base.RootPage(
		h.Div(
			h.Pf("User ID: %s", ctx.Param("id")),
		),
}`

var RootPageSnippet = `func RootPage(children ...h.Ren) *h.Page {
	return h.NewPage(
		h.Html(
		h.HxExtension(h.BaseExtensions()),
		h.Head(
			h.Meta("viewport", "width=device-width, initial-scale=1"),
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
		),
		h.Body(
			h.Class("bg-stone-50 min-h-screen overflow-x-hidden"),
			ui.NavBar(),
			h.Fragment(children...),
		),
	  )
   )
}
`

var PagesSnippet = `// route will be automatically registered based on the file name
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
}`

func Pages(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Pages"),
			Text(`
				Pages are the entry point of an htmgo application.
				A simple page may look like:
			`),
			ui.GoCodeSnippet(PagesSnippet),
			h.Text(`
				htmgo uses std http with chi router as its web server, *h.RequestContext is a thin wrapper around *http.Request. 
				A page must return *h.Page, and accept *h.RequestContext as a parameter
			`),
			autoRegistration(),
			tips(),
			NextStep(
				"mt-4",
				PrevBlock("Getting Started", DocPath("/installation")),
				NextBlock("Partials", DocPath("/core-concepts/partials")),
			),
		),
	)
}

func autoRegistration() *h.Element {
	return h.Div(
		h.Class("flex flex-col gap-2"),
		SubTitle("Auto Registration"),
		Text(`
		htmgo uses file based routing. This means that we will automatically generate and register your routes with chi based on the files you have in the 'pages' directory.
		For example, if you have a directory structure like so below, it will get registered into chi router as follows:

		index.go -> /index
		users.go -> /users
		users.$id.go -> /users/:id 
		`),
		HelpText(`Note: id parameter can be accessed in your page with ctx.Param("id")`),
		Text(`
		You may put any functions you like in your pages file, auto registration will ONLY register functions that return *h.Page
		`),
	)
}

func tips() *h.Element {
	return h.Div(
		h.Class("flex flex-col gap-2"),
		SubTitle("Tips:"),
		Text(`
			Generally it is it recommended to abstract common parts of your page into its own component and re-use it, such as script tags, including styling, etc.
			Example:
		`),
		ui.GoCodeSnippet(RootPageSnippet),
		Text("Usage:"),
		ui.GoCodeSnippet(AbstractedRootPageUsageSnippet),
		Text("You need to then update <strong>htmgo.yml</strong> to exclude that file from auto registration"),
		ui.SingleLineBashCodeSnippet(ExcludeRootSnippet),
		HelpText("In this example, my root page is in a file called root.go in the pages dir, so I need to exclude it from auto registration, otherwise htmgo wil try to generate a route for it."),
	)
}
