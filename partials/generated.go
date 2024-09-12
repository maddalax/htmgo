// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package partials

import "mhtml/h"
import "github.com/gofiber/fiber/v2"
import "mhtml/partials/patient"
import "mhtml/partials/sheet"

func GetPartialFromContext(ctx *fiber.Ctx) *h.Partial {
	path := ctx.Path()
	if path == "NewsSheet" || path == "/mhtml/partials.NewsSheet" {
		return NewsSheet(ctx)
	}
	if path == "NewsSheetOpenCount" || path == "/mhtml/partials.NewsSheetOpenCount" {
		return NewsSheetOpenCount(ctx)
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
	if path == "Close" || path == "/mhtml/partials/sheet.Close" {
		return sheet.Close(ctx)
	}
	return nil
}
