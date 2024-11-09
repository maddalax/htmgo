package ws

import (
	"fmt"
	"github.com/maddalax/htmgo/extensions/websocket/session"
	"github.com/maddalax/htmgo/extensions/websocket/ws"
	"github.com/maddalax/htmgo/framework/h"
	"runtime"
	"time"
	"ws-example/pages"
)

func Metrics(ctx *h.RequestContext) *h.Page {

	ws.RunOnConnected(ctx, func() {
		ws.PushElementCtx(
			ctx,
			metricsView(ctx),
		)
	})

	ws.Every(ctx, time.Second, func() bool {
		return ws.PushElementCtx(
			ctx,
			metricsView(ctx),
		)
	})

	return h.NewPage(
		pages.RootPage(
			ctx,
			h.Div(
				h.Attribute("ws-connect", fmt.Sprintf("/ws?sessionId=%s", session.GetSessionId(ctx))),
				h.Class("flex flex-col gap-4 items-center min-h-screen max-w-2xl mx-auto mt-8"),
				h.H3(
					h.Id("intro-text"),
					h.Text("Websocket Metrics"),
					h.Class("text-2xl"),
				),
				h.Div(
					h.Id("ws-metrics"),
				),
			),
		),
	)
}

func metricsView(ctx *h.RequestContext) *h.Element {
	metrics := ws.MetricsFromCtx(ctx)

	return h.Div(
		h.Id("ws-metrics"),
		List(metrics),
	)
}

func List(metrics ws.Metrics) *h.Element {
	return h.Body(
		h.Div(
			h.Class("flow-root rounded-lg border border-gray-100 py-3 shadow-sm"),
			h.Dl(
				h.Class("-my-3 divide-y divide-gray-100 text-sm"),
				ListItem("Current Time", time.Now().Format("15:04:05")),
				ListItem("Seconds Elapsed", fmt.Sprintf("%d", metrics.Manager.SecondsElapsed)),
				ListItem("Total Messages", fmt.Sprintf("%d", metrics.Manager.TotalMessages)),
				ListItem("Messages Per Second", fmt.Sprintf("%d", metrics.Manager.MessagesPerSecond)),
				ListItem("Total Goroutines For ws.Every", fmt.Sprintf("%d", metrics.Manager.RunningGoroutines)),
				ListItem("Total Goroutines In System", fmt.Sprintf("%d", runtime.NumGoroutine())),
				ListItem("Sockets", fmt.Sprintf("%d", metrics.Manager.TotalSockets)),
				ListItem("Rooms", fmt.Sprintf("%d", metrics.Manager.TotalRooms)),
				ListItem("Session Id To Hashes", fmt.Sprintf("%d", metrics.Handler.SessionIdToHashesCount)),
				ListItem("Total Handlers", fmt.Sprintf("%d", metrics.Handler.TotalHandlers)),
				ListItem("Server Event Names To Hash", fmt.Sprintf("%d", metrics.Handler.ServerEventNamesToHashCount)),
				ListItem("Total Listeners", fmt.Sprintf("%d", metrics.Manager.TotalListeners)),
				h.IterMap(metrics.Manager.SocketsPerRoom, func(key string, value []string) *h.Element {
					return ListBlock(
						fmt.Sprintf("Sockets In Room - %s", key),
						h.IfElse(
							len(value) > 100,
							h.Div(
								h.Pf("%d total sockets", len(value)),
							),
							h.Div(
								h.List(value, func(item string, index int) *h.Element {
									return h.Div(
										h.Pf("%s", item),
									)
								}),
							),
						),
					)
				}),
			),
		),
	)
}

func ListItem(term, description string) *h.Element {
	return h.Div(
		h.Class("grid grid-cols-1 gap-1 p-3 even:bg-gray-50 sm:grid-cols-3 sm:gap-4"),
		DescriptionTerm(term),
		DescriptionDetail(description),
	)
}

func ListBlock(title string, children *h.Element) *h.Element {
	return h.Div(
		h.Class("grid grid-cols-1 gap-1 p-3 even:bg-gray-50 sm:grid-cols-3 sm:gap-4"),
		DescriptionTerm(title),
		h.Dd(
			h.Class("text-gray-700 sm:col-span-2"),
			children,
		),
	)
}

func DescriptionTerm(term string) *h.Element {
	return h.Dt(
		h.Class("font-medium text-gray-900"),
		h.Text(term),
	)
}

func DescriptionDetail(detail string) *h.Element {
	return h.Dd(
		h.Class("text-gray-700 sm:col-span-2"),
		h.Text(detail),
	)
}
