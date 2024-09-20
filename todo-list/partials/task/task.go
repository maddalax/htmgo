package task

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/maddalax/htmgo/framework/h"
	"todolist/ent"
	"todolist/internal/tasks"
)

type Tab = string

const (
	TabAll      Tab = "All"
	TabActive   Tab = "Active"
	TabComplete Tab = "Complete"
)

func getActiveTab(ctx *h.RequestContext) Tab {
	if tab := h.GetQueryParam(ctx, "tab"); tab != "" {
		return tab
	}
	return TabAll
}

func Card(ctx *h.RequestContext) *h.Element {
	service := tasks.NewService(ctx.ServiceLocator())
	list, _ := service.List()

	return h.Div(
		h.Id("task-card"),
		h.Class("bg-white w-full rounded shadow-md"),
		CardBody(list, getActiveTab(ctx)),
	)
}

func CardBody(list []*ent.Task, tab Tab) *h.Element {
	return h.Div(
		h.Id("tasks-card-body"),
		Input(list),
		List(list, tab),
		Footer(list, tab),
	)
}

func Input(list []*ent.Task) *h.Element {
	return h.Div(
		h.Id("task-card-input"),
		h.Class("border border-b-slate-100 relative"),
		h.Input(
			"text",
			h.Attribute("autocomplete", "off"),
			h.Attribute("autofocus", "true"),
			h.Attribute("name", "name"),
			h.Class("pl-12 text-xl p-4 w-full outline-none focus:outline-2 focus:outline-rose-400"),
			h.Placeholder("What needs to be done?"),
			h.Post(h.GetPartialPath(Create)),
			h.Trigger("keyup[keyCode==13]"),
		),
		CompleteAllIcon(list),
	)
}

func CompleteAllIcon(list []*ent.Task) *h.Element {
	notCompletedCount := len(h.Filter(list, func(item *ent.Task) bool {
		return item.CompletedAt == nil
	}))

	return h.Div(
		h.ClassX("absolute top-0 left-0 p-4 rotate-90 text-2xl cursor-pointer", map[string]bool{
			"text-slate-400": notCompletedCount > 0,
		}), h.Text("❯"),
		h.PostPartialOnClickQs(CompleteAll, h.Ternary(notCompletedCount > 0, "complete=true", "complete=false")),
	)
}

func Footer(list []*ent.Task, activeTab Tab) *h.Element {

	notCompletedCount := len(h.Filter(list, func(item *ent.Task) bool {
		return item.CompletedAt == nil
	}))

	tabs := []Tab{TabAll, TabActive, TabComplete}

	return h.Div(
		h.Id("task-card-footer"),
		h.Class("flex items-center justify-between p-4 border-t border-b-slate-100"),
		h.Div(
			h.TextF("%d items left", notCompletedCount),
		),
		h.Div(
			h.Class("flex items-center gap-4"),
			h.List(tabs, func(tab Tab, index int) *h.Element {
				return h.P(
					h.PostOnClick(h.GetPartialPathWithQs(ChangeTab, "tab="+tab)),
					h.ClassX("cursor-pointer px-2 py-1 rounded", map[string]bool{
						"border border-rose-600": activeTab == tab,
					}),
					h.Text(tab),
				)
			}),
		),
		h.Div(
			h.PostPartialOnClick(ClearCompleted),
			h.ClassX("flex gap-2 cursor-pointer", map[string]bool{
				"opacity-0": notCompletedCount == len(list),
			}),
			h.Text("Clear completed"),
		),
	)
}

func List(list []*ent.Task, tab Tab) *h.Element {
	return h.Div(
		h.Id("task-card-list"),
		h.Class("bg-white w-full"),
		h.Div(
			h.List(list, func(item *ent.Task, index int) *h.Element {
				if tab == TabActive && item.CompletedAt != nil {
					return h.Empty()
				}
				if tab == TabComplete && item.CompletedAt == nil {
					return h.Empty()
				}
				return Task(item, false)
			}),
		),
	)
}

func Task(task *ent.Task, editing bool) *h.Element {
	return h.Div(
		h.Id(fmt.Sprintf("task-%s", task.ID.String())),
		h.ClassX("h-[80px] max-h-[80px] max-w-2xl flex items-center p-4 gap-4 cursor-pointer", h.ClassMap{
			"border border-b-slate-100": !editing,
		}),
		CompleteIcon(task),
		h.IfElse(editing,
			h.Div(
				h.Class("flex-1 h-full"),
				h.Form(
					h.Class("h-full"),
					h.Input("text",
						h.Attribute("name", "task"),
						h.Attribute("value", task.ID.String()),
						h.Class("hidden"),
					),
					h.Input(
						"text",
						h.PostPartialOnTrigger(UpdateName, h.TriggerBlur, h.TriggerKeyUpEnter),
						h.Attributes(&h.AttributeMap{
							"placeholder":  "What needs to be done?",
							"autofocus":    "true",
							"autocomplete": "off",
							"name":         "name",
							"class": h.ClassX("", h.ClassMap{
								"pl-1 h-full w-full text-xl outline-none outline-2 outline-rose-300": true,
							}),
							"value": task.Name,
						}),
					),
				),
			),
			h.P(
				h.Trigger("dblclick"),
				h.GetPartialWithQs(EditNameForm, "id="+task.ID.String()),
				h.ClassX("text-xl break-all text-wrap truncate", map[string]bool{
					"line-through text-slate-400": task.CompletedAt != nil,
				}),
				h.Text(task.Name),
			)),
	)
}

