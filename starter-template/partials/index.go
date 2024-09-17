package partials

import (
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
)

func SamplePartial(ctx echo.Context) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text(" asdasasds"))))
}

func NewPartial(ctx echo.Context) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text("This sadsl."))))
}

func NewPartial2(ctx echo.Context) *h.Partial {
	return h.NewPartial(h.Div(h.P(h.Text("This sasdsadasdwl."))))
}
