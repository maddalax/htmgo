package core_concepts

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/partials"
	"htmgo-site/ui"
)

var PartialsSnippet = `func CurrentTimePartial(ctx *h.RequestContext) *h.Partial {
	now := time.Now()
	return h.NewPartial(
		h.Div(
			h.Pf("The current time is %s", now.Format(time.RFC3339)),
		),
	)
}`

var examplePageSnippet = `func CurrentTimePage(ctx *h.RequestContext) *h.Page {
	return RootPage(
		h.GetPartial(partials.CurrentTimePartial, "load, every 1s")
	)
}`

var examplePartialSnippet = `func CurrentTimePartial(ctx *h.RequestContext) *h.Partial {
	now := time.Now()
	return h.NewPartial(
		h.Div(
			h.Pf("The current time is %s", now.Format(time.RFC3339)),
		),
	)
}`

func Partials(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Partials"),
			Text(`
				Partials are where things get interesting.
				Partials allow you to start adding interactivity to your website by swapping in content, setting headers, redirecting, etc.
				Partials have a similar structure to pages. A simple partial may look like:
			`),
			ui.GoCodeSnippet(PartialsSnippet),
			h.Text(`
				This will get automatically registered in the same way that pages are registered, based on the file path. 
				This allows you to reference partials directly via the function itself when rendering them, instead of worrying about the route.
			`),
			example(),
			NextStep(
				"mt-4",
				PrevBlock("Pages", DocPath("/core-concepts/pages")),
				NextBlock("Components", DocPath("/core-concepts/components")),
			),
		),
	)
}

func example() *h.Element {
	return h.Div(
		h.Class("flex flex-col gap-2"),
		SubTitle("Simple Example"),
		Text(`
			I want to build a page that renders the current time, updating every second. Here is how that may look:
		`),
		h.Pf(
			"pages/time.go",
			h.Class("font-semibold"),
		),
		ui.GoCodeSnippet(examplePageSnippet),
		h.Pf(
			"partials/time.go",
			h.Class("font-semibold"),
		),
		ui.GoCodeSnippet(examplePartialSnippet),
		Text(
			`When the page load, the partial will be loaded in via htmx, and then swapped in every 1 second. 
		With this little amount of code and zero written javascript, you have a page that shows the current time and updates every second.`),
		h.Div(
			h.GetPartial(partials.CurrentTimePartial, "load, every 1s"),
		),
	)
}
