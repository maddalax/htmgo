package pages

import (
	"chat/chat"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/js"
	"time"
)

func ChatRoom(ctx *h.RequestContext) *h.Page {
	roomId := chi.URLParam(ctx.Request, "id")
	return h.NewPage(
		RootPage(
			h.Div(
				h.JoinExtensions(
					h.TriggerChildren(),
					h.HxExtension("ws"),
				),

				h.Attribute("ws-connect", fmt.Sprintf("/ws/chat/%s", roomId)),

				h.HxOnWsOpen(
					js.ConsoleLog("Connected to chat room"),
				),

				h.HxOnWsClose(
					js.EvalJs(fmt.Sprintf(`
						const reason = e.detail.event.reason
						if(['invalid room', 'no session', 'invalid user'].includes(reason)) {
							window.location.href = '/?roomId=%s';
						} else if(e.detail.event.code === 1011) { 
							window.location.reload()
						} else if (e.detail.event.code === 1008 || e.detail.event.code === 1006) {
                            window.location.href = '/?roomId=%s';
						} else {
							console.error('Connection closed:', e.detail.event)
						}
					`, roomId, roomId)),
				),

				h.Class("flex flex-row min-h-screen bg-neutral-100"),

				// Sidebar for connected users
				UserSidebar(),

				h.Div(
					h.Class("flex flex-col flex-grow bg-white rounded p-4"),

					// Room name at the top, fixed
					CachedRoomHeader(ctx),

					// Padding to push chat content below the fixed room name
					h.Div(h.Class("pt-[50px]")),

					h.HxAfterWsMessage(
						js.EvalJsOnSibling("#messages",
							`element.scrollTop = element.scrollHeight;`),
					),

					// Chat Messages
					h.Div(
						h.Id("messages"),
						h.Class("flex flex-col gap-4 overflow-auto grow w-full"),
					),

					// Chat Input at the bottom
					h.Div(
						h.Class("mt-auto"),
						Form(),
					),
				),
			),
		),
	)
}

var CachedRoomHeader = h.CachedPerKeyT(time.Hour, func(ctx *h.RequestContext) (string, h.GetElementFunc) {
	roomId := chi.URLParam(ctx.Request, "id")
	return roomId, func() *h.Element {
		return roomNameHeader(ctx)
	}
})

func roomNameHeader(ctx *h.RequestContext) *h.Element {
	roomId := chi.URLParam(ctx.Request, "id")
	service := chat.NewService(ctx.ServiceLocator())
	room, err := service.GetRoom(roomId)
	if err != nil {
		return h.Div()
	}
	return h.Div(
		h.Class("bg-neutral-700 text-white p-3 shadow-sm w-full fixed top-0 left-0 flex justify-center z-10"),
		h.H2F(room.Name, h.Class("text-lg font-bold")),
		h.Div(
			h.Class("absolute right-5 top-3 cursor-pointer"),
			h.Text("Share"),
			h.OnClick(
				js.EvalJs(`
						alert("Share this url with your friends:\n " + window.location.href)
					`),
			),
		),
	)
}

func UserSidebar() *h.Element {
	return h.Div(
		h.Class("pt-[67px] min-w-48 w-48 bg-neutral-200 p-4 flex flex-col justify-between gap-3 rounded-l-lg"),
		h.Div(
			h.H3F("Connected Users", h.Class("text-lg font-bold")),
			chat.ConnectedUsers("", false),
		),
		h.A(
			h.Class("cursor-pointer"),
			h.Href("/"),
			h.Text("Leave Room"),
		),
	)
}

func MessageInput() *h.Element {
	return h.Input("text",
		h.Id("message-input"),
		h.Required(),
		h.Class("p-4 rounded-md border border-slate-200 w-full focus:outline-none focus:ring focus:ring-slate-200"),
		h.Name("message"),
		h.MaxLength(1000),
		h.Placeholder("Type a message..."),
		h.HxAfterWsSend(
			js.SetValue(""),
		),
	)
}

func Form() *h.Element {
	return h.Div(
		h.Class("flex gap-4 items-center"),
		h.Form(
			h.Attribute("ws-send", ""),
			h.Class("flex flex-grow"),
			MessageInput(),
		),
	)
}
