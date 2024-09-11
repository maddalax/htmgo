package pages

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"mhtml/pages/base"
	"mhtml/partials/patient"
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
				h.ViewWithTriggers(patient.List, "load", "patient-added from:body"),
			),
		),
	))
}
