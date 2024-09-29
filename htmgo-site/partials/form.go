package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"time"
)

func SubmitForm(ctx *h.RequestContext) *h.Partial {
	time.Sleep(time.Second * 3)
	return h.NewPartial(
		h.Div(h.Text("Form submitted")),
	)
}
