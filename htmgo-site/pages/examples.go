package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/pages/base"
)

type Example struct {
	Title       string
	Github      string
	Demo        string
	Image       string
	Description string
}

var examples = []Example{
	{
		Title:  "Chat App Example",
		Github: "https://github.com/maddalax/htmgo/tree/master/examples/chat",
		Demo:   "https://chat-example.htmgo.dev",
		Image:  "public/chat-example.jpg",
	},
	{
		Title:  "Todo List MVC",
		Github: "https://github.com/maddalax/htmgo/tree/master/examples/todo-list",
		Demo:   "https://todo-example.htmgo.dev",
		Image:  "public/todo-example.jpg",
	},
	{
		Title:  "htmgo.dev",
		Github: "https://github.com/maddalax/htmgo/tree/master/htmgo-site",
		Demo:   "https://htmgo.dev",
		Image:  "public/htmgo-site.jpg",
	},
	{
		Title:       "Form With Loading State",
		Github:      "https://github.com/maddalax/htmgo/blob/master/htmgo-site/pages/form.go",
		Demo:        "/form",
		Description: "A simple form submission example with a loading state",
	},
}

func ExamplesPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		base.RootPage(ctx, h.Div(
			h.Class("flex items-center justify-center"),
			h.Div(
				h.Class("w-full px-4 flex flex-col prose max-w-[95vw] md:max-w-3xl mt-6"),
				h.Div(
					h.Class("flex flex-col mb-6 md:mb-0 md:flex-row justify-between items-center"),
					h.Div(
						h.H1(
							h.Class("text-center md:text-left"),
							h.Text("htmgo examples"),
						),
						h.H3(
							h.Class("-mt-4"),
							h.TextF("example projects built with htmgo"),
						),
					),
				),
				h.Div(
					h.Class("border-b border-b-slate-200 h-1"),
					h.Div(
						h.Class("mt-4"),
						ExampleCards(),
					),
				),
			)),
		),
	)
}

func ExampleCards() *h.Element {
	return h.Div(
		h.Class("prose-h2:my-1 prose-img:my-1 grid grid-cols-1 md:grid-cols-3 gap-6 text-center pb-8"), // Using grid for 3-column layout
		h.List(examples, func(example Example, index int) *h.Element {
			return h.Div(
				h.Class("border border-gray-200 shadow-sm rounded-md px-4 pb-4 bg-neutral-100"), // Removed specific width, handled by grid
				h.Div(
					h.Class("flex flex-col gap-1 mt-4"),
					h.H2(
						h.Class("text-lg text-center mb-1"), // Reduced margin at the bottom of the title
						h.Text(example.Title),
					),
					h.If(example.Image != "", h.Div(
						h.A(
							h.Href(example.Demo),
							h.Class("not-prose"),
							h.Img(
								h.Src(example.Image),
								h.Class("md:w-full rounded-md mx-auto"),
							),
						), // Ensures image is centered within the card
					)),
					h.If(example.Description != "", h.Pf(example.Description)),
					h.Div(
						h.Div(
							h.Class("flex gap-2 justify-center mt-2"), // Slight margin-top for spacing from the image
							h.A(
								h.Href(example.Github),
								h.Class("not-prose p-2 bg-slate-900 text-white rounded-md"), // Reduced padding for the buttons
								h.Text("Github"),
							),
							h.A(
								h.Href(example.Demo),
								h.Class("not-prose p-2 bg-slate-900 text-white rounded-md"), // Reduced padding for the buttons
								h.Text("Demo"),
							),
						),
					),
				),
			)
		}),
	)
}
