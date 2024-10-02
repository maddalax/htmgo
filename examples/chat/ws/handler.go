package ws

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set the necessary headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*") // Optional for CORS

		cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)
		locator := cc.ServiceLocator()
		manager := service.Get[SocketManager](locator)

		sessionCookie, _ := r.Cookie("session_id")
		sessionId := ""

		if sessionCookie != nil {
			sessionId = sessionCookie.Value
		}

		ctx := r.Context()
		done := make(chan CloseEvent, 1)
		writer := make(WriterChan, 1)

		wg := sync.WaitGroup{}
		wg.Add(1)

		/*
		 * This goroutine is responsible for writing messages to the client
		 */
		go func() {
			defer wg.Done()
			defer manager.Disconnect(sessionId)

			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case reason := <-done:
					fmt.Printf("closing connection: %s\n", reason.Reason)
					return
				case <-ticker.C:
					manager.Ping(sessionId)
				case message := <-writer:
					_, err := fmt.Fprintf(w, message)
					if err != nil {
						done <- CloseEvent{
							Code:   -1,
							Reason: err.Error(),
						}
					} else {
						flusher, ok := w.(http.Flusher)
						if ok {
							flusher.Flush()
						}
					}
				}
			}
		}()

		/**
		 * This goroutine is responsible for adding the client to the room
		 */
		wg.Add(1)
		go func() {
			defer wg.Done()
			if sessionId == "" {
				manager.writeCloseRaw(writer, "no session")
				return
			}

			roomId := chi.URLParam(r, "id")

			if roomId == "" {
				slog.Error("invalid room", slog.String("room_id", roomId))
				manager.writeCloseRaw(writer, "invalid room")
				return
			}

			manager.Add(roomId, sessionId, writer, done)
		}()

		wg.Wait()
	}
}
