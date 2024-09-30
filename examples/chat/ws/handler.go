package ws

import (
	"context"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"net/http"
)

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)
		locator := cc.ServiceLocator()
		manager := service.Get[SocketManager](locator)

		if err != nil {
			return
		}

		id := uuid.NewString()
		manager.Add(id, c)

		defer func() {
			manager.Disconnect(id)
		}()

		for {
			var v map[string]any
			err = wsjson.Read(context.Background(), c, &v)
			if err != nil {
				manager.CloseWithError(id, "failed to read message")
				return
			}
			if v != nil {
				manager.OnMessage(id, v)
			}

		}
	}
}
