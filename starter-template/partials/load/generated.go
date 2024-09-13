// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package load

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/mhtml/framework/h"
	"github.com/maddalax/mhtml/starter-template/partials"
	"github.com/maddalax/mhtml/starter-template/partials/patient"
	"github.com/maddalax/mhtml/starter-template/partials/sheet"
)

func GetPartialFromContext(ctx *fiber.Ctx) *h.Partial {
	path := ctx.Path()
	if path == "NewsSheet" || path == "/mhtml/partials.NewsSheet" {
		return partials.NewsSheet(ctx)
	}
	if path == "NewsSheetOpenCount" || path == "/mhtml/partials.NewsSheetOpenCount" {
		return partials.NewsSheetOpenCount(ctx)
	}
	if path == "Create" || path == "/mhtml/partials/patient.Create" {
		return patient.Create(ctx)
	}
	if path == "List" || path == "/mhtml/partials/patient.List" {
		return patient.List(ctx)
	}
	if path == "AddPatientSheetPartial" || path == "/mhtml/partials/patient.AddPatientSheetPartial" {
		return patient.AddPatientSheetPartial(ctx)
	}
	if path == "ValidateForm" || path == "/mhtml/partials/patient.ValidateForm" {
		return patient.ValidateForm(ctx)
	}
	if path == "Close" || path == "/mhtml/partials/sheet.Close" {
		return sheet.Close(ctx)
	}
	return nil
}
