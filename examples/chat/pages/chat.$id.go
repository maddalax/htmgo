package pages

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/maddalax/htmgo/framework/js"
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
				h.Class("flex flex-row gap-4 min-h-screen bg-neutral-100"),

				// Sidebar for connected users
				UserSidebar(),

				// Chat Area
				h.Div(
					h.Class("flex flex-col flex-grow gap-4 bg-white shadow-md rounded-lg p-4"),

					h.OnEvent("hx-on::ws-after-message",
						// language=JavaScript
						js.EvalJsOnSibling("#messages", `
							element.scrollTop = element.scrollHeight;
						`)),

					// Chat Messages
					h.Div(
						h.Id("messages"),
						h.Class("flex flex-col gap-2 overflow-auto grow w-full"),
					),

					// Chat Input at the bottom
					h.Div(
						h.Class("mt-auto"),
						Form(ctx),
					),
				),
			),
		),
	)
}

func UserSidebar() *h.Element {
	return h.Div(
		h.Class("w-64 bg-slate-200 p-4 flex flex-col gap-4 rounded-l-lg"),
		h.H2F("Connected Users", h.Class("text-lg font-bold")),
		h.Ul(
			h.Class("flex flex-col gap-2"),
			// This would be populated dynamically with connected users
			h.Li(h.Text("User 1"), h.Class("text-slate-700")),
			h.Li(h.Text("User 2"), h.Class("text-slate-700")),
			h.Li(h.Text("User 3"), h.Class("text-slate-700")),
		),
	)
}

func MessageInput() *h.Element {
	return h.Input("text",
		h.Id("message-input"),
		h.Required(),
		h.Class("p-4 rounded-md border border-slate-200 w-full"),
		h.Name("message"),
		h.Placeholder("Type a message..."),
		h.HxBeforeWsSend(
			js.SetValue(""),
		),
		h.OnEvent(hx.KeyDownEvent, js.SubmitFormOnEnter()),
	)
}

func Form(ctx *h.RequestContext) *h.Element {
	return h.Div(
		h.Class("flex gap-4 items-center"),
		h.Form(
			h.Attribute("ws-send", ""),
			h.Class("flex flex-grow"),
			MessageInput(),
		),
	)
}

func Spinner(children ...h.Ren) *h.Element {
	return h.Div(
		h.Children(children...),
		h.Class("spinner spinner-border animate-spin w-4 h-4 border-2 border-t-transparent"),
		h.Attribute("role", "status"),
	)
}
