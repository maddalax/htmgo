package patient

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"mhtml/database"
	"mhtml/h"
	"mhtml/partials/sheet"
	"time"
)

func Create(ctx *fiber.Ctx) *h.Partial {
	name := ctx.FormValue("name")
	reason := ctx.FormValue("reason-for-visit")
	location := ctx.FormValue("location-name")

	database.HSet("patients", uuid.New().String(), Patient{
		Name:            name,
		ReasonForVisit:  reason,
		AppointmentDate: time.Now(),
		LocationName:    location,
	})

	headers := h.CombineHeaders(h.PushQsHeader(ctx, "adding", ""), &map[string]string{
		"HX-Trigger": "patient-added",
	})

	return h.NewPartialWithHeaders(
		headers,
		sheet.Close(ctx))
}
