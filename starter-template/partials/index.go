package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/htmgo/framework/h"
)

func SamplePartial(ctx *fiber.Ctx) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text(" asdas"))))
}

func NewPartial(ctx *fiber.Ctx) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text("This sadsl."))))
}

func NewPartial2(ctx *fiber.Ctx) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text("This sasdsadasdwl."))))
}
