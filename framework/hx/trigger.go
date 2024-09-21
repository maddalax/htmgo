package hx

import (
	"strings"
)

type Trigger struct {
	events []TriggerEvent
}

type TriggerEvent struct {
	event     Event
	modifiers []Modifier
}

func ToHtmxTriggerName(event string) string {
	if strings.HasPrefix(event, "htmx:") {
		return event[5:]
	}
	if strings.HasPrefix(event, "on") {
		return event[2:]
	}
	return event
}

func FormatEventName(event string, isDomEvent bool) string {
	raw := ToHtmxTriggerName(event)
	if isDomEvent {
		return "on" + raw
	}
	return event
}

func NewTrigger(opts ...TriggerEvent) *Trigger {
	t := Trigger{
		events: make([]TriggerEvent, 0),
	}
	if len(opts) > 0 {
		for i, opt := range opts {
			opts[i].event = ToHtmxTriggerName(opt.event)
		}
		t.events = opts
	}
	return &t
}

func NewStringTrigger(trigger string) Trigger {
	t := Trigger{
		events: make([]TriggerEvent, 0),
	}

	split := strings.Split(trigger, ", ")
	for _, s := range split {
		parts := strings.Split(s, " ")
		event := ToHtmxTriggerName(parts[0])

		modifiers := make([]Modifier, 0)
		if len(parts) > 1 {
			for _, m := range parts[1:] {
				modifiers = append(modifiers, RawModifier{modifier: m})
			}
		}
		t.events = append(t.events, TriggerEvent{
			event:     event,
			modifiers: modifiers,
		})
	}
	return t
}

func (t Trigger) AddEvent(event TriggerEvent) Trigger {
	t.events = append(t.events, event)
	return t
}

func (t Trigger) ToString() string {
	builder := strings.Builder{}
	for i, e := range t.events {
		eventName := ToHtmxTriggerName(e.event)
		builder.WriteString(eventName)
		for _, m := range e.modifiers {
			builder.WriteString(" ")
			builder.WriteString(m.Modifier())
		}
		if i < len(t.events)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}

func (t Trigger) Render(builder *strings.Builder) {
	builder.WriteString(t.ToString())
}
