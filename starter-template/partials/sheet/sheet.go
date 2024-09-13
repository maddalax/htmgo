package sheet

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/mhtml/framework/h"
)

type Props struct {
	ClassName   string
	Root        h.Renderable
	OnClosePath string
}

var Id = "#active-modal"

func Opened(props Props) h.Renderable {
	return h.Fragment(h.Div(
		h.Class(`fixed top-0 right-0 h-full shadow-lg z-50`,
			h.Ternary(props.ClassName != "", props.ClassName, "w-96 bg-gray-100")),
		closeButton(props),
		h.Div(
			props.Root,
		)))
}

func Closed() h.Renderable {
	return h.Div(h.Id(Id))
}

func Close(ctx *fiber.Ctx) *h.Partial {
	return h.NewPartialWithHeaders(
		h.Ternary(ctx.Query("path") != "", h.ReplaceUrlHeader(ctx.Query("path")), h.NewHeaders()),
		h.Swap(ctx, Closed()),
	)
}

func closeButton(props Props) h.Renderable {
	return h.Div(
		h.Class("absolute top-0 right-0 p-3"),
		h.Button(
			h.Class("text-gray-500"),
			h.GetPartialWithQs(Close, fmt.Sprintf("path=%s", props.OnClosePath)),
			h.Text("X"),
		),
	)
}
