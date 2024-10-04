package h

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/maddalax/htmgo/framework/service"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type RequestContext struct {
	Request           *http.Request
	Response          http.ResponseWriter
	locator           *service.Locator
	isBoosted         bool
	currentBrowserUrl string
	hxPromptResponse  string
	isHxRequest       bool
	hxTargetId        string
	hxTriggerName     string
	hxTriggerId       string
	kv                map[string]interface{}
}

func (c *RequestContext) FormValue(key string) string {
	return c.Request.FormValue(key)
}

func (c *RequestContext) QueryParam(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *RequestContext) Set(key string, value interface{}) {
	if c.kv == nil {
		c.kv = make(map[string]interface{})
	}
	c.kv[key] = value
}

func (c *RequestContext) Get(key string) interface{} {
	if c.kv == nil {
		return nil
	}
	return c.kv[key]
}

func (c *RequestContext) ServiceLocator() *service.Locator {
	return c.locator
}

type AppOpts struct {
	LiveReload     bool
	ServiceLocator *service.Locator
	Register       func(app *App)
}

type App struct {
	Opts   AppOpts
	Router *chi.Mux
}

func Start(opts AppOpts) {
	router := chi.NewRouter()
	instance := App{
		Opts:   opts,
		Router: router,
	}
	instance.start()
}

const RequestContextKey = "htmgo.request.context"

func populateHxFields(cc *RequestContext) {
	cc.isBoosted = cc.Request.Header.Get(hx.BoostedHeader) == "true"
	cc.isBoosted = cc.Request.Header.Get(hx.BoostedHeader) == "true"
	cc.currentBrowserUrl = cc.Request.Header.Get(hx.CurrentUrlHeader)
	cc.hxPromptResponse = cc.Request.Header.Get(hx.PromptResponseHeader)
	cc.isHxRequest = cc.Request.Header.Get(hx.RequestHeader) == "true"
	cc.hxTargetId = cc.Request.Header.Get(hx.TargetIdHeader)
	cc.hxTriggerName = cc.Request.Header.Get(hx.TriggerNameHeader)
	cc.hxTriggerId = cc.Request.Header.Get(hx.TriggerIdHeader)
}

func (app *App) UseWithContext(h func(w http.ResponseWriter, r *http.Request, context map[string]any)) {
	app.Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cc := r.Context().Value(RequestContextKey).(*RequestContext)
			h(w, r, cc.kv)
			handler.ServeHTTP(w, r)
		})
	})
}

func GetLogLevel() slog.Level {
	// Get the log level from the environment variable
	logLevel := os.Getenv("LOG_LEVEL")
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		// Default to INFO if no valid log level is set
		return slog.LevelInfo
	}
}

func (app *App) start() {

	slog.SetLogLoggerLevel(GetLogLevel())

	app.Router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cc := &RequestContext{
				locator:  app.Opts.ServiceLocator,
				Request:  r,
				Response: w,
				kv:       make(map[string]interface{}),
			}
			populateHxFields(cc)
			ctx := context.WithValue(r.Context(), RequestContextKey, cc)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	if app.Opts.Register != nil {
		app.Opts.Register(app)
	}

	if app.Opts.LiveReload && IsDevelopment() {
		app.AddLiveReloadHandler("/dev/livereload")
	}

	port := ":3000"
	slog.Info(fmt.Sprintf("Server started on port %s", port))
	err := http.ListenAndServe(port, app.Router)

	if err != nil {
		// If we are in watch mode, just try to kill any processes holding that port
		// and try again
		if IsDevelopment() && IsWatchMode() {
			slog.Info("Port already in use, trying to kill the process and start again")
			if runtime.GOOS == "windows" {
				cmd := exec.Command("cmd", "/C", fmt.Sprintf(`for /F "tokens=5" %%i in ('netstat -aon ^| findstr :%s') do taskkill /F /PID %%i`, port))
				cmd.Run()
			} else {
				cmd := exec.Command("bash", "-c", fmt.Sprintf("kill -9 $(lsof -ti%s)", port))
				cmd.Run()
			}
			time.Sleep(time.Millisecond * 50)
			err = http.ListenAndServe(":3000", app.Router)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
		panic(err)
	}
}

func writeHtml(w http.ResponseWriter, element Ren) error {
	w.Header().Set("Content-Type", "text/html")
	_, err := fmt.Fprint(w, Render(element))
	return err
}

func HtmlView(w http.ResponseWriter, page *Page) error {
	return writeHtml(w, page.Root)
}

func PartialViewWithHeaders(w http.ResponseWriter, headers *Headers, partial *Partial) error {
	if partial.Headers != nil {
		for s, a := range *partial.Headers {
			w.Header().Set(s, a)
		}
	}

	if headers != nil {
		for s, a := range *headers {
			w.Header().Set(s, a)
		}
	}

	return writeHtml(w, partial.Root)
}

func PartialView(w http.ResponseWriter, partial *Partial) error {
	if partial.Headers != nil {
		for s, a := range *partial.Headers {
			w.Header().Set(s, a)
		}
	}

	return writeHtml(w, partial.Root)
}
