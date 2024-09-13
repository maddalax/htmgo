package partials

import (
	"github.com/maddalax/mhtml/framework-ui/ui"
	"github.com/maddalax/mhtml/framework/h"
)

func OpenSheetButton(open bool, children ...h.Renderable) h.Renderable {
	if open {
		return ui.PrimaryButton(ui.ButtonProps{
			Id:       "open-sheet",
			Text:     "Close NewsSheet",
			Target:   "#sheet-partial",
			Get:      h.GetPartialPathWithQs(NewsSheet, "open=false"),
			Children: children,
		})
	} else {
		return ui.PrimaryButton(ui.ButtonProps{
			Id:       "open-sheet",
			Text:     "Open NewsSheet",
			Target:   "#sheet-partial",
			Get:      h.GetPartialPathWithQs(NewsSheet, "open=true"),
			Children: children,
		})
	}
}
