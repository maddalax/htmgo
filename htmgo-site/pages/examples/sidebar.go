package examples

import (
	"github.com/maddalax/htmgo/framework/datastructure/orderedmap"
	"github.com/maddalax/htmgo/framework/h"
)

func SnippetSidebar() *h.Element {

	grouped := h.GroupByOrdered(examples, func(example Snippet) string {
		return example.category
	})

	return h.Div(
		h.Class("px-3 py-2 pr-6 md:min-h-screen pb-4 mb:pb-0 bg-neutral-50 border-r border-r-slate-300 overflow-y-auto"),
		h.Div(
			h.Div(
				h.Class("mb-3"),
				h.A(
					h.Href("#"),
					h.Text("Examples"),
					h.Class("md:mt-4 text-xl text-slate-900 font-bold"),
				),
			),
			h.Div(
				h.Class("flex flex-col gap-2"),
				h.List(grouped.Entries(), func(entry orderedmap.Entry[string, []Snippet], index int) *h.Element {
					return h.Div(
						h.P(
							h.Text(entry.Key),
							h.Class("text-slate-800 font-bold"),
						),
						h.Div(
							h.Class("pl-4 flex flex-col"),
							h.List(entry.Value, func(entry Snippet, index int) *h.Element {
								return h.A(
									h.Href(entry.path),
									h.Text(entry.sidebarName),
									h.Class("text-slate-900 hover:text-rose-400"),
								)
							}),
						),
					)
				}),
			),
		),
	)
}
