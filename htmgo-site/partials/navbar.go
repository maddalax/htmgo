package partials

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/internal/httpjson"
	"time"
)

type NavItem struct {
	Name string
	Url  string
}

var navItems = []NavItem{
	{Name: "Docs", Url: "/docs"},
	{Name: "Examples", Url: "/examples"},
	{Name: "Convert HTML", Url: "/html-to-go"},
}

type NavBarProps struct {
	Expanded       bool
	ShowPreRelease bool
}

func ToggleNavbar(ctx *h.RequestContext) *h.Partial {
	return h.SwapManyPartial(
		ctx,
		MobileNav(ctx, h.GetQueryParam(ctx, "expanded") == "true"),
	)
}

var CachedStar = h.CachedT(time.Minute*15, func(t *h.RequestContext) *h.Element {
	return Star(t)
})

func Star(ctx *h.RequestContext) *h.Element {

	type Repo struct {
		StarCount int `json:"stargazers_count"`
	}

	fmt.Printf("making github star request\n")
	count := 0
	response, err := httpjson.Get[Repo]("https://api.github.com/repos/maddalax/htmgo")

	if err == nil && response != nil {
		count = response.StarCount
	}

	return h.A(
		h.Href("https://github.com/maddalax/htmgo"),
		h.Target("_blank"),
		h.Class("inline-flex items-center rounded overflow-hidden shadow-sm"),
		h.Div(
			h.Class("flex items-center px-2 py-1 bg-gray-800 text-white text-sm font-semibold hover:bg-gray-700 transition"),
			h.Svg(
				h.Class("w-4 h-4 -mt-0.5 mr-0.5 stroke-current text-white"),
				h.Attribute("xmlns", "http://www.w3.org/2000/svg"),
				h.Attribute("viewBox", "0 0 24 24"),
				h.Attribute("fill", "none"),           // No fill
				h.Attribute("stroke", "currentColor"), // Apply stroke
				h.Attribute("stroke-width", "2"),      // Stroke width
				h.Path(
					h.D("M12 17.27l5.18 3.05-1.64-5.68 4.46-3.87-5.88-.5L12 3.5l-2.12 6.77-5.88.5 4.46 3.87-1.64 5.68L12 17.27z"),
				),
			),
			h.Text("Star"),
		),
		h.If(count > 0, h.Div(
			h.Class("flex items-center px-3 py-1 bg-black text-white text-sm font-semibold"),
			h.Pf("%d", count),
		)),
	)
}

func NavBar(ctx *h.RequestContext, props NavBarProps) *h.Element {
	//prelease := h.If(props.ShowPreRelease, h.A(
	//	h.Class("bg-blue-200 text-blue-700 text-center p-2 flex items-center justify-center"),
	//	h.Href("https://github.com/maddalax/htmgo/issues"),
	//	h.Attribute("target", "_blank"),
	//	h.Text("htmgo."),
	//))

	desktopNav := h.Nav(
		h.Class("hidden sm:block bg-neutral-100 border border-b-slate-300 p-4 md:p-3 max-h-[100vh - 9rem] overflow-y-auto"),
		h.Div(
			h.Class("max-w-[95%] md:max-w-3xl px-4 mx-auto"),
			h.Div(
				h.Class("flex justify-between items-center"),
				h.Div(
					h.Class("flex items-center"),
					h.A(
						h.Class("text-2xl"),
						h.Href("/"),
						h.Text("htmgo"),
					)),
				h.Div(
					h.Id("search-container"),
				),
				h.Div(
					h.Class("flex gap-4 items-center"),
					h.List(navItems, func(item NavItem, index int) *h.Element {
						return h.Div(
							h.Class("flex items-center"),
							h.A(
								h.Class(""),
								h.Href(item.Url),
								h.Text(item.Name),
							),
						)
					}),
					CachedStar(ctx),
				),
			),
		),
	)

	return h.Div(
		h.Id("navbar"),
		//prelease,
		MobileNav(ctx, props.Expanded),
		desktopNav,
	)
}

func MobileNav(ctx *h.RequestContext, expanded bool) *h.Element {
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
					h.Class("flex items-center gap-3"),
					h.Div(h.Class("mt-1"), CachedStar(ctx)),
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
						h.UnsafeRaw("&#9776;"),
					),
				),
			),
		),
		h.If(expanded, h.Div(
			h.Class("mt-2 ml-2 flex flex-col gap-2"),
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
		)),
	)
}
