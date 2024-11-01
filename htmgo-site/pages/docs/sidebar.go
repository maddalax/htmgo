package docs

import (
	"github.com/maddalax/htmgo/framework/h"
)

type Section struct {
	Title string
	Pages []*Page
}

type Page struct {
	Title string
	Path  string
}

func DocPath(path string) string {
	return "/docs" + path
}

var sections = []Section{
	{
		Title: "Getting Started",
		Pages: []*Page{
			{Title: "Introduction", Path: DocPath("/introduction")},
			{Title: "Quick Start", Path: DocPath("/installation")},
			{Title: "Related Projects", Path: DocPath("/related-projects")},
		},
	},
	{
		Title: "Core Concepts",
		Pages: []*Page{
			{Title: "Pages", Path: DocPath("/core-concepts/pages")},
			{Title: "Partials", Path: DocPath("/core-concepts/partials")},
			{Title: "Components", Path: DocPath("/core-concepts/components")},
			{Title: "Tags and Attributes", Path: DocPath("/core-concepts/tags-and-attributes")},
			{Title: "Raw HTML", Path: DocPath("/core-concepts/raw-html")},
		},
	},
	{
		Title: "Control",
		Pages: []*Page{
			{Title: "Conditionals", Path: DocPath("/control/if-else")},
			{Title: "Rendering Lists", Path: DocPath("/control/loops")},
		},
	},
	{
		Title: "Interactivity",
		Pages: []*Page{
			{Title: "Swapping", Path: DocPath("/interactivity/swapping")},
			{Title: "Events", Path: DocPath("/interactivity/events")},
			{Title: "Evaluating Javascript", Path: DocPath("/interactivity/events")},
			{Title: "Eval Commands", Path: DocPath("/interactivity/eval-commands")},
		},
	},
	{
		Title: "Performance",
		Pages: []*Page{
			{Title: "Caching Globally", Path: DocPath("/performance/caching-globally")},
			{Title: "Caching Per Key", Path: DocPath("/performance/caching-per-key")},
		},
	},
	{
		Title: "Pushing Data",
		Pages: []*Page{
			{Title: "Server Sent Events", Path: DocPath("/pushing-data/sse")},
		},
	},
	{
		Title: "HTMX Extensions",
		Pages: []*Page{
			{Title: "Overview", Path: DocPath("/htmx-extensions/overview")},
			{Title: "Trigger Children", Path: DocPath("/htmx-extensions/trigger-children")},
			{Title: "Mutation Error", Path: DocPath("/htmx-extensions/mutation-error")},
		},
	},
	{
		Title: "Miscellaneous",
		Pages: []*Page{
			{Title: "Tailwind Intellisense", Path: DocPath("/misc/tailwind-intellisense")},
			{Title: "Formatter", Path: DocPath("/misc/formatter")},
		},
	},
	{
		Title: "Configuration",
		Pages: []*Page{
			{Title: "Htmgo Config", Path: DocPath("/config/htmgo-config")},
		},
	},
}

func DocSidebar() *h.Element {
	return h.Div(
		h.Class("px-3 py-2 pr-6 md:min-h-screen pb-4 mb:pb-0 bg-neutral-50 border-r border-r-slate-300 overflow-y-auto"),
		h.Div(
			h.Div(
				h.Class("mb-3"),
				h.A(
					h.Href("#quick-start-introduction"),
					h.Text("Documentation"),
					h.Class("md:mt-4 text-xl text-slate-900 font-bold"),
				),
			),
			h.Div(
				h.Class("flex flex-col gap-4"),
				h.List(sections, func(entry Section, index int) *h.Element {
					return h.Div(
						h.P(
							h.Text(entry.Title),
							h.Class("text-slate-800 font-bold"),
						),
						h.Div(
							h.Class("pl-4 flex flex-col"),
							h.List(entry.Pages, func(page *Page, index int) *h.Element {
								return h.A(
									h.Href(page.Path),
									h.Text(page.Title),
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
