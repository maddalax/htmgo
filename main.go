package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"mhtml/h"
	"mhtml/pages"
	"mhtml/partials"
	"time"
)

func main() {
	f := fiber.New()
	f.Static("/js", "./js")

	f.Use(func(ctx *fiber.Ctx) error {
		if ctx.Cookies("mhtml-session") != "" {
			return ctx.Next()
		}
		id := ctx.IP() + uuid.NewString()
		ctx.Cookie(&fiber.Cookie{
			Name:        "mhtml-session",
			Value:       id,
			SessionOnly: true,
		})
		return ctx.Next()
	})

	f.Use(func(ctx *fiber.Ctx) error {
		if ctx.Path() == "/livereload" {
			return ctx.Next()
		}
		now := time.Now()
		err := ctx.Next()
		duration := time.Since(now)
		ctx.Set("X-Response-Time", duration.String())
		// Log or print the request method, URL, and duration
		log.Printf("Request: %s %s took %v", ctx.Method(), ctx.OriginalURL(), duration)
		return err
	})

	f.Get("/mhtml/partials*", func(ctx *fiber.Ctx) error {
		partial := partials.GetPartialFromContext(ctx)
		if partial == nil {
			return ctx.SendStatus(404)
		}
		return h.PartialView(ctx, partial)
	})

	pages.RegisterPages(f)

	h.Start(f, h.App{
		LiveReload: true,
	})
}
