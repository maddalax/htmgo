package main

import (
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/js"
	"time"
)

func Index(ctx *h.RequestContext) *h.Page {
	return h.NewPage(
		h.Html(
			h.HxExtensions(
				h.BaseExtensions(),
			),
			h.Head(
				h.Meta("viewport", "width=device-width, initial-scale=1"),
				h.Script("/public/htmgo.js"),
			),
			h.Body(
				h.Pf("hello htmgo"),
				h.Div(
					h.Get("/current-time", "load, every 1s"),
				),
				h.Div(
					h.Button(
						h.Text("Click me"),
						h.OnClick(
							js.EvalJs(`
                 console.log("you evalulated javascript");
								 alert("you clicked me");
             `),
						),
					),
				),
			),
		),
	)
}

func CurrentTime(ctx *h.RequestContext) *h.Partial {
	return h.NewPartial(
		h.Pf("It is %s", time.Now().String()),
	)
}
