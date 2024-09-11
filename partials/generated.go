// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package partials

import "mhtml/h"
import "github.com/gofiber/fiber/v2"
import "mhtml/partials/sheet"

func GetPartialFromContext(ctx *fiber.Ctx) *h.Partial {
	path := ctx.Path()
	if path == "NewsSheet" || path == "/mhtml/partials.NewsSheet" {
		return NewsSheet(ctx)
	}
	if path == "NewsSheetOpenCount" || path == "/mhtml/partials.NewsSheetOpenCount" {
		return NewsSheetOpenCount(ctx)
	}
	if path == "PatientList" || path == "/mhtml/partials.PatientList" {
		return PatientList(ctx)
	}
	if path == "AddPatientSheet" || path == "/mhtml/partials.AddPatientSheet" {
		return AddPatientSheet(ctx)
	}
	if path == "Close" || path == "/mhtml/partials/sheet.Close" {
		return sheet.Close(ctx)
	}
	return nil
}
