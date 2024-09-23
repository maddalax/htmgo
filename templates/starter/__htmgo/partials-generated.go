// Package __htmgo THIS FILE IS GENERATED. DO NOT EDIT.
package __htmgo

import "github.com/maddalax/htmgo/framework/h"
import "github.com/labstack/echo/v4"
import "starter-template/partials"

func GetPartialFromContext(ctx echo.Context) *h.Partial {
	path := ctx.Request().URL.Path
	if path == "CounterPartial" || path == "/starter-template/partials.CounterPartial" {
		cc := ctx.(*h.RequestContext)
		return partials.CounterPartial(cc)
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
