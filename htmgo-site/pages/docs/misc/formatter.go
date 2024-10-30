package misc

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

func Formatter(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Formatter"),
			Text(`
				htmgo has a built-in formatter that can be used to format htmgo element blocks.
				It is available through the 'htmgo' cli tool that is installed with htmgo.
			`),
			HelpText(`Note: if you have previously installed htmgo, you will need to run GOPROXY=direct go install github.com/maddalax/htmgo/cli/htmgo@latest to update the cli tool.`),
			Text("Usage:"),
			ui.SingleLineBashCodeSnippet(`htmgo format .`),
			HelpText(`This will format all htmgo element blocks in your project recursively.`),
			ui.SingleLineBashCodeSnippet(`htmgo format ./my-file.go`),
			HelpText(`This will format the file specified.`),
			Text("Before:"),
			ui.GoCodeSnippet(formatBefore),
			Text("After:"),
			ui.GoCodeSnippet(formatAfter),
			h.Div(
				h.Class("hidden md:block w-[800px] h-[800px] rounded"),
				Video(),
			),
			NextStep(
				"mt-4",
				PrevBlock("Tailwind Intellisense", DocPath("/misc/tailwind-intellisense")),
				NextBlock("Configuration", DocPath("/config/htmgo-config")),
			),
		),
	)
}

const formatBefore = `h.Div(
	h.Class("flex gap-2"), h.Text("hello"), h.Text("world"),
)`

const formatAfter = `h.Div(
	h.Class("flex gap-2"),
	h.Text("hello"),
	h.Text("world"),
)
`

func Video() *h.Element {
	return h.Video(
		h.Tag(
			"source",
			h.Src("/public/formatter.mp4"),
			h.Type("video/mp4"),
		),
		h.Controls(),
		h.Class("h-full w-full rounded"),
	)
}
