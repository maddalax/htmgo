package partials

import (
	"github.com/gofiber/fiber/v2"
	"mhtml/database"
	"mhtml/h"
	"mhtml/partials/sheet"
	"mhtml/ui"
	"time"
)

type Patient struct {
	Name            string
	ReasonForVisit  string
	AppointmentDate time.Time
	LocationName    string
}

func PatientList(ctx *fiber.Ctx) *h.Partial {
	patients, err := database.HList[Patient]("patients")

	if err != nil {
		return h.NewPartial(h.Div(
			h.Class("patient-list"),
			h.P("Error loading patients"),
		))
	}

	if len(patients) == 0 {
		return h.NewPartial(h.Div(
			h.Class("patient-list"),
			h.P("No patients found"),
		))
	}

	return h.NewPartial(h.Div(
		h.Id("patient-list"),
		h.List(patients, PatientRow),
	))
}

func AddPatientForm(ctx *fiber.Ctx) *h.Partial {
	return h.NewPartial(sheet.Opened(h.Div(
		h.Class("flex flex-col gap-4"),
		h.P("Add Patient", h.Class("text-lg font-bold")),
	)))
}

func PatientRow(patient *Patient) *h.Node {
	return h.Div(
		h.Class("flex flex-col gap-2"),
		h.Pf("Name: %s", patient.Name),
		h.Pf("Reason for visit: %s", patient.ReasonForVisit),
	)
}

func AddPatientButton() *h.Node {
	return ui.Button(ui.ButtonProps{
		Id:     "add-patient",
		Text:   "Add Patient",
		Class:  "bg-blue-700 text-white rounded p-2 h-12",
		Target: "#active-modal",
		Get:    h.GetPartialPath(AddPatientForm),
	})
}
