package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/pages/base"
	"starter-template/partials"
)

func IndexPage(c echo.Context) *h.Page {
	return h.NewPage(h.Html(
		h.HxExtension(base.Extensions()),
		h.Class("bg-slate-100 flex flex-col items-center h-full w-full"),
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
					h.Text("test"),
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
	return h.Button(h.Class("btn bg-red-500 p-4 rounded text-white"),
		h.Text("this is my nice this works"),
		h.AfterRequest(
			h.SetDisabled(true),
			h.RemoveClass("bg-red-600"),
			h.AddClass("bg-gray-500"),
		),
		h.GetPartial(partials.SamplePartial),
	)
}
