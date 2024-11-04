package docs

import (
	"github.com/maddalax/htmgo/framework/h"
)

func RelatedProjects(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Other languages and related projects"),
			Text(`
				If you're not a Go user but are interested in the idea of what htmgo is, you might want to check out these other projects:
			`),
			h.Ul(
				h.Class("font-bold"),
				h.Text("Python:"),
				h.Class("list-disc list-inside"),
				h.Li(
					h.P(
						h.Class("font-normal"),
						Link("fastht.ml", "https://fastht.ml"),
						h.Text(" - Modern web applications in pure Python, Built on solid web foundations, not the latest fads - with FastHTML you can get started on anything from simple dashboards to scalable web applications in minutes."),
					),
				),
			),
			NextStep(
				"mt-4",
				PrevBlock("Tailwind Intellisense", "/docs/misc/tailwind-intellisense"),
				NextBlock("Adding Interactivity", "/docs/interactivity/swapping"),
			),
		),
	)
}
