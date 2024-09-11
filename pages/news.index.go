package pages

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"mhtml/pages/base"
	"mhtml/partials"
)

func ListPage(ctx *fiber.Ctx) *h.Page {
	return h.NewPage(base.RootPage(
		list(ctx),
	))
}

func list(ctx *fiber.Ctx) *h.Node {
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
