// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package load

import "github.com/maddalax/htmgo/framework/h"
import "github.com/gofiber/fiber/v2"
import "starter-template/partials"

func GetPartialFromContext(ctx *fiber.Ctx) *h.Partial {
	path := ctx.Path()
	if path == "SamplePartial" || path == "/starter-template/partials.SamplePartial" {
		return partials.SamplePartial(ctx)
	}
	if path == "NewPartial" || path == "/starter-template/partials.NewPartial" {
		return partials.NewPartial(ctx)
	}
	if path == "NewPartial2" || path == "/starter-template/partials.NewPartial2" {
		return partials.NewPartial2(ctx)
	}
	return nil
}

func RegisterPartials(f *fiber.App) {
	f.All("starter-template/partials*", func(ctx *fiber.Ctx) error {
		partial := GetPartialFromContext(ctx)
		if partial == nil {
			return ctx.SendStatus(404)
		}
		return h.PartialView(ctx, partial)
	})
}
