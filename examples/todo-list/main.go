package main

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	_ "github.com/mattn/go-sqlite3"
	"io/fs"
	"todolist/__htmgo"
	"todolist/ent"
	"todolist/infrastructure/db"
)

//go:embed assets/dist/*
var StaticAssets embed.FS

func main() {
	locator := service.NewLocator()

	service.Set[ent.Client](locator, service.Singleton, func() *ent.Client {
		return db.Provide()
	})

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
