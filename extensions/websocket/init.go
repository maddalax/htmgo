package websocket

import (
	"github.com/maddalax/htmgo/extensions/websocket/internal/wsutil"
	"github.com/maddalax/htmgo/extensions/websocket/opts"
	"github.com/maddalax/htmgo/extensions/websocket/ws"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
)

func EnableExtension(app *h.App, opts opts.ExtensionOpts) {
	if app.Opts.ServiceLocator == nil {
		app.Opts.ServiceLocator = service.NewLocator()
	}

	if opts.WsPath == "" {
		panic("websocket: WsPath is required")
	}

	if opts.SessionId == nil {
		panic("websocket: SessionId func is required")
	}

	service.Set[wsutil.SocketManager](app.Opts.ServiceLocator, service.Singleton, func() *wsutil.SocketManager {
		manager := wsutil.NewSocketManager(&opts)
		manager.StartMetrics()
		return manager
	})
	ws.StartListener(app.Opts.ServiceLocator)
	app.Router.Handle(opts.WsPath, wsutil.WsHttpHandler(&opts))
}
