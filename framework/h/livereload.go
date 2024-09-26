package h

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

var Version = uuid.NewString()

func sseHandler(w http.ResponseWriter, r *http.Request) {
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

	for {
		// Write an event to the stream
		_, err := fmt.Fprintf(w, "data: %s\n\n", Version)
		if err != nil {
			break
		}
		// Flush the response to ensure the client gets it immediately
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
}

func (app *App) AddLiveReloadHandler(path string) {
	app.Router.Get(path, sseHandler)
}
