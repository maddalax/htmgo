package h

import (
	"fmt"
)

var HxBeforeRequest = "hx-on::before-request"
var HxAfterRequest = "hx-on::after-request"
var HxOnMutationError = "hx-on::mutation-error"

type LifeCycle struct {
	beforeRequest   []JsCommand
	afterRequest    []JsCommand
	onMutationError []JsCommand
}

func NewLifeCycle() *LifeCycle {
	return &LifeCycle{
		beforeRequest:   []JsCommand{},
		afterRequest:    []JsCommand{},
		onMutationError: []JsCommand{},
	}
}

func (l *LifeCycle) BeforeRequest(cmd ...JsCommand) *LifeCycle {
	l.beforeRequest = append(l.beforeRequest, cmd...)
	return l
}

func (l *LifeCycle) AfterRequest(cmd ...JsCommand) *LifeCycle {
	l.afterRequest = append(l.afterRequest, cmd...)
	return l
}

func (l *LifeCycle) OnMutationError(cmd ...JsCommand) *LifeCycle {
	l.onMutationError = append(l.onMutationError, cmd...)
	return l
}

func (l *LifeCycle) Render() *Node {
	beforeRequest := ""
	afterReqest := ""
	onMutationError := ""
	for _, command := range l.beforeRequest {
		beforeRequest += fmt.Sprintf("%s;", command.Command)
	}
	for _, command := range l.afterRequest {
		afterReqest += fmt.Sprintf("%s;", command.Command)
	}
	for _, command := range l.onMutationError {
		onMutationError += fmt.Sprintf("%s;", command.Command)
	}

	return Children(
		If(beforeRequest != "", Attribute(HxBeforeRequest, beforeRequest)),
		If(afterReqest != "", Attribute(HxAfterRequest, afterReqest)),
		If(onMutationError != "", Attribute(HxOnMutationError, onMutationError)),
	).Render()
}

type JsCommand struct {
	Command string
}

func SetText(text string) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.innerText = '%s'", text)}
}

func AddAttribute(name, value string) JsCommand {
	return JsCommand{Command: fmt.Sprintf("this.setAttribute('%s', '%s')", name, value)}
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
