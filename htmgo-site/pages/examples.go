package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/pages/base"
)

type Example struct {
	Title  string
	Github string
	Demo   string
	Image  string
}

var examples = []Example{
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
		h.Class("prose-h2:my-1 prose-img:my-1 flex gap-2"), // Left-aligns and allows multiple cards in a row
		h.List(examples, func(example Example, index int) *h.Element {
			return h.Div(
				h.Class("border border-gray-200 shadow-sm rounded-md px-4 pb-4 w-full md:w-1/2"), // Reduces padding
				h.Div(
					h.Class("flex flex-col gap-1 mt-4"),
					h.H2(
						h.Class("text-lg text-center mb-1"), // Reduced margin at the bottom of the title
						h.Text(example.Title),
					),
					h.Div(
						h.A(
							h.Href(example.Demo),
							h.Class("not-prose"),
							h.Img(
								h.Src(example.Image),
								h.Class("md:w-full rounded-md mx-auto"),
							),
						), // Ensures image is centered within the card
					),
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
