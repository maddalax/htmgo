package pages

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"mhtml/pages/base"
)

func IndexPage(c *fiber.Ctx) *h.Page {
	return h.NewPage(base.RootPage(
		h.Fragment(
			h.Div(
				h.Class("inline-flex flex-col gap-4 p-4"),
				h.Div(
					h.Class("max-w-md flex flex-col gap-4"),
					h.P("Routes"),
				),
			)),
	))
}
