package h

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"regexp"
	"strings"
	"testing"
)

func findScriptById(n *html.Node, id string) *html.Node {
	if n.Type == html.ElementNode && n.Data == "script" {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return n
			}
		}
	}
	// Recursively search in child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findScriptById(c, id); result != nil {
			return result
		}
	}
	return nil
}

func renderJs(t *testing.T, command Command) string {
	result := Render(Div(OnClick(command)))
	parsed, _ := html.Parse(strings.NewReader(result))
	value := parsed.FirstChild.FirstChild.NextSibling.LastChild.Attr[0].Val
	isComplex := strings.HasPrefix(value, "__eval_")
	if !isComplex {
		value = strings.ReplaceAll(value, "var e=event;", "")
		return strings.ReplaceAll(value, "var self=this;", "")
	} else {
		id := strings.TrimSuffix(value, "(this, event);")
		script := findScriptById(parsed, id)
		assert.NotNil(t, script)
		funcCall := script.LastChild.Data
		funcCall = strings.ReplaceAll(funcCall, "\n", "")
		funcCall = strings.ReplaceAll(funcCall, "\t", "")
		start := fmt.Sprintf("function %s(self, event) {", id)
		funcCall = strings.TrimPrefix(funcCall, start)
		funcCall = strings.TrimSuffix(funcCall, "}")
		funcCall = strings.ReplaceAll(funcCall, "let e = event;", "")
		funcCall = strings.ReplaceAll(funcCall, "var self=this;", "")
		return funcCall
	}
}

var re = regexp.MustCompile(`\s+`)

func compareIgnoreSpaces(t *testing.T, actual, expected string) {
	expected = strings.ReplaceAll(expected, "\n", "")
	expected = strings.ReplaceAll(expected, "\t", "")
	expected = re.ReplaceAllString(expected, " ")
	actual = strings.ReplaceAll(actual, "\n", "")
	actual = strings.ReplaceAll(actual, "\t", "")
	actual = re.ReplaceAllString(actual, " ")
	assert.Equal(t, expected, actual)
}

func TestJsEval(t *testing.T) {
	evalTests := []string{
		"alert('hello')",
		"console.log('hello')",
		"document.getElementById('myDiv').style.display = 'none'",
		"self.style.display = 'none'",
	}

	for _, test := range evalTests {
		compareIgnoreSpaces(t, renderJs(t, EvalJs(test)), test)
	}

	compareIgnoreSpaces(t, renderJs(t, EvalJsOnParent("element.style.display = 'none'")), `
		if(!self.parentElement) { return; } let element = self.parentElement; element.style.display = 'none'
	`)

	compareIgnoreSpaces(t, renderJs(t, EvalJsOnSibling("button", "element.style.display = 'none'")), `
		if(!self.parentElement) { return; }let siblings = self.parentElement.querySelectorAll('button');siblings.forEach(function(element) {element.style.display = 'none'});
	`)

}

func TestSetText(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, SetText("Hello World")), "this.innerText = 'Hello World';")
}

func TestSetTextOnChildren(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, SetTextOnChildren("div", "Hello Child")), `
		var children = self.querySelectorAll('div');
		children.forEach(function(child) {
			child.innerText = 'Hello Child';
		});
	`)
}

func TestIncrement(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, Increment(5)), "this.innerText = parseInt(this.innerText) + 5;")
}

func TestSetInnerHtml(t *testing.T) {
	htmlContent := Div(Span(UnsafeRaw("inner content")))
	compareIgnoreSpaces(t, renderJs(t, SetInnerHtml(htmlContent)), "this.innerHTML = `<div><span>inner content</span></div>`;")
}

func TestSetOuterHtml(t *testing.T) {
	htmlContent := Div(Span(UnsafeRaw("outer content")))
	compareIgnoreSpaces(t, renderJs(t, SetOuterHtml(htmlContent)), "this.outerHTML = `<div><span>outer content</span></div>`;")
}

func TestAddAttribute(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, AddAttribute("data-id", "123")), "this.setAttribute('data-id', '123');")
}

func TestSetDisabled(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, SetDisabled(true)), "this.setAttribute('disabled', 'true');")
	compareIgnoreSpaces(t, renderJs(t, SetDisabled(false)), "this.removeAttribute('disabled');")
}

func TestRemoveAttribute(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, RemoveAttribute("data-id")), "this.removeAttribute('data-id');")
}

func TestAddClass(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, AddClass("active")), "this.classList.add('active');")
}

func TestRemoveClass(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, RemoveClass("active")), "this.classList.remove('active');")
}

func TestToggleClass(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, ToggleClass("hidden")), "this.classList.toggle('hidden');")
}

func TestToggleClassOnElement(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, ToggleClassOnElement(".my-class", "hidden")), `
		var el = document.querySelector('.my-class');
		if(el) { el.classList.toggle('hidden'); }
	`)
}

func TestSetClassOnParent(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, SetClassOnParent("active")), `
		if(!self.parentElement) { return; } let element = self.parentElement; element.classList.add('active')
	`)
}

func TestRemoveClassOnParent(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, RemoveClassOnParent("active")), `
		if(!self.parentElement) { return; } let element = self.parentElement; element.classList.remove('active')
	`)
}

func TestSetClassOnChildren(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, SetClassOnChildren("div", "highlight")), `
		let children = self.querySelectorAll('div');
		children.forEach(function(element) {
			element.classList.add('highlight')
		});
	`)
}

func TestRemoveClassOnChildren(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, RemoveClassOnChildren("div", "highlight")), `
		let children = self.querySelectorAll('div');
		children.forEach(function(element) {
			element.classList.remove('highlight')
		});
	`)
}

func TestSetClassOnSibling(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, SetClassOnSibling("button", "selected")), `
		if(!self.parentElement) { return; }let siblings = self.parentElement.querySelectorAll('button');
		siblings.forEach(function(element) {
			element.classList.add('selected')
		});
	`)
}

func TestRemoveClassOnSibling(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, RemoveClassOnSibling("button", "selected")), `
		if(!self.parentElement) { return; }let siblings = self.parentElement.querySelectorAll('button');
		siblings.forEach(function(element) {
			element.classList.remove('selected')
		});
	`)
}

func TestAlert(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, Alert("Warning")), "alert('Warning');")
}

func TestRemove(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, Remove()), "this.remove();")
}

func TestSubmitFormOnEnter(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, SubmitFormOnEnter()), `
		if (event.code === 'Enter') { 
			self.form.dispatchEvent(new Event('submit', { cancelable: true })); 
		}
	`)
}

func TestInjectScript(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, InjectScript("https://example.com/script.js")), `
		var script = document.createElement('script');
		script.src = 'https://example.com/script.js';
        script.async = true;
		document.head.appendChild(script);
	`)
}

func TestInjectScriptIfNotExist(t *testing.T) {
	compareIgnoreSpaces(t, renderJs(t, InjectScriptIfNotExist("https://example.com/script.js")), `
		if(!document.querySelector('script[src="https://example.com/script.js"]')) {
			var script = document.createElement('script');
			script.src = 'https://example.com/script.js';
			script.async = true;
			document.head.appendChild(script);
		}
	`)
}
