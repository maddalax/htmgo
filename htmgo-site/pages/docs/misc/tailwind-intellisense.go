package misc

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/internal/urlhelper"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

func TailwindIntellisense(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("Tailwind Intellisense"),
			Text(`
				Tailwind's language server allows you to specify custom configuration on what it should match to start giving you tailwind intellisense.
			`),
			Text(`To make this work, you will need to update the tailwind lsp config with the config below:`),
			Image("/public/tailwind-intellisense.png"),
			Text(`To make this work, you will need to update your Tailwind LSP configuration with what is below:`),
			SubTitle("Jetbrains IDE's"),
			ui.CodeSnippetFromUrl(urlhelper.ToAbsoluteUrl(ctx, "/public/jetbrains-tailwind.json"), ui.CodeSnippetProps{
				Lang:            "json",
				HideLineNumbers: true,
			}),
			Text(`
	      To find this configuration in GoLand you can go to Settings -> Languages & Frameworks -> Style Sheets -> Tailwind CSS and update the configuration there.
        These changes are additive, add these options to your existing Tailwind LSP configuration, instead of replacing the entire file.
			`),
			SubTitle("Visual Studio Code"),
			Text(`For VSCode, you should be able to update your settings.json with the following values:`),
			ui.CodeSnippetFromUrl(urlhelper.ToAbsoluteUrl(ctx, "/public/vscode-tailwind.json"), ui.CodeSnippetProps{
				Lang:            "json",
				HideLineNumbers: true,
			}),
			NextStep(
				"mt-4",
				PrevBlock("Mutation Error Extension", DocPath("/htmx-extensions/mutation-error")),
				NextBlock("Formatting blocks", DocPath("/misc/formatter")),
			),
		),
	)
}
