package h

// If returns the node if the condition is true, otherwise returns an empty element
func If(condition bool, node Ren) Ren {
	if condition {
		return node
	} else {
		return Empty()
	}
}

// Ternary returns the first argument if the second argument is true, otherwise returns the third argument
func Ternary[T any](value bool, a T, b T) T {
	return IfElse(value, a, b)
}

// ElementIf returns the element if the condition is true, otherwise returns an empty element
func ElementIf(condition bool, element *Element) *Element {
	if condition {
		return element
	} else {
		return Empty()
	}
}

// IfElseE returns element if condition is true, otherwise returns element2
func IfElseE(condition bool, element *Element, element2 *Element) *Element {
	if condition {
		return element
	} else {
		return element2
	}
}

// IfElse returns node if condition is true, otherwise returns node2
func IfElse[T any](condition bool, node T, node2 T) T {
	if condition {
		return node
	} else {
		return node2
	}
}

// IfElseLazy returns node if condition is true, otherwise returns the result of cb2
// This is useful if you want to lazily evaluate a node based on a condition
// For example, If you are rendering a component that requires specific data,
// you can use this to only load the component if the data is available
func IfElseLazy[T any](condition bool, cb1 func() T, cb2 func() T) T {
	if condition {
		return cb1()
	} else {
		return cb2()
	}
}

// IfHtmxRequest returns the node if the request is an htmx request, otherwise returns an empty element
func IfHtmxRequest(ctx *RequestContext, node Ren) Ren {
	if ctx.isHxRequest {
		return node
	}
	return Empty()
}

// ClassIf returns the class attribute if the condition is true, otherwise returns an empty element
func ClassIf(condition bool, value string) Ren {
	if condition {
		return Class(value)
	}
	return Empty()
}

// AttributeIf returns the attribute if the condition is true, otherwise returns an empty element
func AttributeIf(condition bool, name string, value string) Ren {
	if condition {
		return Attribute(name, value)
	}
	return Empty()
}
