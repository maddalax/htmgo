package pages

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"mhtml/pages/base"
)

func IndexPage(c *fiber.Ctx) *h.Page {
	return h.NewPage(base.RootPage(h.P("this is cool")))
}
