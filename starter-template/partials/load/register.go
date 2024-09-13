package load

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/mhtml/framework/h"
)

func RegisterPartials(f *fiber.App) {
	f.All("/mhtml/partials*", func(ctx *fiber.Ctx) error {
		partial := GetPartialFromContext(ctx)
		if partial == nil {
			return ctx.SendStatus(404)
		}
		return h.PartialView(ctx, partial)
	})
}
