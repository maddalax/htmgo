package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/maddalax/htmgo/framework/js"
)

func IndexPage(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		RootPage(
			h.Div(
				h.JoinExtensions(
					h.TriggerChildren(),
					h.HxExtension("ws"),
				),
				h.Attribute("ws-connect", "/chat"),
				h.Class("flex flex-col gap-4 items-center pt-24 min-h-screen bg-neutral-100"),
				Form(ctx),
				h.Div(
					h.Div(
						h.Id("messages"),
						h.Class("flex flex-col gap-2 w-full"),
					),
				),
			),
		),
	)
}

func MessageInput() *h.Element {
	return h.Input("text",
		h.Id("message-input"),
		h.Required(),
		h.Class("p-4 rounded-md border border-slate-200"),
		h.Name("message"),
		h.Placeholder("Message"),
		h.HxBeforeWsSend(
			js.SetValue(""),
		),
		h.OnEvent(hx.KeyDownEvent, js.SubmitFormOnEnter()),
	)
}

func Form(ctx *h.RequestContext) *h.Element {
	return h.Div(
		h.Class("flex flex-col items-center justify-center p-4 gap-6"),
		h.H2F("Form submission with ws example", h.Class("text-2xl font-bold")),
		h.Form(
			h.Attribute("ws-send", ""),
			h.Class("flex flex-col gap-2"),
			h.LabelFor("name", "Your Message"),
			MessageInput(),
			SubmitButton(),
		),
	)
}

func SubmitButton() *h.Element {
	buttonClasses := "rounded items-center px-3 py-2 bg-slate-800 text-white w-full text-center"
	return h.Div(
		h.HxBeforeRequest(
			js.RemoveClassOnChildren(".loading", "hidden"),
			js.SetClassOnChildren(".submit", "hidden"),
		),
		h.HxAfterRequest(
			js.SetClassOnChildren(".loading", "hidden"),
			js.RemoveClassOnChildren(".submit", "hidden"),
		),
		h.Class("flex gap-2 justify-center"),
		h.Button(
			h.Class("loading hidden relative text-center", buttonClasses),
			Spinner(),
			h.Disabled(),
			h.Text("Submitting..."),
		),
		h.Button(
			h.Type("submit"),
			h.Class("submit", buttonClasses),
			h.Text("Submit"),
		),
	)
}

func Spinner(children ...h.Ren) *h.Element {
	return h.Div(
		h.Children(children...),
		h.Class("absolute left-1 spinner spinner-border animate-spin inline-block w-6 h-6 border-4 rounded-full border-slate-200 border-t-transparent"),
		h.Attribute("role", "status"),
	)
}
