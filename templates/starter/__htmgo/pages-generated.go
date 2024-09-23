// Package __htmgo THIS FILE IS GENERATED. DO NOT EDIT.
package __htmgo

import "github.com/labstack/echo/v4"
import "github.com/maddalax/htmgo/framework/h"
import "starter-template/pages"

func RegisterPages(f *echo.Echo) {
	f.GET("/", func(ctx echo.Context) error {
		cc := ctx.(*h.RequestContext)
		return h.HtmlView(ctx, pages.IndexPage(cc))
	})
}
