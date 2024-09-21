// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package load

import "github.com/maddalax/htmgo/framework/h"
import "github.com/labstack/echo/v4"
import "htmgo-site/partials"

func GetPartialFromContext(ctx echo.Context) *h.Partial {
	path := ctx.Request().URL.Path
	if path == "ToggleNavbar" || path == "/htmgo-site/partials.ToggleNavbar" {
		cc := ctx.(*h.RequestContext)
		return partials.ToggleNavbar(cc)
	}
	return nil
}

func RegisterPartials(f *echo.Echo) {
	f.Any("htmgo-site/partials*", func(ctx echo.Context) error {
		partial := GetPartialFromContext(ctx)
		if partial == nil {
			return ctx.NoContent(404)
		}
		return h.PartialView(ctx, partial)
	})
}
