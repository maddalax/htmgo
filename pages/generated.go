// Package pages THIS FILE IS GENERATED. DO NOT EDIT.
package pages

import "github.com/gofiber/fiber/v2"
import "mhtml/h"

func RegisterPages(f *fiber.App) {
	f.Get("/", func(ctx *fiber.Ctx) error {
		return h.HtmlView(ctx, IndexPage(ctx))
	})
	f.Get("/news/:id", func(ctx *fiber.Ctx) error {
		return h.HtmlView(ctx, Test(ctx))
	})
	f.Get("/news", func(ctx *fiber.Ctx) error {
		return h.HtmlView(ctx, ListPage(ctx))
	})
	f.Get("/patients", func(ctx *fiber.Ctx) error {
		return h.HtmlView(ctx, PatientsIndex(ctx))
	})
}
