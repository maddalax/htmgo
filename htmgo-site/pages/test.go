package pages

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
)

func TestFormatPage(ctx *h.RequestContext) *h.Page {
	return RootPage(
		ctx,
		h.Div(
			h.P(
				h.Class("hello"),
				h.Details(
					h.Summary(
						h.Text("Summary"),
					),
					h.Text("Details"),
				),
				h.Id("hi"),
			),
		),
	)
}

func notPage() int {
	test := 1
	fmt.Printf("test: %d\n", test)
	return test
}

func TestOtherPage(ctx *h.RequestContext) *h.Page {
	return RootPage(
		ctx,
		h.Div(
			h.Id("test"),
			h.Details(
				h.Summary(
					h.Text("Summary"),
				),
				h.Text("Details"),
			),
			h.Class("flex flex-col gap-2 bg-white h-full"),
			h.Id("test"),
			h.Details(
				h.Summary(
					h.Text("Summary"),
				),
				h.Text("Details"),
			),
		),
	)
}
