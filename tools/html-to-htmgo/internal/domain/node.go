package domain

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CustomNode struct {
	ParentNode *CustomNode
	Level      uint
	customType bool
	Type       string
	Attrs      []Attr
	Nodes      []*CustomNode
}

func (n *CustomNode) SetType(in string) {
	switch in {
	case "textarea":
		n.Type = "h.TextArea"
	case "head":
		n.Type = "h.Head"
	case "thead":
		n.Type = "h.THead"
	case "tbody":
		n.Type = "h.TBody"
	case "id":
		n.Type = "h.Id"
	case "circle":
		n.Type = "circle"
		n.customType = true
	case "rect":
		n.Type = "rect"
		n.customType = true
	case "line":
		n.Type = "line"
		n.customType = true
	case "polyline":
		n.Type = "line"
		n.customType = true
	case "svg":
		n.Type = "h.Svg"
	default:
		n.Type = fmt.Sprintf("h.%s", cases.Title(language.English).String(in))
	}
}

func (n *CustomNode) AddAttr(key, value string) {
	if slices.Contains([]string{"xmlns", "fill", "viewBox", "stroke", "stroke-width", "fill-rule", "d", "stroke-linecap", "stroke-linejoin", "cx", "cy", "r", "x", "y", "rx", "ry", "x1", "x2", "y1", "y2", "points"}, key) {
		n.Attrs = append(n.Attrs, Attr{
			custom: true,
			key:    key,
			value:  value,
		})
		return
	}

	switch {
	case key == "autocomplete":
		n.Attrs = append(n.Attrs, Attr{key: "h.AutoComplete", value: value})
	case key == "id":
		n.Attrs = append(n.Attrs, Attr{key: "h.Id", value: value})
	case key == "tabindex":
		n.Attrs = append(n.Attrs, Attr{key: "h.TabIndex", value: value})
	case key == "h.Text":
		n.Attrs = append(n.Attrs, Attr{key: key, value: value})
	case strings.ContainsRune(key, '-'):
		n.Attrs = append(n.Attrs, Attr{
			custom: true,
			key:    key,
			value:  value,
		})
		fmt.Printf("key: %s, value: %s\n", key, value)
	default:
		n.Attrs = append(n.Attrs, Attr{key: "h." + cases.Title(language.English).String(key), value: value})
	}
}

func (n *CustomNode) String() string {
	str := ""

	if n.customType {
		str += "h.Tag(\"" + n.Type + "\","
	} else {
		str += n.Type + "("
	}

	if str == "h.Input(" {
		if len(n.Attrs) > 0 {
			for i, attr := range n.Attrs {
				if attr.key == "h.Type" {
					str = str + fmt.Sprintf(`"%s"`, attr.value) + ","
					n.Attrs = append(n.Attrs[:i], n.Attrs[i+1:]...)
				}
			}
		}
	}

	if str == "h.Script(" {
		if len(n.Attrs) > 0 {
			for _, attr := range n.Attrs {
				if attr.key == "h.Src" {
					str = str + fmt.Sprintf(`"%s"`, attr.value) + ","
					n.Attrs = make([]Attr, 0)
				}
			}
		}
	}

	booleanAttributes := []string{
		"h.AllowFullscreen",
		"h.Async",
		"h.Autofocus",
		"h.Autoplay",
		"h.Checked",
		"h.Controls",
		"h.Default",
		"h.Defer",
		"h.Disabled",
		"h.FormNoValidate",
		"h.Hidden",
		"h.IsMap",
		"h.Loop",
		"h.Multiple",
		"h.Muted",
		"h.NoModule",
		"h.NoValidate",
		"h.Open",
		"h.ReadOnly",
		"h.Required",
		"h.Reversed",
		"h.Selected",
	}

	if len(n.Attrs) > 0 {
		for _, v := range n.Attrs {
			switch {
			case v.custom:
				str = fmt.Sprintf("%sh.Attribute(\"%s\",\"%s\"),", str, v.key, v.value)
			case v.hyphenated:
				str = fmt.Sprintf("%s%s(\"%s\", \"%s\"),", str, v.key, v.arg, v.value)
			case len(v.value) > 0:
				if strings.Contains(v.value, "\n") {
					str = fmt.Sprintf("%s%s(`%s`),", str, v.key, v.value)
				} else {
					str = fmt.Sprintf("%s%s(\"%s\"),", str, v.key, v.value)
				}
			case v.value == "" && !slices.Contains(booleanAttributes, v.key):
				str = fmt.Sprintf("%s%s(\"\"),", str, v.key)
			default:
				str = fmt.Sprintf("%s%s(),", str, v.key)
			}
		}
	}

	if len(n.Nodes) > 0 {
		for _, v := range n.Nodes {
			if v.Type != "" {
				str = fmt.Sprintf("%s\n%s%s,", str, strings.Repeat(" ", int(n.Level)), v)
			}
		}
	}

	str = fmt.Sprintf("%s\n%s)", str, strings.Repeat(" ", int(n.Level)))
	return str
}

type Attr struct {
	custom, hyphenated bool
	key, value, arg    string
}
