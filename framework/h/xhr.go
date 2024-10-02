package h

import (
	"github.com/maddalax/htmgo/framework/hx"
	"strings"
)

func Get(path string, trigger ...string) *AttributeMapOrdered {
	return AttributeList(Attribute(hx.GetAttr, path), HxTriggerString(trigger...))
}

func GetPartial(partial PartialFunc, trigger ...string) *AttributeMapOrdered {
	return Get(GetPartialPath(partial), trigger...)
}

func GetPartialWithQs(partial PartialFunc, qs *Qs, trigger string) *AttributeMapOrdered {
	return Get(GetPartialPathWithQs(partial, qs), trigger)
}

func GetWithQs(path string, qs *Qs, trigger string) *AttributeMapOrdered {
	return Get(SetQueryParams(path, qs), trigger)
}

func PostPartial(partial PartialFunc, triggers ...string) *AttributeMapOrdered {
	path := GetPartialPath(partial)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return Post(path, triggers...)
}

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

func PostWithQs(url string, qs *Qs, trigger string) *AttributeMapOrdered {
	return Post(SetQueryParams(url, qs), trigger)
}

func PostOnClick(url string) *AttributeMapOrdered {
	return Post(url, hx.ClickEvent)
}

func PostPartialOnClick(partial PartialFunc) *AttributeMapOrdered {
	return PostOnClick(GetPartialPath(partial))
}

func PostPartialOnClickQs(partial PartialFunc, qs *Qs) *AttributeMapOrdered {
	return PostOnClick(GetPartialPathWithQs(partial, qs))
}
