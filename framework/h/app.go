package h

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/util/process"
	"log/slog"
	"time"
)

type App struct {
	LiveReload bool
	Echo       *echo.Echo
}

var instance *App

func GetApp() *App {
	if instance == nil {
		panic("App instance not initialized")
	}
	return instance
}

func Start(app *echo.Echo, opts App) {
	instance = &opts
	instance.start(app)
}

func (a App) start(app *echo.Echo) {

	a.Echo = app

	if a.LiveReload {
		AddLiveReloadHandler("/dev/livereload", a.Echo)
	}

	port := ":3000"
	err := a.Echo.Start(port)

	if err != nil {
		// If we are in watch mode, just try to kill any processes holding that port
		// and try again
		if IsDevelopment() && IsWatchMode() {
			slog.Info("Port already in use, trying to kill the process and start again")
			process.RunOrExit(fmt.Sprintf("kill -9 $(lsof -t -i%s)", port))
			time.Sleep(time.Millisecond * 50)
			err = a.Echo.Start(port)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}

func HtmlView(c echo.Context, page *Page) error {
	root := page.Root.Render()
	return c.HTML(200,
		Render(
			root,
		),
	)
}

func PartialViewWithHeaders(c echo.Context, headers *Headers, partial *Partial) error {
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

	return c.HTML(200,
		Render(
			partial.Root,
		),
	)
}

func PartialView(c echo.Context, partial *Partial) error {
	c.Set(echo.HeaderContentType, echo.MIMETextHTML)
	if partial.Headers != nil {
		for s, a := range *partial.Headers {
			c.Set(s, a)
		}
	}

	return c.HTML(200,
		Render(
			partial.Root,
		),
	)
}
