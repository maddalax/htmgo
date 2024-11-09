package wsutil

import (
	"fmt"
	"github.com/maddalax/htmgo/extensions/websocket/opts"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/service"
	"github.com/puzpuzpuz/xsync/v3"
	"log/slog"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type EventType string
type WriterChan chan string
type DoneChan chan bool

const (
	ConnectedEvent    EventType = "connected"
	DisconnectedEvent EventType = "disconnected"
	MessageEvent      EventType = "message"
)

type SocketEvent struct {
	SessionId string
	RoomId    string
	Type      EventType
	Payload   map[string]any
}

type CloseEvent struct {
	Code   int
	Reason string
}

type SocketConnection struct {
	Id     string
	RoomId string
	Done   DoneChan
	Writer WriterChan
}

type ManagerMetrics struct {
	RunningGoroutines   int32
	TotalSockets        int
	TotalRooms          int
	TotalListeners      int
	SocketsPerRoomCount map[string]int
	SocketsPerRoom      map[string][]string
	TotalMessages       int64
	MessagesPerSecond   int
	SecondsElapsed      int
}

type SocketManager struct {
	sockets           *xsync.MapOf[string, *xsync.MapOf[string, SocketConnection]]
	idToRoom          *xsync.MapOf[string, string]
	listeners         []chan SocketEvent
	goroutinesRunning atomic.Int32
	opts              *opts.ExtensionOpts
	lock              sync.Mutex
	totalMessages     atomic.Int64
	messagesPerSecond int
	secondsElapsed    int
}

func (manager *SocketManager) StartMetrics() {
	go func() {
		for {
			time.Sleep(time.Second)
			manager.lock.Lock()
			manager.secondsElapsed++
			totalMessages := manager.totalMessages.Load()
			manager.messagesPerSecond = int(float64(totalMessages) / float64(manager.secondsElapsed))
			manager.lock.Unlock()
		}
	}()
}

func (manager *SocketManager) Metrics() ManagerMetrics {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	count := manager.goroutinesRunning.Load()
	metrics := ManagerMetrics{
		RunningGoroutines:   count,
		TotalSockets:        0,
		TotalRooms:          0,
		TotalListeners:      len(manager.listeners),
		SocketsPerRoom:      make(map[string][]string),
		SocketsPerRoomCount: make(map[string]int),
		TotalMessages:       manager.totalMessages.Load(),
		MessagesPerSecond:   manager.messagesPerSecond,
		SecondsElapsed:      manager.secondsElapsed,
	}

	roomMap := make(map[string]int)

	manager.idToRoom.Range(func(socketId string, roomId string) bool {
		roomMap[roomId]++
		return true
	})

	metrics.TotalRooms = len(roomMap)

	manager.sockets.Range(func(roomId string, sockets *xsync.MapOf[string, SocketConnection]) bool {
		metrics.SocketsPerRoomCount[roomId] = sockets.Size()
		sockets.Range(func(socketId string, conn SocketConnection) bool {
			if metrics.SocketsPerRoom[roomId] == nil {
				metrics.SocketsPerRoom[roomId] = []string{}
			}
			metrics.SocketsPerRoom[roomId] = append(metrics.SocketsPerRoom[roomId], socketId)
			metrics.TotalSockets++
			return true
		})
		return true
	})

	return metrics
}

func SocketManagerFromCtx(ctx *h.RequestContext) *SocketManager {
	locator := ctx.ServiceLocator()
	return service.Get[SocketManager](locator)
}

func NewSocketManager(opts *opts.ExtensionOpts) *SocketManager {
	return &SocketManager{
		sockets:           xsync.NewMapOf[string, *xsync.MapOf[string, SocketConnection]](),
		idToRoom:          xsync.NewMapOf[string, string](),
		opts:              opts,
		goroutinesRunning: atomic.Int32{},
	}
}

func (manager *SocketManager) ForEachSocket(roomId string, cb func(conn SocketConnection)) {
	sockets, ok := manager.sockets.Load(roomId)
	if !ok {
		return
	}
	sockets.Range(func(id string, conn SocketConnection) bool {
		cb(conn)
		return true
	})
}

func (manager *SocketManager) RunIntervalWithSocket(socketId string, interval time.Duration, cb func() bool) {
	socketIdSlog := slog.String("socketId", socketId)
	slog.Debug("ws-extension: starting every loop", socketIdSlog, slog.Duration("duration", interval))

	go func() {
		manager.goroutinesRunning.Add(1)
		defer manager.goroutinesRunning.Add(-1)
		tries := 0
		for {
			socket := manager.Get(socketId)
			// This can run before the socket is established, lets try a few times and kill it if socket isn't connected after a bit.
			if socket == nil {
				if tries > 200 {
					slog.Debug("ws-extension: socket disconnected, killing goroutine", socketIdSlog)
					return
				} else {
					time.Sleep(time.Millisecond * 15)
					tries++
					slog.Debug("ws-extension: socket not connected yet, trying again", socketIdSlog, slog.Int("attempt", tries))
					continue
				}
			}
			success := cb()
			if !success {
				return
			}
			time.Sleep(interval)
		}
	}()
}

func (manager *SocketManager) Listen(listener chan SocketEvent) {
	if manager.listeners == nil {
		manager.listeners = make([]chan SocketEvent, 0)
	}
	if listener != nil {
		manager.listeners = append(manager.listeners, listener)
	}
}

func (manager *SocketManager) dispatch(event SocketEvent) {
	done := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-time.After(5 * time.Second):
				fmt.Printf("havent dispatched event after 5s, chan blocked: %s\n", event.Type)
			}
		}
	}()
	for _, listener := range manager.listeners {
		listener <- event
	}
	done <- struct{}{}
}

