package sheet

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
)

type Props struct {
	ClassName string
	Root      *h.Node
}

var Id = "#active-modal"

func Opened(props Props) *h.Node {
	return h.Fragment(h.Div(
		h.Class(`fixed top-0 right-0 h-full shadow-lg z-50`,
			h.Ternary(props.ClassName != "", props.ClassName, "w-96 bg-gray-100")),
		closeButton(),
		h.Div(
			props.Root,
		)))
}

func Closed() *h.Node {
	return h.Div(h.Id(Id))
}

func Close(ctx *fiber.Ctx) *h.Partial {
	return h.NewPartial(
		h.Swap(ctx, Closed()),
	)
}

func closeButton() *h.Node {
	return h.Div(
		h.Class("absolute top-0 right-0 p-3"),
		h.Button(
			h.Class("text-gray-500"),
			h.GetPartial(Close),
			h.Text("X"),
		),
	)
}
