package main

import (
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/htmgo/service"
	_ "github.com/mattn/go-sqlite3"
	"htmgo-site/internal/markdown"
	"htmgo-site/pages"
	"htmgo-site/partials/load"
)

func main() {
	locator := service.NewLocator()

	service.Set(locator, service.Singleton, markdown.NewRenderer)

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(e *echo.Echo) {
			e.Static("/public", "./assets/dist")
			load.RegisterPartials(e)
			pages.RegisterPages(e)
			pages.RegisterMarkdown(e, "md", func(ctx echo.Context, path string) error {
				return pages.MarkdownHandler(ctx.(*h.RequestContext), path)
			})
		},
	})
}
