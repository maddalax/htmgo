package h

import (
	"github.com/maddalax/htmgo/framework/hx"
	"strings"
)

// Get adds two attributes to the element: hx-get and hx-trigger.
func Get(path string, trigger ...string) *AttributeMapOrdered {
	return AttributeList(Attribute(hx.GetAttr, path), HxTriggerString(trigger...))
}

// GetPartial adds two attributes to the element: hx-get and hx-trigger, and uses the partial path for the hx-get attribute.
func GetPartial(partial PartialFunc, trigger ...string) *AttributeMapOrdered {
	return Get(GetPartialPath(partial), trigger...)
}

// GetPartialWithQs adds two attributes to the element: hx-get and hx-trigger, and uses the partial path for the hx-get attribute. It also sets the query string parameters.
func GetPartialWithQs(partial PartialFunc, qs *Qs, trigger string) *AttributeMapOrdered {
	return Get(GetPartialPathWithQs(partial, qs), trigger)
}

// GetWithQs adds two attributes to the element: hx-get and hx-trigger, and uses the path for the hx-get attribute. It also sets the query string parameters.
func GetWithQs(path string, qs *Qs, trigger string) *AttributeMapOrdered {
	return Get(SetQueryParams(path, qs), trigger)
}

// PostPartial adds two attributes to the element: hx-post and hx-trigger, and uses the partial path for the hx-post attribute.
func PostPartial(partial PartialFunc, triggers ...string) *AttributeMapOrdered {
	path := GetPartialPath(partial)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return Post(path, triggers...)
}

// PostPartialWithQs adds two attributes to the element: hx-post and hx-trigger, and uses the partial path for the hx-post attribute. It also sets the query string parameters.
func PostPartialWithQs(partial PartialFunc, qs *Qs, trigger ...string) *AttributeMapOrdered {
	path := GetPartialPathWithQs(partial, qs)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return Post(GetPartialPathWithQs(partial, qs), trigger...)
}

func Post(url string, trigger ...string) *AttributeMapOrdered {
	return AttributeList(Attribute(hx.PostAttr, url), HxTriggerString(trigger...))
}

// PostWithQs adds two attributes to the element: hx-post and hx-trigger, and uses the path for the hx-post attribute. It also sets the query string parameters.
func PostWithQs(url string, qs *Qs, trigger string) *AttributeMapOrdered {
	return Post(SetQueryParams(url, qs), trigger)
}

// PostOnClick adds two attributes to the element: hx-post and hx-trigger, and uses the path for the hx-post attribute. It also sets the hx-trigger to hx-click.
func PostOnClick(url string) *AttributeMapOrdered {
	return Post(url, hx.ClickEvent)
}

// PostPartialOnClick adds two attributes to the element: hx-post and hx-trigger, and uses the partial path for the hx-post attribute. It also sets the hx-trigger to hx-click.
func PostPartialOnClick(partial PartialFunc) *AttributeMapOrdered {
	return PostOnClick(GetPartialPath(partial))
}

// PostPartialOnClickQs adds two attributes to the element: hx-post and hx-trigger, and uses the partial path for the hx-post attribute. It also sets the hx-trigger to hx-click. It also sets the query string parameters.
func PostPartialOnClickQs(partial PartialFunc, qs *Qs) *AttributeMapOrdered {
	return PostOnClick(GetPartialPathWithQs(partial, qs))
}
