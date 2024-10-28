package ui

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"github.com/maddalax/htmgo/framework/js"
)

func CopyButton(selector string) *h.Element {
	return h.Div(
		h.Class("absolute top-0 right-0 p-2 bg-slate-800 text-white rounded-bl-md cursor-pointer"),
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
