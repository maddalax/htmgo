package pages

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"mhtml/pages/base"
	"mhtml/partials"
	"mhtml/ui"
)

func ListPage(ctx *fiber.Ctx) *h.Page {
	return h.NewPage(base.RootPage(
		list(),
	))
}

func list() *h.Node {
	return h.Fragment(
		partials.SheetClosed(),
		h.Div(
			h.Class("inline-flex flex-col gap-4 p-4"),
			h.Div(
				h.Class("max-w-md flex flex-col gap-4 "),
				openButton(),
			),
			h.Div(
				h.View(partials.SheetOpenCount, h.ReloadParams{
					Triggers: h.CreateTriggers("load", "sheetOpened from:body"),
				}),
				h.Text("you opened sheet 0 times")),
		))
}

func openButton() *h.Node {
	return h.VStack(
		ui.PrimaryButton(ui.ButtonProps{
			Text:   "Open Sheet",
			Target: "#sheet-partial",
			Get:    h.GetPartialPathWithQs(partials.Sheet, "open=true"),
		}),
	)
}
