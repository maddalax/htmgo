package ui

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/js"
)

func CopyButton(selector string, classes ...string) *h.Element {
	classes = append(classes, "flex p-2 bg-slate-800 text-white cursor-pointer items-center")
	return h.Div(
		h.Class(classes...),
		h.Text("Copy"),
		h.OnClick(
			// language=JavaScript
			js.EvalJs(fmt.Sprintf(`
				if(!navigator.clipboard) {
					return;
				}
				let text = document.querySelector("%s").innerText;
				navigator.clipboard.writeText(text);
				self.innerText = "Copied!";
				setTimeout(() => {
					self.innerText = "Copy";
				}, 1000);
			`, selector)),
		),
	)
}

func AbsoluteCopyButton(selector string) *h.Element {
	return CopyButton(selector, "absolute top-0 right-0 rounded-bl-md")
}
