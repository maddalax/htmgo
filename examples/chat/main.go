package main

import (
	"chat/__htmgo"
	"chat/chat"
	"chat/internal/db"
	"chat/ws"
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"io/fs"
	"net/http"
	"runtime"
	"time"
)

func main() {
	locator := service.NewLocator()

	service.Set[db.Queries](locator, service.Singleton, db.Provide)
	service.Set[ws.SocketManager](locator, service.Singleton, func() *ws.SocketManager {
		return ws.NewSocketManager()
	})

	chatManager := chat.NewManager(locator)
	go chatManager.StartListener()

	go func() {
		for {
			count := runtime.NumGoroutine()
			fmt.Printf("goroutines: %d\n", count)
			time.Sleep(10 * time.Second)
		}
	}()

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
			app.Router.Handle("/ws/chat/{id}", ws.Handle())

			__htmgo.Register(app.Router)
		},
	})
}
