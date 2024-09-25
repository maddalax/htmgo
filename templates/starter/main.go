package main

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"io/fs"
	"starter-template/__htmgo"
)

//go:embed assets/dist/*
var StaticAssets embed.FS

func main() {
	locator := service.NewLocator()

	sub, err := fs.Sub(StaticAssets, "assets/dist")

	if err != nil {
		panic(err)
	}

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(e *echo.Echo) {
			e.StaticFS("/public", sub)
			__htmgo.Register(e)
		},
	})
}
