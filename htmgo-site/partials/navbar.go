package partials

import (
	"github.com/maddalax/htmgo/framework/h"
)

type NavItem struct {
	Name string
	Url  string
}

func ToggleNavbar(ctx *h.RequestContext) *h.Partial {
	return h.SwapManyPartial(
		ctx,
		MobileNav(h.GetQueryParam(ctx, "expanded") == "true"),
	)
}

var navItems = []NavItem{
	{Name: "Docs", Url: "/docs"},
	{Name: "Examples", Url: "/examples"},
}

func Star() *h.RawContent {
	return h.Raw(`
		<a 
		class="github-button" 
		href="https://github.com/maddalax/htmgo" 
		data-color-scheme="no-preference: light; light: light; dark: dark;" 
		data-icon="octicon-star" 
		data-size="large" 
		data-show-count="true"
		aria-label="Star maddalax/htmgo on GitHub">Star</a>
	`)
}

func NavBar(expanded bool) *h.Element {
	prelease := h.Div(h.Class("bg-yellow-200 text-yellow-800 text-center p-2"),
		h.Text("This is a prerelease version and generally should not be used at this time. Watch on github for the stable release!"),
	)

	desktopNav := h.Nav(
		h.Script("https://buttons.github.io/buttons.js"),
		h.Class("hidden sm:block bg-neutral-100 border border-b-slate-300 p-4 md:p-3"),
		h.Div(
			h.Class("max-w-[95%] md:max-w-prose mx-auto"),
			h.Div(
				h.Class("flex justify-between items-center"),
				h.Div(
					h.Class("flex items-center"),
					h.A(
						h.Boost(),
						h.Class("text-2xl"),
						h.Href("/"),
						h.Text("htmgo"),
					)),
				h.Div(
					h.Class("flex gap-4 items-center"),
					h.List(navItems, func(item NavItem, index int) *h.Element {
						return h.Div(
							h.Class("flex items-center"),
							h.A(
								h.Boost(),
								h.Class(""),
								h.Href(item.Url),
								h.Text(item.Name),
							),
						)
					}),
					h.Div(h.Class("ml-2 hidden md:block"), Star()),
				),
			),
		),
	)

	return h.Div(
		h.Class("mb-4"),
		h.Id("navbar"),
		prelease,
		MobileNav(expanded),
		desktopNav,
	)
}

func MobileNav(expanded bool) *h.Element {
	return h.Nav(
		h.Id("mobile-nav"),
		h.Class("block sm:hidden bg-neutral-100 border border-b-slate-300 p-4 md:p-3"),
		h.Div(
			h.Class("max-w-[95%] md:max-w-prose mx-auto"),
			h.Div(
				h.Class("flex justify-between items-center"),
				h.Div(
					h.Class("flex items-center"),
					h.A(
						h.Boost(),
						h.Class("text-2xl"),
						h.Href("/"),
						h.Text("htmgo"),
					)),
				h.Div(
					h.Class("flex items-center"),
					h.Button(
						h.Boost(),

						h.GetPartialWithQs(
							ToggleNavbar,
							h.NewQs("expanded", h.Ternary(expanded, "false", "true"), "test", "true"),
							"click",
						),

						h.AttributePairs(
							"class", "text-2xl",
							"aria-expanded", h.Ternary(expanded, "true", "false"),
						),

						h.Class("text-2xl"),
						h.Text("☰"),
					),
				),
			),
		),
		h.If(expanded, h.Div(
			h.Class("mt-2 ml-2 flex flex-col gap-2"),
			h.Script("https://buttons.github.io/buttons.js"),
			h.List(navItems, func(item NavItem, index int) *h.Element {
				return h.Div(
					h.Class("flex items-center"),
					h.A(
						h.Boost(),
						h.Class(""),
						h.Href(item.Url),
						h.Text(item.Name),
					),
				)
			}),
			h.Div(Star()),
		)),
	)
}
