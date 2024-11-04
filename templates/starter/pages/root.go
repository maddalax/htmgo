package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/__htmgo/assets"
)

func RootPage(children ...h.Ren) *h.Page {
	title := "htmgo template"
	description := "an example of the htmgo template"
	author := "htmgo"
	url := "https://htmgo.dev"

	return h.NewPage(
		h.Html(
			h.HxExtensions(
				h.BaseExtensions(),
			),
			h.Head(
				h.Title(
					h.Text(title),
				),
				h.Meta("viewport", "width=device-width, initial-scale=1"),
				h.Link(assets.FaviconIco, "icon"),
				h.Link(assets.AppleTouchIconPng, "apple-touch-icon"),
				h.Meta("title", title),
				h.Meta("charset", "utf-8"),
				h.Meta("author", author),
				h.Meta("description", description),
				h.Meta("og:title", title),
				h.Meta("og:url", url),
				h.Link("canonical", url),
				h.Meta("og:description", description),
				h.Link(assets.MainCss, "stylesheet"),
				h.Script(assets.HtmgoJs),
			),
			h.Body(
				h.Div(
					h.Class("flex flex-col gap-2 bg-white h-full"),
					h.Fragment(children...),
				),
			),
		),
	)
}
