package partials

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework-ui/ui"
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/news"
)

func NewsSheet(ctx echo.Context) *h.Partial {
	open := h.GetQueryParam(ctx, "open") == "true"
	return h.NewPartialWithHeaders(
		&map[string]string{
			"hx-trigger":  "sheetOpened",
			"hx-push-url": fmt.Sprintf("/news%s", h.Ternary(open, "?open=true", "")),
		},
		SheetWrapper(
			h.IfElseLazy(open, SheetOpen, SheetClosed),
			h.OobSwap(ctx, OpenSheetButton(open)),
			h.OobSwap(ctx, NewsSheetOpenCount(ctx).Root),
		),
	)
}

func NewsSheetOpenCount(ctx echo.Context) *h.Partial {

	open := h.GetQueryParam(ctx, "open") == "true"

	return h.NewPartial(h.Div(
		h.Id("sheet-open-count"),
		h.IfElse(open,
			h.Text(fmt.Sprintf("you opened sheet %d times", 1)),
			h.Text("sheet is not open")),
	),
	)
}

func SheetWrapper(children ...h.Ren) h.Ren {
	return h.Div(h.Id("sheet-partial"), h.Fragment(children...))
}

func SheetClosed() h.Ren {
	return h.Div()
}

func SheetOpen() h.Ren {
	return h.Fragment(h.Div(
		h.Class(`fixed top-0 right-0 h-full w-96 bg-gray-100 shadow-lg z-50`),
		h.Div(
			h.Class("p-4 overflow-y-auto h-full w-full flex flex-col gap-4"),
			h.P(h.Text("News Sheet"),
				h.Class("text-lg font-bold"),
			),
			h.P(h.Text("Here are the latest news stories."),
				h.Class("text-sm mt-2"),
			),
			ui.Button(ui.ButtonProps{
				Text:   "Close NewsSheet",
				Target: "#sheet-partial",
				Get:    h.GetPartialPathWithQs(NewsSheet, "open=false"),
			}),
			news.StoryList(),
		)))
}
