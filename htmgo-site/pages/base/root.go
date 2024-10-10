package base

import (
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/partials"
)

var Version = uuid.NewString()[0:6]

func RootPage(ctx *h.RequestContext, children ...h.Ren) *h.Element {
	title := "htmgo"
	description := "build simple and scalable systems with go + htmx"

	return h.Html(
		h.HxExtension(h.BaseExtensions()),
		h.Head(
			h.Meta("viewport", "width=device-width, initial-scale=1"),
			h.Meta("title", title),
			h.Link("/public/favicon.ico", "icon"),
			h.Link("/public/apple-touch-icon.png", "apple-touch-icon"),
			h.Meta("charset", "utf-8"),
			h.Meta("author", "htmgo"),
			h.Meta("description", description),
			h.Meta("og:title", title),
			h.Meta("og:url", "https://htmgo.dev"),
			h.Link("canonical", "https://htmgo.dev"),
			h.Meta("og:description", description),
			h.LinkWithVersion("/public/main.css", "stylesheet", Version),
			h.ScriptWithVersion("/public/htmgo.js", Version),
			h.Style(`
				html {
					scroll-behavior: smooth;
				}
			`),
		),
		h.Body(
			h.Class("bg-stone-50 min-h-screen overflow-x-hidden"),
			partials.NavBar(ctx, false),
			h.Fragment(children...),
		),
	)
}
