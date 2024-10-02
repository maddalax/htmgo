package ws

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"log/slog"
	"net/http"
	"time"
)

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)

		sessionCookie, _ := r.Cookie("session_id")

		if sessionCookie == nil {
			slog.Error("session cookie not found")
			return
		}

		locator := cc.ServiceLocator()
		manager := service.Get[SocketManager](locator)

		sessionId := sessionCookie.Value

		roomId := chi.URLParam(r, "id")

		if roomId == "" {
			slog.Error("invalid room", slog.String("room_id", roomId))
			manager.CloseWithError(sessionId, 1008, "invalid room")
			return
		}

		done := make(chan CloseEvent, 50)
		flush := make(chan bool, 50)

		manager.Add(roomId, sessionId, w, done, flush)

		defer func() {
			manager.Disconnect(sessionId)
		}()

		// Set the necessary headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*") // Optional for CORS

		// Flush the headers immediately
		flusher, ok := w.(http.Flusher)

		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				manager.Ping(sessionId)
			case <-flush:
				if flusher != nil {
					flusher.Flush()
				}
			case <-done: // Client closed the connection
				fmt.Println("Client disconnected")
				return
			}
		}
	}
}
