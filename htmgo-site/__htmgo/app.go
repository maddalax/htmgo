package __htmgo

import (
	"github.com/labstack/echo/v4"
	"htmgo-site/pages"
	"htmgo-site/partials"
)
import "github.com/maddalax/htmgo/framework/h"

func RegisterPages(f *echo.Echo) {
	f.GET("/docs", func(ctx echo.Context) error {
		cc := ctx.(*h.RequestContext)
		return h.HtmlView(ctx, pages.DocsPage(cc))
	})
	f.GET("/", func(ctx echo.Context) error {
		cc := ctx.(*h.RequestContext)
		return h.HtmlView(ctx, pages.IndexPage(cc))
	})
}

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
