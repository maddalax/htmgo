package h

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/hx"
	"github.com/maddalax/htmgo/framework/internal/util"
	"strings"
)

type LifeCycle struct {
	handlers map[hx.Event][]Command
}

func NewLifeCycle() *LifeCycle {
	return &LifeCycle{
		handlers: make(map[hx.Event][]Command),
	}
}

func validateCommands(cmds []Command) {
	for _, cmd := range cmds {
		switch t := cmd.(type) {
		case SimpleJsCommand:
			break
		case ComplexJsCommand:
			break
		case *AttributeMapOrdered:
			break
		case *Element:
			panic(fmt.Sprintf("element is not allowed in lifecycle events. Got: %v", t))
		default:
			panic(fmt.Sprintf("type is not allowed in lifecycle events. Got: %v", t))
		}
	}
}

func (l *LifeCycle) OnEvent(event hx.Event, cmd ...Command) *LifeCycle {
	validateCommands(cmd)

	if strings.HasPrefix(event, "htmx:") {
		event = event[5:]
		event = util.ConvertCamelToDash(fmt.Sprintf("hx-on::%s", event))
	}

	if l.handlers[event] == nil {
		l.handlers[event] = []Command{}
	}

	l.handlers[event] = append(l.handlers[event], cmd...)
	return l
}

// OnLoad executes the given commands when the element is loaded into the DOM, it also executes when the element is replaced / swapped in.
// This will work on any element because of the htmgo htmx extension to trigger it, instead of the browser.
func OnLoad(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.LoadDomEvent, cmd...)
}

func (l *LifeCycle) HxBeforeRequest(cmd ...Command) *LifeCycle {
	l.OnEvent(hx.BeforeRequestEvent, cmd...)
	return l
}

// HxOnLoad executes the given commands when the element is loaded into the DOM.
// Deprecated: Use OnLoad instead.
func HxOnLoad(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.LoadEvent, cmd...)
}

// HxOnAfterSwap executes the given commands when the element is swapped in.
func HxOnAfterSwap(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.AfterSwapEvent, cmd...)
}

// OnClick executes the given commands when the element is clicked.
func OnClick(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.ClickEvent, cmd...)
}

// OnEvent executes the given commands when the given event is triggered.
func OnEvent(event hx.Event, cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(event, cmd...)
}

// HxBeforeSseMessage executes the given commands when a message is received from the server via SSE, but before it is processed.
func HxBeforeSseMessage(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.SseBeforeMessageEvent, cmd...)
}

// HxAfterSseMessage executes the given commands when a message is received from the server via SSE, and after it is processed.
func HxAfterSseMessage(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.SseAfterMessageEvent, cmd...)
}

// OnSubmit executes the given commands when the form is submitted.
func OnSubmit(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.SubmitEvent, cmd...)
}

// HxOnSseError executes the given commands when an error occurs while connecting to the server via SSE.
func HxOnSseError(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.SseErrorEvent, cmd...)
}

// HxOnSseClose executes the given commands when the connection to the server via SSE is closed.
func HxOnSseClose(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.SseClosedEvent, cmd...)
}

// HxOnSseConnecting executes the given commands when the connection to the server via SSE is being established.
func HxOnSseConnecting(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.SseConnectingEvent, cmd...)
}

// HxOnSseOpen executes the given commands when the connection to the server via SSE is established.
func HxOnSseOpen(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.SseConnectedEvent, cmd...)
}

// HxBeforeRequest executes the given commands before the request is sent.
func HxBeforeRequest(cmd ...Command) *LifeCycle {
	return NewLifeCycle().HxBeforeRequest(cmd...)
}

// HxAfterRequest executes the given commands after the request is sent.
func HxAfterRequest(cmd ...Command) *LifeCycle {
	return NewLifeCycle().HxAfterRequest(cmd...)
}

// HxOnMutationError executes the given commands when a mutation error of a request occurs.
func HxOnMutationError(cmd ...Command) *LifeCycle {
	return NewLifeCycle().HxOnMutationError(cmd...)
}

func (l *LifeCycle) HxAfterRequest(cmd ...Command) *LifeCycle {
	l.OnEvent(hx.AfterRequestEvent, cmd...)
	return l
}

func (l *LifeCycle) HxOnMutationError(cmd ...Command) *LifeCycle {
	l.OnEvent(hx.OnMutationErrorEvent, cmd...)
	return l
}

type Command = Ren

type SimpleJsCommand struct {
	Command string
}

type ComplexJsCommand struct {
	Command      string
	TempFuncName string
}

// NewComplexJsCommand creates a new complex JavaScript command.
func NewComplexJsCommand(command string) ComplexJsCommand {
	name := fmt.Sprintf("__eval_%s", util.RandSeq(6))
	return ComplexJsCommand{Command: command, TempFuncName: name}
}

// SetText sets the inner text of the element.
func SetText(text string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.innerText = '%s'", text)}
}

