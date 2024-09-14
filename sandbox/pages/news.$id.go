package pages

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/htmgo/framework/h"
)

func Test(ctx *fiber.Ctx) *h.Page {
	text := fmt.Sprintf("News ID: %s", ctx.Params("id"))
	return h.NewPage(
		h.Div(h.Text(text)),
	)
}
