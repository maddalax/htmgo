package h

func If(condition bool, node Ren) Ren {
	if condition {
		return node
	} else {
		return Empty()
	}
}

func Ternary[T any](value bool, a T, b T) T {
	return IfElse(value, a, b)
}

func ElementIf(condition bool, element *Element) *Element {
	if condition {
		return element
	} else {
		return Empty()
	}
}

func IfElse[T any](condition bool, node T, node2 T) T {
	if condition {
		return node
	} else {
		return node2
	}
}

func IfElseLazy[T any](condition bool, cb1 func() T, cb2 func() T) T {
	if condition {
		return cb1()
	} else {
		return cb2()
	}
}

func IfHtmxRequest(ctx *RequestContext, node Ren) Ren {
	if ctx.Get("HX-Request") != "" {
		return node
	}
	return Empty()
}

func ClassIf(condition bool, value string) Ren {
	if condition {
		return Class(value)
	}
	return Empty()
}

func AttributeIf(condition bool, name string, value string) Ren {
	if condition {
		return Attribute(name, value)
	}
	return Empty()
}