// SetTextOnChildren sets the inner text of all the children of the element that match the selector.
func SetTextOnChildren(selector, text string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		var children = self.querySelectorAll('%s');
		children.forEach(function(child) {
			child.innerText = '%s';
		});
	`, selector, text))
}

// Increment increments the inner text of the element by the given amount.
func Increment(amount int) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.innerText = parseInt(this.innerText) + %d", amount)}
}

// SetInnerHtml sets the inner HTML of the element.
func SetInnerHtml(r Ren) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.innerHTML = `%s`", Render(r))}
}

// SetOuterHtml sets the outer HTML of the element.
func SetOuterHtml(r Ren) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.outerHTML = `%s`", Render(r))}
}

// AddAttribute adds the given attribute to the element.
func AddAttribute(name, value string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.setAttribute('%s', '%s')", name, value)}
}

// SetDisabled sets the disabled attribute on the element.
func SetDisabled(disabled bool) SimpleJsCommand {
	if disabled {
		return AddAttribute("disabled", "true")
	} else {
		return RemoveAttribute("disabled")
	}
}

// RemoveAttribute removes the given attribute from the element.
func RemoveAttribute(name string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.removeAttribute('%s')", name)}
}

// AddClass adds the given class to the element.
func AddClass(class string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.classList.add('%s')", class)}
}

// RemoveClass removes the given class from the element.
func RemoveClass(class string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.classList.remove('%s')", class)}
}

// ToggleClass toggles the given class on the element.
func ToggleClass(class string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.classList.toggle('%s')", class)}
}

// ToggleClassOnElement toggles the given class on the elements returned by the selector.
func ToggleClassOnElement(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		var el = document.querySelector('%s');
		if(el) { el.classList.toggle('%s'); }`,
		selector, class,
	))
}

// EvalJsOnParent evaluates the given JavaScript code on the parent of the element. Reference the element using 'element'.
func EvalJsOnParent(js string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		if(!self.parentElement) { return; }
        let element = self.parentElement;
        %s
	`, js))
}

// EvalJsOnChildren evaluates the given JavaScript code on the children of the element. Reference the element using 'element'.
func EvalJsOnChildren(selector, js string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		let children = self.querySelectorAll('%s');
		children.forEach(function(element) {
			%s
		});
	`, selector, js))
}

// EvalJsOnSibling evaluates the given JavaScript code on the siblings of the element. Reference the element using 'element'.
func EvalJsOnSibling(selector, js string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		if(!self.parentElement) { return; }
		let siblings = self.parentElement.querySelectorAll('%s');
		siblings.forEach(function(element) {
			%s
		});
	`, selector, js))
}

// SetClassOnParent sets the given class on the parent of the element. Reference the element using 'element'.
func SetClassOnParent(class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnParent(fmt.Sprintf("element.classList.add('%s')", class))
}

// RemoveClassOnParent removes the given class from the parent of the element. Reference the element using 'element'.
func RemoveClassOnParent(class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnParent(fmt.Sprintf("element.classList.remove('%s')", class))
}

// SetClassOnChildren sets the given class on the children of the element. Reference the element using 'element'.
func SetClassOnChildren(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnChildren(selector, fmt.Sprintf("element.classList.add('%s')", class))
}

// SetClassOnSibling sets the given class on the siblings of the element. Reference the element using 'element'.
func SetClassOnSibling(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnSibling(selector, fmt.Sprintf("element.classList.add('%s')", class))
}

// RemoveClassOnSibling removes the given class from the siblings of the element. Reference the element using 'element'.
func RemoveClassOnSibling(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnSibling(selector, fmt.Sprintf("element.classList.remove('%s')", class))

}

// RemoveClassOnChildren removes the given class from the children of the element. Reference the element using 'element'.
func RemoveClassOnChildren(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnChildren(selector, fmt.Sprintf("element.classList.remove('%s')", class))
}

// Alert displays an alert dialog with the given text.
func Alert(text string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("alert('%s')", text)}
}

// Remove removes the element from the DOM.
func Remove() SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: "this.remove()"}
}

// EvalJs evaluates the given JavaScript code.
func EvalJs(js string) ComplexJsCommand {
	return NewComplexJsCommand(js)
}

// PreventDefault prevents the default action of the event.
func PreventDefault() SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: "event.preventDefault()"}
}

// ConsoleLog logs a message to the console.
func ConsoleLog(text string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("console.log('%s')", text)}
}

// SetValue sets the value of the element.
func SetValue(value string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.value = '%s'", value)}
}

// SubmitFormOnEnter submits the form when the user presses the enter key.
func SubmitFormOnEnter() ComplexJsCommand {
	// language=JavaScript
	return EvalJs(`
		if (event.code === 'Enter') { self.form.dispatchEvent(new Event('submit', { cancelable: true })); }
	`)
}

// InjectScript injects a script tag into the document.
func InjectScript(src string) ComplexJsCommand {
	// language=JavaScript
	return NewComplexJsCommand(fmt.Sprintf(`
		var script = document.createElement('script');
		script.src = '%s';
        script.async = true;
		document.head.appendChild(script);
	`, src))
}

// InjectScriptIfNotExist injects a script tag into the document if it does not already exist.
func InjectScriptIfNotExist(src string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		if(!document.querySelector('script[src="%s"]')) {
			var script = document.createElement('script');
			script.src = '%s';
			script.async = true;
			document.head.appendChild(script);
		}
	`, src, src))
}
