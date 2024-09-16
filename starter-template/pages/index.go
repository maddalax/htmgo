package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/partials"
)

func IndexPage(c *fiber.Ctx) *h.Page {
	return h.NewPage(h.Html(
		h.Class("bg-blue-400 flex flex-col items-center h-full w-full"),
		h.Head(
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
		),
		h.Body(
			h.Class("flex flex-col gap-4"),
			h.Div(
				h.Class("flex flex-col items-center justify-center gap-6 p-12 text-center"),
				h.H1(
					h.Class("text-4xl sm:text-5xl font-bold max-w-3xl"),
					h.Text("Welcome to my fast!!"),
				),
				h.P(
					h.Class("text-lg sm:text-xl max-w-1xl"),
				),
				h.Div(
					Button(),
				),
			),
		),
	))
}

func Button() h.Renderable {
	return h.Button(h.Class("btn bg-slate-500 p-4 rounded text-white"),
		h.Text("Ctest"),
		h.AfterRequest(
			h.SetDisabled(true),
			h.RemoveClass("bg-red-600"),
			h.AddClass("bg-gray-500"),
		),
		h.GetPartial(partials.SamplePartial),
	)
}