func CompleteIcon(task *ent.Task) *h.Element {
	return h.Div(
		h.Trigger("click"),
		h.Post(h.GetPartialPathWithQs(ToggleCompleted, "id="+task.ID.String())),
		h.Class("flex items-center justify-center cursor-pointer"),
		h.Div(
			h.ClassX("w-10 h-10 border rounded-full flex items-center justify-center", map[string]bool{
				"border-green-500": task.CompletedAt != nil,
				"border-slate-400": task.CompletedAt == nil,
			}),
			h.If(task.CompletedAt != nil, h.Raw(`
				<svg class="w-6 h-6 text-green-500" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
      			<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7"></path>
    			</svg>
			`)),
		),
	)
}

func UpdateName(ctx *h.RequestContext) *h.Partial {
	id, err := uuid.Parse(ctx.FormValue("task"))
	if err != nil {
		return h.NewPartial(h.Div(h.Text("invalid id")))
	}

	name := ctx.FormValue("name")
	if name == "" {
		return h.NewPartial(h.Div(h.Text("name is required")))
	}

	service := tasks.NewService(ctx.ServiceLocator())
	task, err := service.Get(id)

	if task == nil {
		return h.NewPartial(h.Div(h.Text("task not found")))
	}

	task, err = service.SetName(task.ID, name)

	if err != nil {
		return h.NewPartial(h.Div(h.Text("failed to update")))
	}

	return h.NewPartial(
		h.OobSwap(ctx, Task(task, false)))
}

func EditNameForm(ctx *h.RequestContext) *h.Partial {
	id, err := uuid.Parse(ctx.QueryParam("id"))
	if err != nil {
		return h.NewPartial(h.Div(h.Text("invalid id")))
	}

	service := tasks.NewService(ctx.ServiceLocator())
	task, err := service.Get(id)

	if task == nil {
		return h.NewPartial(h.Div(h.Text("task not found")))
	}

	return h.NewPartial(
		h.OobSwap(ctx, Task(task, true)),
	)
}

func ToggleCompleted(ctx *h.RequestContext) *h.Partial {
	id, err := uuid.Parse(ctx.QueryParam("id"))
	if err != nil {
		return h.NewPartial(h.Div(h.Text("invalid id")))
	}

	service := tasks.NewService(ctx.ServiceLocator())
	task, err := service.Get(id)

	if task == nil {
		return h.NewPartial(h.Div(h.Text("task not found")))
	}

	task, err = service.SetCompleted(task.ID, h.
		Ternary(task.CompletedAt == nil, true, false))

	if err != nil {
		return h.NewPartial(h.Div(h.Text("failed to update")))
	}

	list, _ := service.List()

	return h.NewPartial(h.Fragment(
		h.OobSwap(ctx, List(list, getActiveTab(ctx))),
		h.OobSwap(ctx, Footer(list, getActiveTab(ctx))),
		h.OobSwap(ctx, CompleteAllIcon(list)),
	))
}

func CompleteAll(ctx *h.RequestContext) *h.Partial {
	service := tasks.NewService(ctx.ServiceLocator())

	service.SetAllCompleted(ctx.QueryParam("complete") == "true")

	list, _ := service.List()

	return h.NewPartial(h.OobSwap(ctx, CardBody(list, getActiveTab(ctx))))
}

func ClearCompleted(ctx *h.RequestContext) *h.Partial {
	service := tasks.NewService(ctx.ServiceLocator())
	_ = service.ClearCompleted()

	list, _ := service.List()

	return h.NewPartial(h.OobSwap(ctx, CardBody(list, getActiveTab(ctx))))
}

func Create(ctx *h.RequestContext) *h.Partial {
	name := ctx.FormValue("name")
	if name == "" {
		return h.NewPartial(h.Div(h.Text("name is required")))
	}

	service := tasks.NewService(ctx.ServiceLocator())
	_, err := service.Create(tasks.CreateRequest{
		Name: name,
	})

	if err != nil {
		return h.NewPartial(h.Div(h.Text("failed to create")))
	}

	list, _ := service.List()

	return h.NewPartial(h.Fragment(h.OobSwap(ctx, CardBody(list, getActiveTab(ctx)))))
}

func ChangeTab(ctx *h.RequestContext) *h.Partial {
	service := tasks.NewService(ctx.ServiceLocator())
	list, _ := service.List()

	tab := ctx.QueryParam("tab")

	return h.NewPartialWithHeaders(&h.Headers{"hx-push-url": fmt.Sprintf("/tasks?tab=%s", tab)},
		h.Fragment(
			h.OobSwap(ctx, List(list, tab)),
			h.OobSwap(ctx, Footer(list, tab)),
		),
	)
}
