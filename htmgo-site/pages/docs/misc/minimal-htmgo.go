package misc

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
)

func MinimalHtmgo(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Minimal Htmgo"),
			Text(`
				Looking to use htmgo as an html builder but don't want to install the entire framework? View our minimal htmgo example below:
			`),
			Link("Minimal htmgo template", "https://github.com/maddalax/htmgo/tree/master/examples/minimal-htmgo"),
			Text("This is the most minimal example of htmgo you can get, it does not require the htmgo cli, and can be used with any go project."),
			Text("Since this is a minimal example, it does not include many of the main features that the htmgo framework provides, which are: live reload, automatic tailwind css generation, and automatic partial/page registration. Those features are only available in the full htmgo framework that utilizes the htmgo cli for preprocessing."),
			Text("This is useful for those who want to use htmgo as an html builder with js evaluation support, but do not want to use the full htmgo framework, such as when embedding into a different application."),
			NextStep(
				"mt-4",
				PrevBlock("Introduction", DocPath("/introduction")),
				NextBlock("Core Concepts", DocPath("/core-concepts/pages")),
			),
		),
	)
}
