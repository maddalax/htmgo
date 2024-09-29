package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/js"
	"htmgo-site/pages/base"
	"htmgo-site/partials"
)

func Form(ctx *h.RequestContext) *h.Page {
	return h.NewPage(base.RootPage(ctx,
		h.Div(
			h.Class("flex flex-col items-center justify-center p-4 gap-6"),
			h.H2F("Form submission with loading state example", h.Class("text-2xl font-bold")),
			h.Form(
				h.TriggerChildren(),
				h.PostPartial(partials.SubmitForm),
				h.Class("flex flex-col gap-2"),
				h.LabelFor("name", "Your Name"),
				h.Input("text",
					h.Class("p-4 rounded-md border border-slate-200"),
					h.Name("name"),
					h.Placeholder("Name")),
				SubmitButton(),
			),
		),
	))
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
			h.Class("submit", buttonClasses),
			h.Text("Submit"),
		),
	)
}

func Spinner() *h.Element {
	return h.Div(
		h.Class("absolute left-1 spinner spinner-border animate-spin inline-block w-6 h-6 border-4 rounded-full border-slate-200 border-t-transparent"),
		h.Attribute("role", "status"),
	)
}
