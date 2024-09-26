// Package __htmgo THIS FILE IS GENERATED. DO NOT EDIT.
package __htmgo

import "github.com/maddalax/htmgo/framework/h"
import "github.com/labstack/echo/v4"
import "todolist/partials/task"

func GetPartialFromContext(ctx echo.Context) *h.Partial {
	path := ctx.Request().URL.Path
	if path == "UpdateName" || path == "/todolist/partials/task.UpdateName" {
		cc := ctx.(*h.RequestContext)
		return task.UpdateName(cc)
	}
	if path == "EditNameForm" || path == "/todolist/partials/task.EditNameForm" {
		cc := ctx.(*h.RequestContext)
		return task.EditNameForm(cc)
	}
	if path == "ToggleCompleted" || path == "/todolist/partials/task.ToggleCompleted" {
		cc := ctx.(*h.RequestContext)
		return task.ToggleCompleted(cc)
	}
	if path == "CompleteAll" || path == "/todolist/partials/task.CompleteAll" {
		cc := ctx.(*h.RequestContext)
		return task.CompleteAll(cc)
	}
	if path == "ClearCompleted" || path == "/todolist/partials/task.ClearCompleted" {
		cc := ctx.(*h.RequestContext)
		return task.ClearCompleted(cc)
	}
	if path == "Create" || path == "/todolist/partials/task.Create" {
		cc := ctx.(*h.RequestContext)
		return task.Create(cc)
	}
	if path == "ChangeTab" || path == "/todolist/partials/task.ChangeTab" {
		cc := ctx.(*h.RequestContext)
		return task.ChangeTab(cc)
	}
	return nil
}

func RegisterPartials(f *echo.Echo) {
	f.Any("todolist/partials*", func(ctx echo.Context) error {
		partial := GetPartialFromContext(ctx)
		if partial == nil {
			return ctx.NoContent(404)
		}
		return h.PartialView(ctx, partial)
	})
}
