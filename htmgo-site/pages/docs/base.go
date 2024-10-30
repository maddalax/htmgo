package docs

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/pages/base"
	"htmgo-site/partials"
	"strings"
)

func Title(title string) *h.Element {
	return h.H1(
		h.Text(title),
		h.Class("text-2xl font-bold"),
	)
}

func SubTitle(title string) *h.Element {
	return h.H2(
		h.Text(title),
		h.Class("text-xl font-bold"),
	)
}

func StepTitle(title string) *h.Element {
	return h.H2(
		h.Text(title),
		h.Class("text-lg font-bold"),
	)
}

func NextStep(classes string, prev *h.Element, next *h.Element) *h.Element {
	return h.Div(
		h.Class("flex gap-2 justify-between", classes),
		prev,
		next,
	)
}

func NextBlock(text string, url string) *h.Element {
	return h.A(
		h.Href(url),
		h.Class("w-[50%] border border-slate-300 p-4 rounded text-right hover:border-blue-400 cursor-pointer"),
		h.P(
			h.Text("Next"),
			h.Class("text-slate-600 text-sm"),
		),
		h.P(
			h.Text(text),
			h.Class("text-blue-500 hover:text-blue-400"),
		),
	)
}

func PrevBlock(text string, url string) *h.Element {
	return h.A(
		h.Href(url),
		h.Class("w-[50%] border border-slate-300  p-4 rounded text-left hover:border-blue-400 cursor-pointer"),
		h.P(
			h.Text("Previous"),
			h.Class("text-slate-600 text-sm"),
		),
		h.P(
			h.Text(text),
			h.Class("text-blue-500 hover:text-blue-400"),
		),
	)
}

func Image(src string) *h.Element {
	return h.Img(
		h.Src(src),
		h.Class("rounded w-full"),
	)
}

func Text(text string) *h.Element {
	split := strings.Split(text, "\n")
	return h.Div(
		h.Class("flex flex-col gap-2 leading-relaxed text-slate-900 break-words"),
		h.List(split, func(item string, index int) *h.Element {
			return h.P(
				h.UnsafeRaw(item),
			)
		}),
	)
}

func Inline(elements ...h.Ren) *h.Element {
	return h.Div(
		h.Class("flex gap-1 items-center"),
		h.Children(elements...),
	)
}

func HelpText(text string) *h.Element {
	return h.Div(
		h.Class("text-slate-600 text-sm"),
		h.UnsafeRaw(text),
	)
}

func Link(text string, href string, additionalClasses ...string) *h.Element {
	additionalClasses = append(additionalClasses, "text-blue-500 hover:text-blue-400")
	return h.A(
		h.Href(href),
		h.Text(text),
		h.Class(
			additionalClasses...,
		),
	)
}

func DocPage(ctx *h.RequestContext, children ...h.Ren) *h.Page {

	title := "htmgo"
	for _, section := range sections {
		for _, page := range section.Pages {
			if page.Path == ctx.Request.URL.Path {
				title = fmt.Sprintf("Docs - %s", page.Title)
				break
			}
		}
	}

	return base.ConfigurableRootPage(
		ctx,
		base.RootPageProps{
			Title:       title,
			Description: "build simple and scalable systems with go + htmx",
			Children: h.Div(
				h.Class("flex h-full"),
				h.Aside(
					h.Class("hidden md:block md:min-w-60 text-white overflow-y-auto"),
					DocSidebar(),
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
							DocSidebar(),
						),
						h.Class("overflow-y-auto overflow-x-hidden pb-6 items-center w-full"),
						h.Div(
							h.Class("flex flex-col mx-auto"),
							h.Div(
								h.Class("flex flex-col justify-center items-center md:mt-6 mx-auto"),
								h.Div(
									h.Class(
										"w-full flex flex-col max-w-[90vw] md:max-w-[65vw] xl:max-w-4xl",
									),
									h.Children(children...),
								),
							),
						),
					),
				),
			),
		},
	)
}
