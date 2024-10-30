package docs2

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/ui"
)

func Installation(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Getting Started"),
			h.Ul(
				h.Text("Prerequisites:"),
				h.Class("list-disc list-outside"),
				h.Li(
					Inline(
						Link("Go 1.23 or above", "https://go.dev/doc/install"),
					),
				),
				Inline(
					Text("Familiarity with "),
					Link("https://htmx.org", "https://htmx.org"),
					Text(" and html/hypermedia"),
				),
			),
			HelpText("If you have not read the htmx docs, please do so before continuing, a lot of concepts below will be much more clear after."),
			StepTitle("1. Install htmgo"),
			ui.SingleLineBashCodeSnippet(`GOPROXY=direct go install github.com/maddalax/htmgo/cli/htmgo@latest`),
			StepTitle("2. Create new project"),
			ui.SingleLineBashCodeSnippet(`htmgo template myapp`),
			HelpText("this will ask you for a new app name, and it will clone our starter template to a new directory it creates with your app name."),
			StepTitle("3. Running the dev server"),
			ui.SingleLineBashCodeSnippet(`htmgo watch`),
			HelpText("htmgo has built in live reload on the dev server, to use this, run this command in the root of your project"),
			HelpText("If you prefer to run the dev server yourself (no live reload), use `htmgo run`"),
			StepTitle("4. Building for production"),
			ui.SingleLineBashCodeSnippet(`htmgo build`),
			HelpText("it will be output to `./dist`"),
			NextStep(
				"mt-4",
				PrevBlock("Introduction", "/docs2/introduction"),
				NextBlock("Core Concepts", "/docs2/core-concepts/pages"),
			),
		),
	)
}
