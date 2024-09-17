package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"log"
	"starter-template/pages"
	"starter-template/partials/load"
	"time"
)

func main() {
	f := echo.New()

	f.Static("/public", "./assets/dist")

	f.Use(func(ctx echo.Context) error {
		if ctx.Cookies("htmgo-session") != "" {
			return ctx.Next()
		}
		id := ctx.IP() + uuid.NewString()
		ctx.Cookie(&echo.Cookie{
			Name:        "htmgo-session",
			Value:       id,
			SessionOnly: true,
		})
		return ctx.Next()
	})

	f.Use(func(ctx echo.Context) error {
		if ctx.Path() == "/livereload" {
			return ctx.Next()
		}
		now := time.Now()
		err := ctx.Next()
		duration := time.Since(now)
		ctx.Set("X-Response-Time", duration.String())
		// Log or print the request method, URL, and duration
		log.Printf("Request: %s %s took %dms", ctx.Method(), ctx.OriginalURL(), duration.Milliseconds())
		return err
	})

	load.RegisterPartials(f)
	pages.RegisterPages(f)

	h.Start(f, h.App{
		LiveReload: true,
	})
}
