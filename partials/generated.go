// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package partials

import "mhtml/h"
import "github.com/gofiber/fiber/v2"

func GetPartialFromContext(ctx *fiber.Ctx) *h.Partial {
	path := ctx.Path()
	if path == "SheetOpenCount" || path == "/mhtml/partials.SheetOpenCount" {
		return SheetOpenCount(ctx)
	}
	if path == "Sheet" || path == "/mhtml/partials.Sheet" {
		return Sheet(ctx)
	}
	return nil
}
