package pages

import (
	"github.com/maddalax/htmgo/framework/js"
	"todolist/pages/base"
	"todolist/partials/task"

	"github.com/maddalax/htmgo/framework/h"
)

func TaskListPage(ctx *h.RequestContext) *h.Page {

	title := h.Div(
		h.H1(h.Class("text-7xl font-extralight text-rose-500 tracking-wide"), h.Text("todos")),
	)

	return h.NewPage(base.RootPage(
		h.Div(
			h.Class("bg-neutral-100 min-h-screen"),
			h.Div(
				h.Class("flex flex-col gap-6 p-4 items-center max-w-xl mx-auto pb-12"),
				title,
				h.Button(
					h.Text("Add Task"),
					h.OnClick(
						js.InjectScript("https://htmgo.dev"),
					),
				),
				task.Card(ctx),
				h.Children(
					h.Div(h.Text("Double-click to edit a todo")),
				),
			),
		),
	))
}
