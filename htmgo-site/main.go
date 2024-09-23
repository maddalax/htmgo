package main

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	_ "github.com/mattn/go-sqlite3"
	"htmgo-site/__htmgo"
	"htmgo-site/internal/markdown"
	"htmgo-site/pages"
	"io/fs"
)

//go:embed assets/dist/*
var StaticAssets embed.FS

//go:embed md/*
var MarkdownAssets embed.FS

func main() {
	locator := service.NewLocator()

	service.Set(locator, service.Singleton, markdown.NewRenderer)

	sub, err := fs.Sub(StaticAssets, "assets/dist")

	if err != nil {
		panic(err)
	}

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(e *echo.Echo) {
			e.StaticFS("/public", sub)

			e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					c.Set("embeddedMarkdown", &MarkdownAssets)
					return next(c)
				}
			})

			__htmgo.RegisterPartials(e)
			__htmgo.RegisterPages(e)

			pages.RegisterMarkdown(e, "md", MarkdownAssets, func(ctx echo.Context, path string) error {
				return pages.MarkdownHandler(ctx.(*h.RequestContext), path)
			})
		},
	})
}
