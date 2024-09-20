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
		h.Class("bg-red-200 flex flex-col items-center h-full w-full"),
		h.Head(
			h.Link("/public/main.css", "stylesheet"),
			h.Script("/public/htmgo.js"),
		),
		h.Body(
			h.Class("flex flex-col gap-4"),
			h.Div(h.Class("flex gap-2 mt-6"),
				Button(),
				Button(),
				Button(),
				Button(),
			),
		),
	))
}

func Button() h.Ren {
	return h.Button(h.Class("btn bg-green-500 p-4 rounded text-white"),
		h.Text("my button"),
		h.AfterRequest(
			h.SetDisabled(true),
			h.RemoveClass("bg-red-600"),
			h.AddClass("bg-gray-500"),
		),
		h.GetPartial(partials.SamplePartial),
	)
}
