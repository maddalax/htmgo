package ws

import (
	ws2 "github.com/maddalax/htmgo/extensions/websocket/opts"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createManager() *SocketManager {
	return NewSocketManager(&ws2.ExtensionOpts{
		WsPath: "/ws",
		SessionId: func(ctx *h.RequestContext) string {
			return "test"
		},
	})
}

func addSocket(manager *SocketManager, roomId string, id string) (socketId string, writer WriterChan, done DoneChan) {
	writer = make(chan string, 10)
	done = make(chan bool, 10)
	manager.Add(roomId, id, writer, done)
	return id, writer, done
}

func TestManager(t *testing.T) {
	manager := createManager()
	socketId, _, _ := addSocket(manager, "123", "456")
	socket := manager.Get(socketId)
	assert.NotNil(t, socket)
	assert.Equal(t, socketId, socket.Id)

	manager.OnClose(socketId)
	socket = manager.Get(socketId)
	assert.Nil(t, socket)
}

func TestManagerForEachSocket(t *testing.T) {
	manager := createManager()
	addSocket(manager, "all", "456")
	addSocket(manager, "all", "789")
	var count int
	manager.ForEachSocket("all", func(conn SocketConnection) {
		count++
	})
	assert.Equal(t, 2, count)
}

func TestSendText(t *testing.T) {
	manager := createManager()
	socketId, writer, done := addSocket(manager, "all", "456")
	manager.SendText(socketId, "hello")
	assert.Equal(t, "hello", <-writer)
	manager.SendText(socketId, "hello2")
	assert.Equal(t, "hello2", <-writer)
	done <- true
	assert.Equal(t, true, <-done)
}

func TestBroadcastText(t *testing.T) {
	manager := createManager()
	_, w1, d1 := addSocket(manager, "all", "456")
	_, w2, d2 := addSocket(manager, "all", "789")
	manager.BroadcastText("all", "hello", func(conn SocketConnection) bool {
		return true
	})
	assert.Equal(t, "hello", <-w1)
	assert.Equal(t, "hello", <-w2)
	d1 <- true
	d2 <- true
	assert.Equal(t, true, <-d1)
	assert.Equal(t, true, <-d2)
}

func TestBroadcastTextWithPredicate(t *testing.T) {
	manager := createManager()
	_, w1, _ := addSocket(manager, "all", "456")
	_, w2, _ := addSocket(manager, "all", "789")
	manager.BroadcastText("all", "hello", func(conn SocketConnection) bool {
		return conn.Id != "456"
	})

	assert.Equal(t, 0, len(w1))
	assert.Equal(t, 1, len(w2))
}

func TestSendHtml(t *testing.T) {
	manager := createManager()
	socketId, writer, _ := addSocket(manager, "all", "456")
	rendered := h.Render(
		h.Div(
			h.P(
				h.Text("hello"),
			),
		))
	manager.SendHtml(socketId, rendered)
	assert.Equal(t, "<div><p>hello</p></div>", <-writer)
}

func TestOnMessage(t *testing.T) {
	manager := createManager()
	socketId, _, _ := addSocket(manager, "all", "456")

	listener := make(chan SocketEvent, 10)

	manager.Listen(listener)

	manager.OnMessage(socketId, map[string]any{
		"message": "hello",
	})

	event := <-listener
	assert.Equal(t, "hello", event.Payload["message"])
	assert.Equal(t, "456", event.SessionId)
	assert.Equal(t, MessageEvent, event.Type)
	assert.Equal(t, "all", event.RoomId)
}

func TestOnClose(t *testing.T) {
	manager := createManager()
	socketId, _, _ := addSocket(manager, "all", "456")
	listener := make(chan SocketEvent, 10)
	manager.Listen(listener)
	manager.OnClose(socketId)
	event := <-listener
	assert.Equal(t, "456", event.SessionId)
	assert.Equal(t, DisconnectedEvent, event.Type)
	assert.Equal(t, "all", event.RoomId)
}

func TestOnAdd(t *testing.T) {
	manager := createManager()

	listener := make(chan SocketEvent, 10)
	manager.Listen(listener)

	socketId, _, _ := addSocket(manager, "all", "456")
	event := <-listener

	assert.Equal(t, socketId, event.SessionId)
	assert.Equal(t, ConnectedEvent, event.Type)
	assert.Equal(t, "all", event.RoomId)
}

func TestCloseWithMessage(t *testing.T) {
	manager := createManager()
	socketId, w, _ := addSocket(manager, "all", "456")
	manager.CloseWithMessage(socketId, "internal error")
	assert.Equal(t, "internal error", <-w)
	assert.Nil(t, manager.Get(socketId))
}

func TestDisconnect(t *testing.T) {
	manager := createManager()
	socketId, _, _ := addSocket(manager, "all", "456")
	manager.Disconnect(socketId)
	assert.Nil(t, manager.Get(socketId))
}

func TestPing(t *testing.T) {
	manager := createManager()
	socketId, w, _ := addSocket(manager, "all", "456")
	manager.Ping(socketId)
	assert.Equal(t, "ping", <-w)
}

func TestMultipleRooms(t *testing.T) {
	manager := createManager()
	socketId1, _, _ := addSocket(manager, "room1", "456")
	socketId2, _, _ := addSocket(manager, "room2", "789")

	room1Count := 0
	room2Count := 0

	manager.ForEachSocket("room1", func(conn SocketConnection) {
		room1Count++
	})

	manager.ForEachSocket("room2", func(conn SocketConnection) {
		room2Count++
	})

	assert.Equal(t, 1, room1Count)
	assert.Equal(t, 1, room2Count)

	room1Count = 0
	room2Count = 0

	manager.OnClose(socketId1)
	manager.OnClose(socketId2)

	manager.ForEachSocket("room1", func(conn SocketConnection) {
		room1Count++
	})

	manager.ForEachSocket("room2", func(conn SocketConnection) {
		room2Count++
	})

	assert.Equal(t, 0, room1Count)
	assert.Equal(t, 0, room2Count)
}
