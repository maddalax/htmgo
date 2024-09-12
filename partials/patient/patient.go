package patient

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

func List(ctx *fiber.Ctx) *h.Partial {
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
		h.Class("mt-8"),
		h.Id("patient-list"),
		h.List(patients, Row),
	))
}

func AddPatientSheetPartial(ctx *fiber.Ctx) *h.Partial {
	return h.NewPartialWithHeaders(
		h.PushQsHeader(ctx, "adding", "true"),
		AddPatientSheet(h.CurrentPath(ctx)),
	)
}

func AddPatientSheet(onClosePath string) h.Renderable {
	return sheet.Opened(
		sheet.Props{
			OnClosePath: onClosePath,
			ClassName:   "w-[400px] bg-gray-100 p-4",
			Root: h.Div(
				h.Class("flex flex-col gap-4"),
				h.P("Add Patient", h.Class("text-lg font-bold")),
				addPatientForm(),
			),
		})
}

func addPatientForm() h.Renderable {
	return h.Form(
		h.TriggerChildren(),
		h.Post(h.GetPartialPath(Create)),
		h.Class("flex flex-col gap-2"),
		ui.Input(ui.InputProps{
			Type:         "text",
			Label:        "Name",
			Name:         "name",
			DefaultValue: "fart",
		}),
		ui.Input(ui.InputProps{
			Type:  "text",
			Label: "Reason for visit",
			Name:  "reason-for-visit",
		}),
		ui.Input(ui.InputProps{
			Type:  "date",
			Label: "Appointment Date",
			Name:  "appointment-date",
		}),
		ui.Input(ui.InputProps{
			Type:  "text",
			Label: "Location Name",
			Name:  "location-name",
		}),
		ui.PrimaryButton(ui.ButtonProps{
			Text:  "Add Patient",
			Class: "rounded p-2",
			Type:  "submit",
		}),
	)
}

func Row(patient *Patient, index int) h.Renderable {
	return h.Div(
		h.Class("flex flex-col gap-2 rounded p-4", h.Ternary(index%2 == 0, "bg-red-100", "")),
		h.Pf("Name: %s", patient.Name),
		h.Pf("Reason for visit: %s", patient.ReasonForVisit),
	)
}

func AddPatientButton() h.Renderable {
	return ui.Button(ui.ButtonProps{
		Id:     "add-patient",
		Text:   "Add Patient",
		Class:  "bg-blue-700 text-white rounded p-2 h-12",
		Target: sheet.Id,
		Get:    h.GetPartialPath(AddPatientSheetPartial),
	})
}
