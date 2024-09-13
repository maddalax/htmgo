package patient

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/mhtml/framework/h"
	"github.com/maddalax/mhtml/starter-template/features/patient"
	"github.com/maddalax/mhtml/starter-template/partials/sheet"
	"strings"
)

func List(ctx *fiber.Ctx) *h.Partial {
	patients, err := patient.NewService(ctx).List()

	if err != nil {
		return h.NewPartial(h.Div(
			h.Class("patient-list"),
			h.Pf("Error loading patients"),
		))
	}

	if len(patients) == 0 {
		return h.NewPartial(h.Div(
			h.Class("patient-list"),
			h.Pf("No patients found"),
		))
	}

	return h.NewPartial(h.Div(
		h.Class("mt-8"),
		h.Id("patient-list"),
		h.List(patients, Row),
	))
}

func AddPatientSheetPartial(ctx *fiber.Ctx) *h.Partial {
	closePathQs := h.GetQueryParam(ctx, "onClosePath")
	return h.NewPartialWithHeaders(
		h.PushQsHeader(ctx, "adding", "true"),
		AddPatientSheet(
			h.Ternary(closePathQs != "", closePathQs, h.CurrentPath(ctx)),
		),
	)
}

func AddPatientSheet(onClosePath string) h.Renderable {
	return sheet.Opened(
		sheet.Props{
			OnClosePath: onClosePath,
			ClassName:   "w-[400px] bg-gray-100 p-4",
			Root: h.Div(
				h.Class("flex flex-col gap-4"),
				h.P(h.Text("Add Patient"), h.Class("text-lg font-bold")),
				addPatientForm(),
			),
		})
}

func ValidateForm(ctx *fiber.Ctx) *h.Partial {
	trigger := h.GetTriggerName(ctx)
	value := ctx.FormValue(trigger)

	if trigger == "name" {
		if strings.ToLower(value) == "sydne" {
			return h.NewPartial(h.Pf("that name is reserved"))
		}
	}

	if trigger == "reason-for-visit" {
		if strings.ToLower(value) == "arm hurts" {
			return h.NewPartial(h.Pf("lol that reason is fake"))
		}
	}

	if trigger == "location-name" {
		if strings.ToLower(value) == "hospital" {
			return h.NewPartial(h.Pf("that location is reserved"))
		}
	}

	return h.NewPartial(h.Fragment())
}

func addPatientForm() h.Renderable {
	return h.Form(
		h.HxExtension("debug, trigger-children"),
		h.Attribute("hx-target-5*", "#submit-error"),
		h.Post(h.GetPartialPath(Create)),
		h.Class("flex flex-col gap-2"),
		ui.Input(ui.InputProps{
			Type:           "text",
			Label:          "Name",
			Name:           "name",
			DefaultValue:   "",
			ValidationPath: h.GetPartialPath(ValidateForm),
		}),
		ui.Input(ui.InputProps{
			Type:           "text",
			Label:          "Reason for visit",
			Name:           "reason-for-visit",
			ValidationPath: h.GetPartialPath(ValidateForm),
		}),
		ui.Input(ui.InputProps{
			Type:  "date",
			Label: "Appointment Date",
			Name:  "appointment-date",
		}),
		ui.Input(ui.InputProps{
			Type:           "text",
			Label:          "Location Name",
			Name:           "location-name",
			ValidationPath: h.GetPartialPath(ValidateForm),
		}),
		ui.PrimaryButton(ui.ButtonProps{
			Text:  "Add Patient",
			Class: "rounded p-2",
			Type:  "submit",
		}),
		h.Div(
			h.Id("submit-error"),
			h.Class("text-red-500"),
		),
	)
}

func Row(patient *patient.Patient, index int) h.Renderable {
	return h.Div(
		h.Class("flex flex-col gap-2 rounded p-4", h.Ternary(index%2 == 0, "bg-red-100", "")),
		h.Pf("Name: %s", patient.Name),
		h.Pf("Reason for visit: %s", patient.ReasonForVisit),
	)
}

func AddPatientButton() h.Renderable {
	return ui.Button(ui.ButtonProps{
		Id:      "add-patient",
		Text:    "Add Patient",
		Class:   "bg-blue-700 text-white rounded p-2 h-12",
		Trigger: "qs:adding, click",
		Target:  sheet.Id,
		Get:     h.GetPartialPathWithQs(AddPatientSheetPartial, "onClosePath=/patients"),
	})
}
