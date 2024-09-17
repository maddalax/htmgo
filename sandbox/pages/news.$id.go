package pages

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maddalax/htmgo/framework/h"
)

func Test(ctx echo.Context) *h.Page {
	text := fmt.Sprintf("News ID: %s", ctx.Params("id"))
	return h.NewPage(
		h.Div(h.Text(text)),
	)
}
