package h

func If(condition bool, node Ren) Ren {
	if condition {
		return node
	} else {
		return Empty()
	}
}

func Ternary[T any](value bool, a T, b T) T {
	if value {
		return a
	}
	return b
}

func IfElse(condition bool, node Ren, node2 Ren) Ren {
	if condition {
		return node
	} else {
		return node2
	}
}

func IfElseLazy(condition bool, cb1 func() Ren, cb2 func() Ren) Ren {
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
