package h

import (
	"fmt"
)

type LifeCycle struct {
	handlers map[HxEvent][]JsCommand
}

func NewLifeCycle() *LifeCycle {
	return &LifeCycle{
		handlers: make(map[HxEvent][]JsCommand),
	}
}

func (l *LifeCycle) OnEvent(event HxEvent, cmd ...JsCommand) *LifeCycle {
	if l.handlers[event] == nil {
		l.handlers[event] = []JsCommand{}
	}
	l.handlers[event] = append(l.handlers[event], cmd...)
	return l
}

func (l *LifeCycle) BeforeRequest(cmd ...JsCommand) *LifeCycle {
	l.OnEvent(HxBeforeRequest, cmd...)
	return l
}

func OnEvent(event HxEvent, cmd ...JsCommand) *LifeCycle {
	return NewLifeCycle().OnEvent(event, cmd...)
}

func BeforeRequest(cmd ...JsCommand) *LifeCycle {
	return NewLifeCycle().BeforeRequest(cmd...)
}

func AfterRequest(cmd ...JsCommand) *LifeCycle {
	return NewLifeCycle().AfterRequest(cmd...)
}

func OnMutationError(cmd ...JsCommand) *LifeCycle {
	return NewLifeCycle().OnMutationError(cmd...)
}

func (l *LifeCycle) AfterRequest(cmd ...JsCommand) *LifeCycle {
	l.OnEvent(HxAfterRequest, cmd...)
	return l
}

func (l *LifeCycle) OnMutationError(cmd ...JsCommand) *LifeCycle {
	l.OnEvent(HxOnMutationError, cmd...)
	return l
}

func (l *LifeCycle) Render() *Node {
	m := make(map[string]string)

	for event, commands := range l.handlers {
		m[event] = ""
		for _, command := range commands {
			m[event] += fmt.Sprintf("%s;", command.Command)
		}
	}

	children := make([]Renderable, 0)

	for event, js := range m {
		children = append(children, Attribute(event, js))
	}

	return Children(children...).Render()
}

type JsCommand struct {
	Command string
}

func SetText(text string) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.innerText = '%s'", text)}
}

func Increment(amount int) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.innerText = parseInt(this.innerText) + %d", amount)}
}

func SetInnerHtml(r Renderable) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.innerHTML = `%s`", Render(r.Render()))}
}

func SetOuterHtml(r Renderable) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.outerHTML = `%s`", Render(r.Render()))}
}

func AddAttribute(name, value string) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.setAttribute('%s', '%s')", name, value)}
}

func SetDisabled(disabled bool) JsCommand {
	if disabled {
		return AddAttribute("disabled", "true")
	} else {
		return RemoveAttribute("disabled")
	}
}

func RemoveAttribute(name string) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.removeAttribute('%s')", name)}
}

func AddClass(class string) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.classList.add('%s')", class)}
}

func RemoveClass(class string) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.classList.remove('%s')", class)}
}

func Alert(text string) JsCommand {
	return JsCommand{Command: fmt.Sprintf("alert('%s')", text)}
}

func EvalJs(js string) JsCommand {
	return JsCommand{Command: js}
}
