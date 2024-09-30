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

func OnLoad(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.LoadDomEvent, cmd...)
}

func (l *LifeCycle) HxBeforeRequest(cmd ...Command) *LifeCycle {
	l.OnEvent(hx.BeforeRequestEvent, cmd...)
	return l
}

func (l *LifeCycle) HxBeforeWsSend(cmd ...Command) *LifeCycle {
	l.OnEvent(hx.BeforeWsSendEvent, cmd...)
	return l
}

func (l *LifeCycle) HxAfterWsSend(cmd ...Command) *LifeCycle {
	l.OnEvent(hx.AfterWsSendEvent, cmd...)
	return l
}

func HxOnLoad(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.LoadEvent, cmd...)
}

func HxOnAfterSwap(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.AfterSwapEvent, cmd...)
}

func OnClick(cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(hx.ClickEvent, cmd...)
}

func OnEvent(event hx.Event, cmd ...Command) *LifeCycle {
	return NewLifeCycle().OnEvent(event, cmd...)
}

func HxBeforeWsSend(cmd ...Command) *LifeCycle {
	return NewLifeCycle().HxBeforeWsSend(cmd...)
}

func HxAfterWsSend(cmd ...Command) *LifeCycle {
	return NewLifeCycle().HxAfterWsSend(cmd...)
}

func HxBeforeRequest(cmd ...Command) *LifeCycle {
	return NewLifeCycle().HxBeforeRequest(cmd...)
}

func HxAfterRequest(cmd ...Command) *LifeCycle {
	return NewLifeCycle().HxAfterRequest(cmd...)
}

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

func NewComplexJsCommand(command string) ComplexJsCommand {
	name := fmt.Sprintf("__eval_%s", util.RandSeq(6))
	return ComplexJsCommand{Command: command, TempFuncName: name}
}

func SetText(text string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.innerText = '%s'", text)}
}

func SetTextOnChildren(selector, text string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		var children = self.querySelectorAll('%s');
		children.forEach(function(child) {
			child.innerText = '%s';
		});
	`, selector, text))
}

func Increment(amount int) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.innerText = parseInt(this.innerText) + %d", amount)}
}

func SetInnerHtml(r Ren) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.innerHTML = `%s`", Render(r))}
}

func SetOuterHtml(r Ren) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.outerHTML = `%s`", Render(r))}
}

func AddAttribute(name, value string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.setAttribute('%s', '%s')", name, value)}
}

func SetDisabled(disabled bool) SimpleJsCommand {
	if disabled {
		return AddAttribute("disabled", "true")
	} else {
		return RemoveAttribute("disabled")
	}
}

func RemoveAttribute(name string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.removeAttribute('%s')", name)}
}

func AddClass(class string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.classList.add('%s')", class)}
}

func RemoveClass(class string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.classList.remove('%s')", class)}
}

func ToggleClass(class string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.classList.toggle('%s')", class)}
}

func ToggleClassOnElement(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		var el = document.querySelector('%s');
		if(el) { el.classList.toggle('%s'); }`,
		selector, class,
	))
}

func EvalJsOnParent(js string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		if(!self.parentElement) { return; }
        let element = self.parentElement;
        %s
	`, js))
}

func EvalJsOnChildren(selector, js string) ComplexJsCommand {
	// language=JavaScript
	return EvalJs(fmt.Sprintf(`
		let children = self.querySelectorAll('%s');
		children.forEach(function(element) {
			%s
		});
	`, selector, js))
}

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

func SetClassOnParent(class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnParent(fmt.Sprintf("element.classList.add('%s')", class))
}

func RemoveClassOnParent(class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnParent(fmt.Sprintf("element.classList.remove('%s')", class))
}

func SetClassOnChildren(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnChildren(selector, fmt.Sprintf("element.classList.add('%s')", class))
}

func SetClassOnSibling(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnSibling(selector, fmt.Sprintf("element.classList.add('%s')", class))
}

func RemoveClassOnSibling(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnSibling(selector, fmt.Sprintf("element.classList.remove('%s')", class))

}

func RemoveClassOnChildren(selector, class string) ComplexJsCommand {
	// language=JavaScript
	return EvalJsOnChildren(selector, fmt.Sprintf("element.classList.remove('%s')", class))
}

func Alert(text string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("alert('%s')", text)}
}

func Remove() SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: "this.remove()"}
}

func EvalJs(js string) ComplexJsCommand {
	return NewComplexJsCommand(js)
}

func SetValue(value string) SimpleJsCommand {
	// language=JavaScript
	return SimpleJsCommand{Command: fmt.Sprintf("this.value = '%s'", value)}
}

func SubmitFormOnEnter() ComplexJsCommand {
	// language=JavaScript
	return EvalJs(`
		if (event.code === 'Enter') { 
			self.form.dispatchEvent(new Event('submit', { cancelable: true })); 
		}
	`)
}

func InjectScript(src string) ComplexJsCommand {
	// language=JavaScript
	return NewComplexJsCommand(fmt.Sprintf(`
		var script = document.createElement('script');
		script.src = '%s';
        script.async = true;
		document.head.appendChild(script);
	`, src))
}

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
