package partials

import (
	"mhtml/h"
	"mhtml/ui"
)

func OpenSheetButton(open bool, children ...*h.Node) *h.Node {
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
