package patient

import (
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
	"starter-template/features/patient"
	"starter-template/partials/sheet"
)

func Create(ctx echo.Context) *h.Partial {
	name := ctx.FormValue("name")
	reason := ctx.FormValue("reason-for-visit")
	location := ctx.FormValue("location-name")

	err := patient.NewService(ctx).Create(patient.CreatePatientRequest{
		Name:           name,
		ReasonForVisit: reason,
		LocationName:   location,
	})

	if err != nil {
		ctx.Status(500)
		return h.NewPartialWithHeaders(h.NewHeaders(""),
			h.Div(
				h.Text("Error creating patient"),
				h.Class("text-red-500"),
			),
		)
	}

	headers := h.CombineHeaders(h.PushQsHeader(ctx, "adding", ""), &map[string]string{
		"HX-HxTrigger": "patient-added",
	})

	return h.NewPartialWithHeaders(
		headers,
		sheet.Close(ctx))
}
