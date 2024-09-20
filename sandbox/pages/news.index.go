package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/pages/base"
	"starter-template/partials"
)

func ListPage(ctx echo.Context) *h.Page {
	return h.NewPage(base.RootPage(
		list(ctx),
	))
}

func list(ctx echo.Context) h.Ren {
	return h.Fragment(
		h.ViewOnLoad(partials.NewsSheet),
		h.Div(
			h.Class("inline-flex flex-col gap-4 p-4"),
			h.Div(
				h.Class("max-w-md flex flex-col gap-4 "),
				partials.OpenSheetButton(h.GetQueryParam(ctx, "open") == "true"),
			),
			h.Div(
				h.ViewOnLoad(partials.NewsSheetOpenCount),
				h.Text("you opened sheet 0 times")),
		))
}
