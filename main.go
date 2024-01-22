package main

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/database"
	"mhtml/h"
	"mhtml/news"
	"strconv"
	"time"
)

type User struct {
	Name  string
	Email string
}

var Version = time.Now().Nanosecond()

func LiveReloadHandler(c *fiber.Ctx) error {
	v := strconv.FormatInt(int64(Version), 10)
	current := c.Cookies("version", v)

	if current != v {
		c.Set("HX-Refresh", "true")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "version",
		Value: v,
	})
	return c.SendString("")
}

func Page(children ...*h.Node) *h.Node {
	return h.Page(
		h.Head(
			h.Script("https://cdn.tailwindcss.com"),
			h.Script("https://unpkg.com/htmx.org@1.9.2"),
		),
		h.Body(
			h.VStack(
				h.Class("flex flex-col gap-2 bg-gray-100 h-full w-full items-center p-12"),
				h.Fragment(children...),
			),
		),
	)
}

func IndexPage(c *fiber.Ctx) error {
	page := Page(
		h.Div(
			h.P("Hacker News - Top Stories"),
		),
		news.StoryList(),
	)
	return HtmlView(c, page)
}

func HtmlView(c *fiber.Ctx, child *h.Node) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(
		h.Render(
			child,
		),
	)
}

func main() {

	database.HSet("users", "sydne", User{
		Name:  "Sydne Anschutz",
		Email: "sanschutz0808@gmail.com",
	})

	app := fiber.New()

	app.Get("/", IndexPage)
	app.Get("/news/:id", func(c *fiber.Ctx) error {
		return HtmlView(c, Page(
			news.StoryFull(c.Params("id")),
		))
	})
	app.Get("/livereload", LiveReloadHandler)

	app.Listen(":3000")
}
