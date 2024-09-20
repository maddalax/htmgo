package partials

import "github.com/maddalax/htmgo/framework/h"

type Link struct {
	Name string
	Path string
}

func NavBar() h.Ren {

	links := []Link{
		{"Home", "/"},
		{"News", "/news"},
		{"Patients", "/patients"},
	}

	return h.Nav(h.Class("flex gap-4 items-center p-4 text-slate-600"),
		h.Boost(),
		h.Children(
			h.Map(links, func(link Link) h.Ren {
				return h.A(h.Text(link.Name), h.Href(link.Path), h.Class("cursor-pointer hover:text-blue-400"))
			})...,
		))
}
