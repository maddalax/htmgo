package h

import "github.com/maddalax/htmgo/framework/hx"

func Get(path string, trigger ...string) *AttributeMap {
	return AttributeList(Attribute(hx.GetAttr, path), TriggerString(trigger...))
}

func GetPartial(partial PartialFunc, trigger ...string) *AttributeMap {
	return Get(GetPartialPath(partial), trigger...)
}

func GetPartialWithQs(partial PartialFunc, qs *Qs, trigger string) *AttributeMap {
	return Get(GetPartialPathWithQs(partial, qs), trigger)
}

func GetWithQs(path string, qs *Qs, trigger string) *AttributeMap {
	return Get(SetQueryParams(path, qs), trigger)
}

func PostPartial(partial PartialFunc, triggers ...string) *AttributeMap {
	return Post(GetPartialPath(partial), triggers...)
}

func PostPartialWithQs(partial PartialFunc, qs *Qs, trigger ...string) *AttributeMap {
	return Post(GetPartialPathWithQs(partial, qs), trigger...)
}

func Post(url string, trigger ...string) *AttributeMap {
	return AttributeList(Attribute(hx.PostAttr, url), TriggerString(trigger...))
}

func PostOnClick(url string) *AttributeMap {
	return Post(url, hx.ClickEvent)
}

func PostPartialOnClick(partial PartialFunc) *AttributeMap {
	return PostOnClick(GetPartialPath(partial))
}

func PostPartialOnClickQs(partial PartialFunc, qs *Qs) *AttributeMap {
	return PostOnClick(GetPartialPathWithQs(partial, qs))
}
