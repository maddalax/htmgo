package wsutil

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	ws2 "github.com/maddalax/htmgo/extensions/websocket/opts"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

func WsHttpHandler(opts *ws2.ExtensionOpts) http.HandlerFunc {

	if opts.RoomName == nil {
		opts.RoomName = func(ctx *h.RequestContext) string {
			return "all"
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		cc := r.Context().Value(h.RequestContextKey).(*h.RequestContext)
		locator := cc.ServiceLocator()
		manager := service.Get[SocketManager](locator)

		sessionId := opts.SessionId(cc)

		if sessionId == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			slog.Info("failed to upgrade", slog.String("error", err.Error()))
			return
		}

		roomId := opts.RoomName(cc)
		/*
			Large buffer in case the client disconnects while we are writing
			we don't want to block the writer
		*/
		done := make(chan bool, 1000)
		writer := make(WriterChan, 1000)

		wg := sync.WaitGroup{}

		manager.Add(roomId, sessionId, writer, done)

		/*
		 * This goroutine is responsible for writing messages to the client
		 */
		wg.Add(1)
		go func() {
			defer manager.Disconnect(sessionId)
			defer wg.Done()

			defer func() {
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
