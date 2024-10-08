package partials

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/h"
	"sse-with-state/event"
)

type RepeaterProps struct {
	Item         func(index int) *h.Element
	RemoveButton func(index int, children ...h.Ren) *h.Element
	AddButton    *h.Element
	DefaultItems []*h.Element
	Id           string
	currentIndex int
}

func (props *RepeaterProps) itemId(index int) string {
	return fmt.Sprintf("%s-repeater-item-%d", props.Id, index)
}

func (props *RepeaterProps) addButtonId() string {
	return fmt.Sprintf("%s-repeater-add-button", props.Id)
}

func repeaterItem(ctx *h.RequestContext, item *h.Element, index int, props *RepeaterProps) *h.Element {
	id := props.itemId(index)
	return h.Div(
		h.Class("flex gap-2 items-center"),
		h.Id(id),
		item,
		props.RemoveButton(index,
			h.ClassIf(index == 0, "opacity-0 disabled"),
			h.If(index == 0, h.Disabled()),
			OnClick(ctx, func(data event.HandlerData) {
				props.currentIndex--
				event.PushElement(data,
					h.Div(
						h.Attribute("hx-swap-oob", fmt.Sprintf("delete:#%s", id)),
						repeaterItem(
							ctx, item, index, props,
						),
					),
				)
			}),
		),
	)
}

func Repeater(ctx *h.RequestContext, props RepeaterProps) *h.Element {
	if props.Id == "" {
		props.Id = h.GenId()
	}
	return h.Div(
		h.Class("flex flex-col gap-2"),
		h.List(props.DefaultItems, func(item *h.Element, index int) *h.Element {
			return repeaterItem(ctx, item, index, &props)
		}),
		h.Div(
			h.Id(props.addButtonId()),
			h.Class("flex justify-center"),
			props.AddButton,
			OnClick(ctx, func(data event.HandlerData) {
				event.PushServerSideEvent(data.SessionId, "increment")
				event.PushElement(data,
					h.Div(
						h.Attribute("hx-swap-oob", "beforebegin:#"+props.addButtonId()),
						repeaterItem(
							ctx, props.Item(props.currentIndex), props.currentIndex, &props,
						),
					),
				)
				props.currentIndex++
			}),
		),
	)
}
