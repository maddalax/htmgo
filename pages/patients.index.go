package pages

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"mhtml/pages/base"
	"mhtml/partials/patient"
	"mhtml/partials/sheet"
)

func PatientsIndex(ctx *fiber.Ctx) *h.Page {
	return h.NewPage(base.RootPage(
		h.Div(
			h.Class("flex flex-col p-4 w-full"),
			h.Div(
				h.Div(
					h.Class("flex justify-between items-center"),
					h.P("Manage Patients", h.Class("text-lg font-bold")),
					patient.AddPatientButton(),
				),
				h.PartialWithTriggers(patient.List, "load", "patient-added from:body", "every 5s"),
				h.If(
					h.GetQueryParam(ctx, "adding") == "true",
					h.View(patient.AddPatientSheetPartial, h.ReloadParams{
						Triggers: h.CreateTriggers("load"),
						Target:   sheet.Id,
					})),
			),
		),
	))
}
