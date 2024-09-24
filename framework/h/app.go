package h

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/maddalax/htmgo/framework/internal/process"
	"github.com/maddalax/htmgo/framework/service"
	"log/slog"
	"os/exec"
	"runtime"
	"time"
)

type RequestContext struct {
	echo.Context
	locator           *service.Locator
	isBoosted         bool
	currentBrowserUrl string
	hxPromptResponse  string
	isHxRequest       bool
	hxTargetId        string
	hxTriggerName     string
	hxTriggerId       string
}

func (c *RequestContext) ServiceLocator() *service.Locator {
	return c.locator
}

type AppOpts struct {
	LiveReload     bool
	ServiceLocator *service.Locator
	Register       func(echo *echo.Echo)
}

type App struct {
	Opts AppOpts
	Echo *echo.Echo
}

var instance *App

func GetApp() *App {
	if instance == nil {
		panic("App instance not initialized")
	}
	return instance
}

func Start(opts AppOpts) {
	e := echo.New()
	instance := App{
		Opts: opts,
		Echo: e,
	}
	instance.start()
}

func populateHxFields(cc *RequestContext) {
	cc.isBoosted = cc.Request().Header.Get(hx.BoostedHeader) == "true"
	cc.currentBrowserUrl = cc.Request().Header.Get(hx.CurrentUrlHeader)
	cc.hxPromptResponse = cc.Request().Header.Get(hx.PromptResponseHeader)
	cc.isHxRequest = cc.Request().Header.Get(hx.RequestHeader) == "true"
	cc.hxTargetId = cc.Request().Header.Get(hx.TargetIdHeader)
	cc.hxTriggerName = cc.Request().Header.Get(hx.TriggerNameHeader)
	cc.hxTriggerId = cc.Request().Header.Get(hx.TriggerIdHeader)
}

func (a App) start() {

	if a.Opts.Register != nil {
		a.Opts.Register(a.Echo)
	}

	a.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &RequestContext{
				Context: c,
				locator: a.Opts.ServiceLocator,
			}
			populateHxFields(cc)
			return next(cc)
		}
	})

	if a.Opts.LiveReload && IsDevelopment() {
		AddLiveReloadHandler("/dev/livereload", a.Echo)
	}

	port := ":3000"
	err := a.Echo.Start(port)

	if err != nil {
		// If we are in watch mode, just try to kill any processes holding that port
		// and try again
		if IsDevelopment() && IsWatchMode() {
			slog.Info("Port already in use, trying to kill the process and start again")
			if runtime.GOOS == "windows" {
				cmd := exec.Command("cmd", "/C", fmt.Sprintf(`for /F "tokens=5" %%i in ('netstat -aon ^| findstr :%s') do taskkill /F /PID %%i`, port))
				cmd.Run()
			} else {
				process.RunOrExit(fmt.Sprintf("kill -9 $(lsof -t -i%s)", port))
			}

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
	root := page.Root
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
			c.Response().Header().Set(s, a)
		}
	}

	return c.HTML(200,
		Render(
			partial,
		),
	)
}

func PartialView(c echo.Context, partial *Partial) error {
	c.Set(echo.HeaderContentType, echo.MIMETextHTML)
	if partial.Headers != nil {
		for s, a := range *partial.Headers {
			c.Response().Header().Set(s, a)
		}
	}

	return c.HTML(200,
		Render(
			partial,
		),
	)
}
