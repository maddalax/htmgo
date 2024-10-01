package ws

import (
	"context"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/go-chi/chi/v5"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"log/slog"
	"net/http"
)

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)

		sessionCookie, _ := r.Cookie("session_id")

		c, err := websocket.Accept(w, r, nil)

		if err != nil {
			return
		}

		if sessionCookie == nil {
			slog.Error("session cookie not found")
			c.Close(websocket.StatusPolicyViolation, "no session")
			return
		}

		locator := cc.ServiceLocator()
		manager := service.Get[SocketManager](locator)

		sessionId := sessionCookie.Value

		roomId := chi.URLParam(r, "id")

		if roomId == "" {
			slog.Error("invalid room", slog.String("room_id", roomId))
			manager.CloseWithError(sessionId, websocket.StatusPolicyViolation, "invalid room")
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
				slog.Error("failed to read message", slog.String("room_id", roomId))
				manager.CloseWithError(sessionId, websocket.StatusInternalError, "failed to read message")
				return
			}
			if v != nil {
				manager.OnMessage(sessionId, v)
			}

		}
	}
}
