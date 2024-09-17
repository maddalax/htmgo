// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package load

import "github.com/maddalax/htmgo/framework/h"
import "github.com/labstack/echo/v4"
import "starter-template/partials"

func GetPartialFromContext(ctx echo.Context) *h.Partial {
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

func RegisterPartials(f *echo.Echo) {
	f.Any("starter-template/partials*", func(ctx echo.Context) error {
		partial := GetPartialFromContext(ctx)
		if partial == nil {
			return ctx.NoContent(404)
		}
		return h.PartialView(ctx, partial)
	})
}
