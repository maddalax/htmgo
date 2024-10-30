package base

import (
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/partials"
)

var Version = uuid.NewString()[0:6]

type RootPageProps struct {
	Title       string
	Description string
	Canonical   string
	Children    h.Ren
}

func ConfigurableRootPage(ctx *h.RequestContext, props RootPageProps) *h.Page {
	title := "htmgo"
	description := "build simple and scalable systems with go + htmx"
	canonical := ctx.Request.URL.String()

	if props.Canonical != "" {
		canonical = props.Canonical
	}

	if props.Title != "" {
		title = props.Title
	}

	if props.Description != "" {
		description = props.Description
	}

	return h.NewPage(
		h.Html(
			h.HxExtension(
				h.BaseExtensions(),
			),
			h.Head(
				h.Meta("viewport", "width=device-width, initial-scale=1"),
				h.Meta("title", title),
				h.Title(
					h.Text(title),
				),
				h.Link("/public/favicon.ico", "icon"),
				h.Link("/public/apple-touch-icon.png", "apple-touch-icon"),
				h.Meta("charset", "utf-8"),
				h.Meta("author", "htmgo"),
				h.Meta("description", description),
				h.Meta("og:title", title),
				h.Meta("og:url", ctx.Request.URL.String()),
				h.Link("canonical", canonical),
				h.Link("https://cdn.jsdelivr.net/npm/@docsearch/css@3", "stylesheet"),
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
				h.Class("bg-white h-screen"),
				props.Children,
				h.Script("https://cdn.jsdelivr.net/npm/@docsearch/js@3"),
				h.UnsafeRawScript(`
				docsearch({
					insights: true,
					appId: "9IO2WZA8L1",
					apiKey: "d8cd8b6f8f8a0c961ce971e09dbde90a",
					indexName: "htmgo",
					container: "#search-container",
					debug: false
				});
			`),
			),
		),
	)
}

func RootPage(ctx *h.RequestContext, children ...h.Ren) *h.Page {
	return ConfigurableRootPage(
		ctx,
		RootPageProps{
			Title:       "htmgo",
			Description: "build simple and scalable systems with go + htmx",
			Children:    h.Fragment(children...),
		},
	)
}

func PageWithNav(ctx *h.RequestContext, children ...h.Ren) *h.Page {
	return RootPage(
		ctx,
		h.Fragment(
			partials.NavBar(ctx, partials.NavBarProps{
				Expanded:       false,
				ShowPreRelease: true,
			}),
			h.Div(
				children...,
			),
		),
	)
}
