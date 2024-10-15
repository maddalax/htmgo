package components

import "github.com/maddalax/htmgo/framework/h"

func Badge(text string, active bool, children ...h.Ren) *h.Element {
	return h.Button(
		h.Text(text),
		h.ClassX("font-semibold px-3 py-1 rounded-full cursor-pointer h-[32px]", h.ClassMap{
			"bg-rose-500 text-white": active,
			"bg-neutral-300":         !active,
		}),
		h.Children(children...),
	)
}
