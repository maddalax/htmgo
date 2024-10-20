package ui

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/js"
)

func SubmitButton(submitText string) *h.Element {
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
			spinner(),
			h.Disabled(),
			h.Text("Submitting..."),
		),
		h.Button(
			h.Type("submit"),
			h.Class("submit", buttonClasses),
			h.Text(submitText),
		),
	)
}

func spinner(children ...h.Ren) *h.Element {
	return h.Div(
		h.Children(children...),
		h.Class("absolute left-1 spinner spinner-border animate-spin inline-block w-6 h-6 border-4 rounded-full border-slate-200 border-t-transparent"),
		h.Attribute("role", "status"),
	)
}
