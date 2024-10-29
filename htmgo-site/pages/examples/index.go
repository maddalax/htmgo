package examples

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/pages/base"
	"htmgo-site/partials"
)

func Index(ctx *h.RequestContext) *h.Page {
	snippet := GetSnippet(ctx)
	return base.RootPage(
		ctx,
		h.Div(
			h.Class("flex h-full"),
			h.Aside(
				h.Class("hidden md:block md:min-w-60 text-white overflow-y-auto"),
				SnippetSidebar(),
			),
			h.Div(
				h.Class("flex flex-col flex-1 overflow-hidden"),
				partials.NavBar(ctx, partials.NavBarProps{
					Expanded:       false,
					ShowPreRelease: false,
				}),
				h.Main(
					h.Div(
						h.Class("w-full md:hidden bg-neutral-50 overflow-y-auto mb-4 border-b border-b-slate-300"),
						SnippetSidebar(),
					),
					h.Class("overflow-y-auto justify-center overflow-x-hidden pb-6 items-center w-full"),
					h.Div(
						h.Class("flex flex-col mx-auto"),
						h.Div(
							h.Class("flex flex-col justify-center items-center md:mt-6 mx-auto"),
						),
						h.IfElseLazy(
							snippet != nil,
							func() *h.Element {
								return snippetView(ctx, snippet)
							},
							emptyState,
						),
					),
				),
			),
		),
	)
}

func viewSourceButton(snippet *Snippet) *h.Element {
	return h.Div(
		h.Class("flex gap-2 items-center"),
		h.A(
			h.Fragment(
				githubLogo(),
				h.If(
					snippet.externalRoute != "",
					h.Text("View source"),
				),
			),
			h.Href(
				h.Ternary(snippet.sourceCodePath == "", GetGithubPath(snippet.path), snippet.sourceCodePath),
			),
			h.Class("flex gap-2 items-center font-sm text-blue-500 hover:text-blue-400"),
		),
		h.If(
			snippet.externalRoute == "",
			h.H3(
				h.Text("Source Code"),
				h.Class("text-lg font-bold"),
			),
		),
	)
}

func snippetView(ctx *h.RequestContext, snippet *Snippet) *h.Element {
	return h.Div(
		h.Class("flex flex-col mx-auto items-center gap-6 max-w-[90vw] md:max-w-[75vw] xl:max-w-4xl px-8"),
		h.Div(
			h.Class("flex flex-col gap-1 w-full"),
			h.H2(
				h.Text(snippet.name),
				h.Class("text-2xl font-bold"),
			),
			h.If(
				snippet.description != "",
				h.P(
					h.Text(snippet.description),
					h.Class("text-slate-900"),
				),
			),
			h.If(
				snippet.externalRoute != "",
				h.Div(
					h.Class("mt-3"),
					viewSourceButton(snippet),
				),
			),
		),
		h.Div(
			h.ClassX("", map[string]bool{
				"mb-6 border px-8 py-4 rounded-md shadow-sm border-slate-200 w-full": snippet.externalRoute == "",
			}),
			h.IfElse(
				snippet.externalRoute != "",
				h.Div(
					h.Class("relative"),
					h.IFrame(
						snippet.externalRoute,
						h.Class("h-full min-h-[800px] w-[50vw] rounded"),
					),
					h.A(
						h.Class("w-[50vw] rounded absolute top-0 left-0 h-full bg-gray-800 bg-opacity-50 backdrop-blur-[2px] flex items-center justify-center cursor-pointer"),
						h.Href(
							snippet.externalRoute,
						),
						h.Target("_blank"),
						h.Span(
							h.Text("Click to view"),
							h.Class("text-white text-lg font-bold"),
						),
					),
				),
				h.Div(
					h.IfElseLazy(snippet.partial != nil, func() *h.Element {
						return snippet.partial(ctx).Root
					}, h.Empty),
				),
			),
		),
		h.If(
			snippet.externalRoute == "",
			h.Div(
				h.Class("flex flex-col gap-2 justify-center"),
				viewSourceButton(snippet),
				RenderCodeToStringCached(snippet),
			),
		),
	)
}

func emptyState() *h.Element {
	return h.Div(
		h.Class("flex flex-col gap-2 justify-center items-center mt-8"),
		h.Div(
			h.Class("flex gap-2 items-center"),
			h.H3(
				h.Text("Choose an example on the sidebar to view"),
				h.Class("text-lg"),
			),
		),
	)
}

func githubLogo() *h.Element {
	return h.Body(
		h.Svg(
			h.Attribute("xmlns", "http://www.w3.org/2000/svg"),
			h.Width(24),
			h.Height(24),
			h.Attribute("viewBox", "0 0 24 24"),
			h.Path(
				h.Attribute("d", "M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"),
			),
		),
	)
}
