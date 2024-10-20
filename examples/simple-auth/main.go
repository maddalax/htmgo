package main

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"io/fs"
	"net/http"
	"simpleauth/__htmgo"
	"simpleauth/internal/db"
)

func main() {
	locator := service.NewLocator()

	service.Set(locator, service.Singleton, func() *db.Queries {
		return db.Provide()
	})

	h.Start(h.AppOpts{
		ServiceLocator: locator,
		LiveReload:     true,
		Register: func(app *h.App) {
			sub, err := fs.Sub(GetStaticAssets(), "assets/dist")

			if err != nil {
				panic(err)
			}

			http.FileServerFS(sub)

			app.Router.Handle("/public/*", http.StripPrefix("/public", http.FileServerFS(sub)))
			__htmgo.Register(app.Router)
		},
	})
}
