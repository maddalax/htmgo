// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package load

import "github.com/maddalax/htmgo/framework/h"
import "github.com/labstack/echo/v4"
import "starter-template/partials"

func GetPartialFromContext(ctx echo.Context) *h.Partial {
	path := ctx.Request().URL.Path
	if path == "SamplePartial" || path == "/starter-template/partials.SamplePartial" {
		cc := ctx.(*h.RequestContext)
		return partials.SamplePartial(cc)
	}
	if path == "NewPartial" || path == "/starter-template/partials.NewPartial" {
		cc := ctx.(*h.RequestContext)
		return partials.NewPartial(cc)
	}
	if path == "NewPartial2" || path == "/starter-template/partials.NewPartial2" {
		cc := ctx.(*h.RequestContext)
		return partials.NewPartial2(cc)
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
