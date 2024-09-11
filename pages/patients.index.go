package pages

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/h"
	"mhtml/pages/base"
	"mhtml/partials"
)

func PatientsIndex(ctx *fiber.Ctx) *h.Page {
	return h.NewPage(base.RootPage(
		h.Div(
			h.Class("flex flex-col p-4"),
			h.Div(
				h.Class("flex justify-between items-center"),
				h.P("Manage Patients", h.Class("text-lg font-bold")),
				partials.AddPatientButton()),
			h.ViewWithTriggers(partials.PatientList, "load", "every 3s"),
		),
	))
}
