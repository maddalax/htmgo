package pages

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
)

func IndexPage(c *fiber.Ctx) *h.Page {
	return h.NewPage(h.Html(
		h.HxExtension("path-deps, response-targets, mutation-error"),
		h.Head(
			h.Script("https://cdn.tailwindcss.com"),
			h.Script("/js/dist/mhtml.js"),
		),
		h.Body(),
	))
}
