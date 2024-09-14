package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/htmgo/framework/h"
	"log"
	"starter-template/pages"
	"starter-template/partials/load"
	"time"
)

func main() {
	f := fiber.New()

	f.Static("/public", "./assets/dist")

	f.Use(func(ctx *fiber.Ctx) error {
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
