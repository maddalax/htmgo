package sse

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

func HandleWs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)
		locator := cc.ServiceLocator()
		manager := service.Get[SocketManager](locator)

		sessionCookie, _ := r.Cookie("state")
		sessionId := ""

		if sessionCookie != nil {
			sessionId = sessionCookie.Value
		}

		if sessionId == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			slog.Info("failed to upgrade", slog.String("error", err.Error()))
			return
		}
		/*
			Large buffer in case the client disconnects while we are writing
			we don't want to block the writer
		*/
		done := make(chan bool, 1000)
		writer := make(WriterChan, 1000)

		wg := sync.WaitGroup{}

		manager.Add("all", sessionId, writer, done)

		/*
		 * This goroutine is responsible for writing messages to the client
		 */
		wg.Add(1)
		go func() {
			defer manager.Disconnect(sessionId)
			defer wg.Done()

			defer func() {
				fmt.Printf("empting channels\n")
				for len(writer) > 0 {
					<-writer
				}
				for len(done) > 0 {
					<-done
				}
			}()

			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-done:
					fmt.Printf("closing connection: \n")
					return
				case <-ticker.C:
					manager.Ping(sessionId)
				case message := <-writer:
					err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(message))
					if err != nil {
						return
					}
				}
			}
		}()

		/*
		 * This goroutine is responsible for reading messages from the client
		 */
		go func() {
			defer conn.Close()
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					return
				}
				if op != ws.OpText {
					return
				}
				m := make(map[string]any)
				err = json.Unmarshal(msg, &m)
				if err != nil {
					return
				}
				manager.OnMessage(sessionId, m)
			}
		}()

		wg.Wait()
	}
}

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

		sessionCookie, _ := r.Cookie("state")
		sessionId := ""

		if sessionCookie != nil {
			sessionId = sessionCookie.Value
		}

		ctx := r.Context()

		/*
			Large buffer in case the client disconnects while we are writing
			we don't want to block the writer
		*/
		done := make(chan bool, 1000)
		writer := make(WriterChan, 1000)

		wg := sync.WaitGroup{}
		wg.Add(1)

		/*
		 * This goroutine is responsible for writing messages to the client
		 */
		go func() {
			defer wg.Done()
			defer manager.Disconnect(sessionId)

			defer func() {
				fmt.Printf("empting channels\n")
				for len(writer) > 0 {
					<-writer
				}
				for len(done) > 0 {
					<-done
				}
			}()

			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-done:
					fmt.Printf("closing connection: \n")
					return
				case <-ticker.C:
					manager.Ping(sessionId)
				case message := <-writer:
					_, err := fmt.Fprintf(w, message)
					if err != nil {
						done <- true
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

			manager.Add("all", sessionId, writer, done)
		}()

		wg.Wait()
	}
}
