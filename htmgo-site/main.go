package main

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"htmgo-site/__htmgo"
	"htmgo-site/internal/cache"
	"htmgo-site/internal/markdown"
	"htmgo-site/internal/sitemap"
	"io/fs"
	"net/http"
)

func main() {
	locator := service.NewLocator()
	staticAssets := GetStaticAssets()
	markdownAssets := GetMarkdownAssets()

	service.Set(locator, service.Singleton, markdown.NewRenderer)
	service.Set(locator, service.Singleton, cache.NewSimpleCache)

	fmt.Printf("starting up server\n")

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {

			app.UseWithContext(func(w http.ResponseWriter, r *http.Request, context map[string]any) {
				context["embeddedMarkdown"] = markdownAssets
			})

			sub, err := fs.Sub(staticAssets, "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			app.Router.Handle("/sitemap.xml", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				s, err := sitemap.Generate(app.Router)
				if err != nil {
					http.Error(w, "failed to generate sitemap", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/xml")
				w.Write(s)
			}))

			app.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServerFS(sub)))

			__htmgo.Register(app.Router)
		},
	})
}
