package control

import (
	"github.com/maddalax/htmgo/framework/h"
	. "htmgo-site/pages/docs"
	"htmgo-site/ui"
)

func IfElse(ctx *h.RequestContext) *h.Page {
	return DocPage(
		ctx,
		h.Div(
			h.Class("flex flex-col gap-3"),
			Title("If / Else Statements"),
			Text(`
				If / else statements are useful when you want to conditionally render attributes or elements / components.
				htmgo provides a couple of utilities to do so:
			`),
			Text("Example: Rendering an icon if the task is complete"),
			ui.GoCodeSnippet(IfElseExample),
			Text("Example: Using ternary operator to call different partials based on a condition"),
			ui.GoCodeSnippet(TenaryExample),
			Text(`Example: Rendering multiple classes based on a condition`),
			ui.GoCodeSnippet(ConditionalClassExample),
			Text("Example: Rendering a single class based on a condition"),
			ui.GoCodeSnippet(ClassIfElseExample),
			Text("Example: Rendering different elements based on a condition"),
			ui.GoCodeSnippetSingleLine(IfElseExample2),
			Text(`
				<b>Note:</b> This will execute both <b>EditTaskForm</b> and <b>ViewTask</b>, no matter if the condition is true or false, since a function is being called here.
				If you do not want to call the function at all unless the condition is true, use <b>h.IfElseLazy</b>
			`),
			ui.GoCodeSnippetSingleLine(IfElseExample3),
			NextStep(
				"mt-4",
				PrevBlock("Raw HTML", DocPath("/core-concepts/raw-html")),
				NextBlock("Rendering Lists", DocPath("/control/loops")),
			),
		),
	)
}

const IfElseExample = `
h.Div(
	h.If(
		task.CompletedAt != nil,
		CompleteIcon()
	)
)
`

const TenaryExample = `h.Div(
	h.PostPartialWithQs(
		h.Ternary(!editing, StartEditing, SaveEditing),
		h.NewQs("id", record.Id),
	),
)
`

const ConditionalClassExample = `h.ClassX("w-10 h-10 border rounded-full", map[string]bool {
				"border-green-500": task.CompletedAt != nil,
				"border-slate-400": task.CompletedAt == nil,
})`

const IfElseExample2 = `h.IfElse(editing, EditTaskForm(), ViewTask())`

const IfElseExample3 = `h.IfElseLazy(editing, EditTaskForm, ViewTask)`

const ClassIfElseExample = `
h.Div(
	h.ClassIf(task.CompletedAt != nil, "border-green-500"),
)
`
