package ws

import (
	"context"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/go-chi/chi/v5"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"net/http"
)

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)

		sessionCookie, err := r.Cookie("session_id")

		cookies := r.Cookies()

		println(cookies)
		// no session
		if err != nil {
			return
		}

		c, err := websocket.Accept(w, r, nil)

		locator := cc.ServiceLocator()
		manager := service.Get[SocketManager](locator)

		if err != nil {
			return
		}

		sessionId := sessionCookie.Value

		roomId := chi.URLParam(r, "id")

		if roomId == "" {
			manager.CloseWithError(sessionId, "invalid room")
			return
		}

		manager.Add(roomId, sessionId, c)

		defer func() {
			manager.Disconnect(sessionId)
		}()

		for {
			var v map[string]any
			err = wsjson.Read(context.Background(), c, &v)
			if err != nil {
				manager.CloseWithError(sessionId, "failed to read message")
				return
			}
			if v != nil {
				manager.OnMessage(sessionId, v)
			}

		}
	}
}
