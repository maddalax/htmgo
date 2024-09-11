package partials

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"mhtml/news"
	"mhtml/ui"
)

func SheetOpenCount(ctx *fiber.Ctx) *h.Partial {
	rnd := h.SessionGet[int64](ctx, "sheet-open-count")
	if rnd == nil {
		rnd = new(int64)
	}
	return h.NewPartial(h.Div(
		h.Text(fmt.Sprintf("you opened sheet %d times", *rnd)),
	))
}

func SheetClosed() *h.Node {
	return h.Div(h.Id("sheet-partial"))
}

func Sheet(ctx *fiber.Ctx) *h.Partial {
	open := ctx.Query("open")
	if open == "true" {
		h.SessionIncr(ctx, "sheet-open-count")
	}
	return h.NewPartialWithHeaders(
		&map[string]string{
			"hx-trigger": "sheetOpened",
		},
		h.IfElse(open == "true", SheetOpen(), SheetClosed()),
	)
}

func SheetOpen() *h.Node {
	return h.Div(
		h.Class(`fixed top-0 right-0 h-full w-96 bg-gray-100 shadow-lg z-50`),
		h.Div(
			h.Class("p-4 overflow-y-auto h-full w-full flex flex-col gap-4"),
			h.P("My Sheet",
				h.Class("text-lg font-bold"),
			),
			h.P("This is a sheet",
				h.Class("text-sm mt-2"),
			),
			ui.Button(ui.ButtonProps{
				Text:   "Close Sheet",
				Target: "#sheet-partial",
				Get:    h.GetPartialPathWithQs(Sheet, "open=false"),
			}),
			news.StoryList(),
		))
}
