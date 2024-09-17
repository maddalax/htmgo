package h

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/htmgo/framework/util/process"
	"log/slog"
	"time"
)

type App struct {
	LiveReload bool
	Fiber      *fiber.App
}

var instance *App

func GetApp() *App {
	if instance == nil {
		panic("App instance not initialized")
	}
	return instance
}

func Start(app *fiber.App, opts App) {
	instance = &opts
	instance.start(app)
}

func (a App) start(app *fiber.App) {

	a.Fiber = app

	if a.LiveReload {
		AddLiveReloadHandler("/livereload", a.Fiber)
	}

	port := ":3000"
	err := a.Fiber.Listen(port)

	if err != nil {
		// If we are in watch mode, just try to kill any processes holding that port
		// and try again
		if IsDevelopment() && IsWatchMode() {
			slog.Info("Port already in use, trying to kill the process and start again")
			process.RunOrExit(fmt.Sprintf("kill -9 $(lsof -t -i%s)", port))
			time.Sleep(time.Millisecond * 50)
			err = a.Fiber.Listen(port)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}

func HtmlView(c *fiber.Ctx, page *Page) error {
	root := page.Root.Render()
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	if GetApp().LiveReload && root.tag == "html" {
		root.AppendChild(
			LiveReload(),
		)
	}

	return c.SendString(
		Render(
			root,
		),
	)
}

func PartialViewWithHeaders(c *fiber.Ctx, headers *Headers, partial *Partial) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	if partial.Headers != nil {
		for s, a := range *partial.Headers {
			c.Set(s, a)
		}
	}

	if headers != nil {
		for s, a := range *headers {
			c.Set(s, a)
		}
	}

	return c.SendString(
		Render(
			partial.Root,
		),
	)
}

func PartialView(c *fiber.Ctx, partial *Partial) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	if partial.Headers != nil {
		for s, a := range *partial.Headers {
			c.Set(s, a)
		}
	}

	return c.SendString(
		Render(
			partial.Root,
		),
	)
}
