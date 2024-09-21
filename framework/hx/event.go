package hx

import "fmt"

func OnEvent(event Event, modifiers ...Modifier) TriggerEvent {
	return TriggerEvent{
		event:     event,
		modifiers: modifiers,
	}
}

func OnClick(modifiers ...Modifier) TriggerEvent {
	return OnEvent(ClickEvent, modifiers...)
}

func OnLoad(modifiers ...Modifier) TriggerEvent {
	return OnEvent(LoadEvent, modifiers...)
}

func OnChange(modifiers ...Modifier) TriggerEvent {
	return OnEvent(ChangeEvent, modifiers...)
}

func OnPoll(durationSeconds int) TriggerEvent {
	return OnEvent(PollingEvent, StringModifier(fmt.Sprintf("%ds", durationSeconds)))
}
