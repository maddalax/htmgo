package sheet

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
)

func Opened(children ...*h.Node) *h.Node {
	return h.Fragment(h.Div(
		h.Class(`fixed top-0 right-0 h-full w-96 bg-gray-100 shadow-lg z-50`),
		CloseButton(),
		h.Div(
			children...,
		)))
}

func Closed() *h.Node {
	return h.Div(h.Id("active-modal"))
}

func Close(ctx *fiber.Ctx) *h.Partial {
	return h.NewPartial(
		h.Swap(ctx, Closed()),
	)
}

func CloseButton() *h.Node {
	return h.Div(
		h.Class("absolute top-0 right-0 p-3"),
		h.Button(
			h.Class("text-gray-500"),
			h.GetPartial(Close),
			h.Text("X"),
		),
	)
}
