package main

import (
	"github.com/maddalax/htmgo/framework/h"
	"net/http"
)

func RenderToString(element *h.Element) string {
	return h.Render(element)
}

func RenderPage(req *http.Request, w http.ResponseWriter, page func(ctx *h.RequestContext) *h.Page) {
	ctx := h.RequestContext{
		Request:  req,
		Response: w,
	}
	h.HtmlView(w, page(&ctx))
}

func RenderPartial(req *http.Request, w http.ResponseWriter, partial func(ctx *h.RequestContext) *h.Partial) {
	ctx := h.RequestContext{
		Request:  req,
		Response: w,
	}
	h.PartialView(w, partial(&ctx))
}