func (manager *SocketManager) OnMessage(id string, message map[string]any) {
	socket := manager.Get(id)
	if socket == nil {
		return
	}

	manager.totalMessages.Add(1)
	manager.dispatch(SocketEvent{
		SessionId: id,
		Type:      MessageEvent,
		Payload:   message,
		RoomId:    socket.RoomId,
	})
}

func (manager *SocketManager) Add(roomId string, id string, writer WriterChan, done DoneChan) {
	manager.idToRoom.Store(id, roomId)

	sockets, ok := manager.sockets.LoadOrCompute(roomId, func() *xsync.MapOf[string, SocketConnection] {
		return xsync.NewMapOf[string, SocketConnection]()
	})

	sockets.Store(id, SocketConnection{
		Id:     id,
		Writer: writer,
		RoomId: roomId,
		Done:   done,
	})

	s, ok := sockets.Load(id)
	if !ok {
		return
	}

	manager.dispatch(SocketEvent{
		SessionId: s.Id,
		Type:      ConnectedEvent,
		RoomId:    s.RoomId,
		Payload:   map[string]any{},
	})
}

func (manager *SocketManager) OnClose(id string) {
	socket := manager.Get(id)
	if socket == nil {
		return
	}
	slog.Debug("ws-extension: removing socket from manager", slog.String("socketId", id))
	manager.dispatch(SocketEvent{
		SessionId: id,
		Type:      DisconnectedEvent,
		RoomId:    socket.RoomId,
		Payload:   map[string]any{},
	})
	roomId, ok := manager.idToRoom.Load(id)
	if !ok {
		return
	}
	sockets, ok := manager.sockets.Load(roomId)
	if !ok {
		return
	}
	sockets.Delete(id)
	manager.idToRoom.Delete(id)
	slog.Debug("ws-extension: removed socket from manager", slog.String("socketId", id))

}

func (manager *SocketManager) CloseWithMessage(id string, message string) {
	conn := manager.Get(id)
	if conn != nil {
		defer manager.OnClose(id)
		manager.writeText(*conn, message)
		conn.Done <- true
	}
}

func (manager *SocketManager) Disconnect(id string) {
	conn := manager.Get(id)
	if conn != nil {
		manager.OnClose(id)
		conn.Done <- true
	}
}

func (manager *SocketManager) Get(id string) *SocketConnection {
	roomId, ok := manager.idToRoom.Load(id)
	if !ok {
		return nil
	}
	sockets, ok := manager.sockets.Load(roomId)
	if !ok {
		return nil
	}
	conn, ok := sockets.Load(id)
	return &conn
}

func (manager *SocketManager) Ping(id string) bool {
	conn := manager.Get(id)
	if conn != nil {
		return manager.writeText(*conn, "ping")
	}
	return false
}

func (manager *SocketManager) writeCloseRaw(writer WriterChan, message string) {
	manager.writeTextRaw(writer, message)
}

func (manager *SocketManager) writeTextRaw(writer WriterChan, message string) {
	timeout := 3 * time.Second
	select {
	case writer <- message:
	case <-time.After(timeout):
		fmt.Printf("could not send %s to channel after %s\n", message, timeout)
	}
}

func (manager *SocketManager) writeText(socket SocketConnection, message string) bool {
	if socket.Writer == nil {
		return false
	}
	manager.writeTextRaw(socket.Writer, message)
	return true
}

func (manager *SocketManager) BroadcastText(roomId string, message string, predicate func(conn SocketConnection) bool) {
	sockets, ok := manager.sockets.Load(roomId)

	if !ok {
		return
	}

	sockets.Range(func(id string, conn SocketConnection) bool {
		if predicate(conn) {
			manager.writeText(conn, message)
		}
		return true
	})
}

func (manager *SocketManager) SendHtml(id string, message string) bool {
	conn := manager.Get(id)
	minified := strings.ReplaceAll(message, "\n", "")
	minified = strings.ReplaceAll(minified, "\t", "")
	minified = strings.TrimSpace(minified)
	if conn != nil {
		return manager.writeText(*conn, minified)
	}
	return false
}

func (manager *SocketManager) SendText(id string, message string) bool {
	conn := manager.Get(id)
	if conn != nil {
		return manager.writeText(*conn, message)
	}
	return false
}
