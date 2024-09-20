package partials

import "github.com/maddalax/htmgo/framework/h"

func SideBar() *h.Element {
	return h.Div(
		h.Class("w-40 top-[57px] absolute min-h-screen bg-neutral-50 border border-r-slate-300 p-3"),
		h.Div(
			h.Class("max-w-prose mx-auto"),
			h.Div(
				h.Class("flex flex-col gap-4"),
				h.A(
					h.Href("/docs"),
					h.Text("Docs"),
				),
				h.A(
					h.Href("/examples"),
					h.Text("Examples"),
				),
			),
		),
	)
}
