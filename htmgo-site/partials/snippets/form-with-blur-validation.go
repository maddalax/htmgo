package snippets

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/hx"
)

func FormWithBlurValidation(ctx *h.RequestContext) *h.Partial {
	buttonClasses := "rounded items-center px-3 py-2 bg-slate-800 text-white w-full text-center"
	validationPath := h.GetPartialPath(
		Validate,
	)
	return h.NewPartial(
		h.Form(
			h.TriggerChildren(),
			h.Id("my-form"),
			// hx-swap: none is required so the traditional swap doesn't happen, only oob swap
			h.NoSwap(),
			h.PostPartial(SubmitFormExample),
			h.Class("flex flex-col gap-2 max-w-[300px] mx-auto"),
			h.LabelFor("name", "Your Name"),
			h.Input(
				"text",
				h.Required(),
				h.Class("p-4 rounded-md border border-slate-200"),
				h.Name("name"),
				h.Placeholder("Name"),
				h.Post(validationPath, hx.BlurEvent),
			),
			h.Div(
				h.Id("name-error"),
				h.Class("text-red-500"),
			),
			h.LabelFor("occupation", "Occupation"),
			h.Input(
				"text",
				h.Required(),
				h.Class("p-4 rounded-md border border-slate-200"),
				h.Name("occupation"),
				h.Placeholder("Software Developer"),
			),
			h.Button(
				h.Type("submit"),
				h.Class(buttonClasses),
				h.Text("Submit"),
			),
		),
	)
}

func Validate(ctx *h.RequestContext) *h.Partial {
	name := ctx.FormValue("name")

	if name == "htmgo" {
		ctx.Response.WriteHeader(400)
		return h.SwapPartial(
			ctx,
			h.Div(
				h.Id("name-error"),
				h.Text("Name is already taken"),
				h.Class("p-4 bg-rose-400 text-white rounded-md"),
			),
		)
	}

	return h.EmptyPartial()
}

func SubmitFormExample(ctx *h.RequestContext) *h.Partial {

	if !ctx.IsHttpPost() {
		return h.EmptyPartial()
	}

	validate := Validate(ctx)

	// if there is a validation error, swap it in
	if !h.IsEmptyPartial(validate) {
		return validate
	}

	// submit the form
	return h.SwapPartial(
		ctx,
		h.Div(
			h.Id("my-form"),
			h.Text("Form submitted with name: "+ctx.FormValue("name")),
		),
	)
}
