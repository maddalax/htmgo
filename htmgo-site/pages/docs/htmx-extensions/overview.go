package htmx_extensions

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

func Overview(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("HTMX Extensions"),
			Text(`
				htmgo provides a few extra htmx extensions to make common tasks easier.
				Some of these extensions are optional, and some of these are required for htmgo to work correctly.
			`),
			Text(`
			 The following extensions are provided by htmgo:
			`),
			Link("Trigger Children", "/docs/htmx-extensions/trigger-children"),
			Link("Mutation Error", "/docs/htmx-extensions/mutation-error"),
			Link("Path Deps", "https://github.com/bigskysoftware/htmx-extensions/blob/main/src/path-deps/README.md"),
			h.P(
				h.Class("mt-3"),
				h.Text("Default extensions should be included in your project by adding the following attribute to your html tag."),
				ui.GoCodeSnippet(DefaultExtensions),
				h.Text("If you need to combine multiple extensions, you can use:"),
				ui.GoCodeSnippet(CombineMultipleExtensions),
				h.Text("or"),
				ui.GoCodeSnippet(CombineMultipleExtensions2),
			),
			Text(`
				<b>Important:</b> h.BaseExtensions will add the 'htmgo' extension, which is a required extension for inline scripts to work properly, please always include it in your project.
			`),
			NextStep(
				"mt-4",
				PrevBlock("Pushing Data", DocPath("/pushing-data/sse")),
				NextBlock("Trigger Children", DocPath("/htmx-extensions/trigger-children")),
			),
		),
	)
}

const DefaultExtensions = `
h.Html(
    h.HxExtension(h.BaseExtensions())
)
`

const CombineMultipleExtensions = `
h.HxExtensions(
	h.BaseExtensions(), "my-extension"
)
`

const CombineMultipleExtensions2 = `
h.JoinExtensions(
    h.HxExtension("sse"),
    h.HxExtension("my-extension"),
)
`
