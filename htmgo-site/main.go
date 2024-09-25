package main

import (
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"htmgo-site/__htmgo"
	"htmgo-site/internal/markdown"
	"io/fs"
)

func main() {
	locator := service.NewLocator()
	staticAssets := GetStaticAssets()
	markdownAssets := GetMarkdownAssets()

	service.Set(locator, service.Singleton, markdown.NewRenderer)

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(e *echo.Echo) {
			sub, err := fs.Sub(staticAssets, "assets/dist")
			if err != nil {
				panic(err)
			}
			e.StaticFS("/public", sub)

			e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					c.Set("embeddedMarkdown", markdownAssets)
					return next(c)
				}
			})

			__htmgo.Register(e)
		},
	})
}
