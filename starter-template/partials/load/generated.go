// Package partials THIS FILE IS GENERATED. DO NOT EDIT.
package load

import "github.com/maddalax/mhtml/framework/h"
import "github.com/gofiber/fiber/v2"
import "github.com/maddalax/mhtml/starter-template/partials"
import "github.com/maddalax/mhtml/starter-template/partials/patient"
import "github.com/maddalax/mhtml/starter-template/partials/sheet"

func GetPartialFromContext(ctx *fiber.Ctx) *h.Partial {
	path := ctx.Path()
	if path == "NewsSheet" || path == "/github.com/maddalax/mhtml/starter-template/partials.NewsSheet" {
		return partials.NewsSheet(ctx)
	}
	if path == "NewsSheetOpenCount" || path == "/github.com/maddalax/mhtml/starter-template/partials.NewsSheetOpenCount" {
		return partials.NewsSheetOpenCount(ctx)
	}
	if path == "Create" || path == "/github.com/maddalax/mhtml/starter-template/partials/patient.Create" {
		return patient.Create(ctx)
	}
	if path == "List" || path == "/github.com/maddalax/mhtml/starter-template/partials/patient.List" {
		return patient.List(ctx)
	}
	if path == "AddPatientSheetPartial" || path == "/github.com/maddalax/mhtml/starter-template/partials/patient.AddPatientSheetPartial" {
		return patient.AddPatientSheetPartial(ctx)
	}
	if path == "ValidateForm" || path == "/github.com/maddalax/mhtml/starter-template/partials/patient.ValidateForm" {
		return patient.ValidateForm(ctx)
	}
	if path == "Close" || path == "/github.com/maddalax/mhtml/starter-template/partials/sheet.Close" {
		return sheet.Close(ctx)
	}
	return nil
}

func RegisterPartials(f *fiber.App) {
	f.All("github.com/maddalax/mhtml/starter-template/partials*", func(ctx *fiber.Ctx) error {
		partial := GetPartialFromContext(ctx)
		if partial == nil {
			return ctx.SendStatus(404)
		}
		return h.PartialView(ctx, partial)
	})
}
