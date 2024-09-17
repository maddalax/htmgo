// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package load

import "github.com/maddalax/htmgo/framework/h"
import "github.com/labstack/echo/v4"
import "starter-template/partials"
import "starter-template/partials/patient"
import "starter-template/partials/sheet"

func GetPartialFromContext(ctx echo.Context) *h.Partial {
	path := ctx.Path()
	if path == "NewsSheet" || path == "/starter-template/partials.NewsSheet" {
		return partials.NewsSheet(ctx)
	}
	if path == "NewsSheetOpenCount" || path == "/starter-template/partials.NewsSheetOpenCount" {
		return partials.NewsSheetOpenCount(ctx)
	}
	if path == "Create" || path == "/starter-template/partials/patient.Create" {
		return patient.Create(ctx)
	}
	if path == "List" || path == "/starter-template/partials/patient.List" {
		return patient.List(ctx)
	}
	if path == "AddPatientSheetPartial" || path == "/starter-template/partials/patient.AddPatientSheetPartial" {
		return patient.AddPatientSheetPartial(ctx)
	}
	if path == "ValidateForm" || path == "/starter-template/partials/patient.ValidateForm" {
		return patient.ValidateForm(ctx)
	}
	if path == "Close" || path == "/starter-template/partials/sheet.Close" {
		return sheet.Close(ctx)
	}
	return nil
}

func RegisterPartials(f *echo.Echo) {
	f.All("starter-template/partials*", func(ctx echo.Context) error {
		partial := GetPartialFromContext(ctx)
		if partial == nil {
			return ctx.SendStatus(404)
		}
		return h.PartialView(ctx, partial)
	})
}
