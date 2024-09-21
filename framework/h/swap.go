package h

import (
	"fmt"
	"github.com/maddalax/htmgo/framework/hx"
)

type SwapArg struct {
	Content *Element
	Option  SwapOption
}

type SwapOption struct {
	Selector string
	SwapType hx.SwapType
	Modifier string
}

func NewSwap(content *Element, opts ...SwapOption) SwapArg {
	option := SwapOption{}
	if len(opts) > 0 {
		option = opts[0]
	}
	return SwapArg{
		Content: content,
		Option:  option,
	}
}

func OobSwap(ctx *RequestContext, content *Element, option ...SwapOption) *Element {
	return OobSwapWithSelector(ctx, "", content, option...)
}

func OobSwapWithSelector(ctx *RequestContext, selector string, content *Element, option ...SwapOption) *Element {
	if ctx == nil || !ctx.isHxRequest {
		return Empty()
	}
	return content.AppendChild(outOfBandSwap(selector, option...))
}

func outOfBandSwap(selector string, option ...SwapOption) Ren {
	swapType := hx.SwapTypeTrue

	if len(option) > 0 {
		o := option[0]

		if o.SwapType != "" {
			swapType = o.SwapType
		}

		modifier := o.Modifier
		if modifier != "" {
			swapType = fmt.Sprintf("%s %s", swapType, modifier)
		}
	}

	return Attribute(hx.SwapOobAttr,
		Ternary(selector == "", swapType, selector))
}

func SwapMany(ctx *RequestContext, elements ...*Element) *Element {
	if !ctx.isHxRequest {
		return Empty()
	}
	for _, element := range elements {
		element.AppendChild(outOfBandSwap(""))
	}
	return Template(Map(elements, func(arg *Element) Ren {
		return arg
	})...)
}

func SwapManyX(ctx *RequestContext, args ...SwapArg) *Element {
	if !ctx.isHxRequest {
		return Empty()
	}
	for _, arg := range args {
		arg.Content.AppendChild(outOfBandSwap("", arg.Option))
	}
	return Template(Map(args, func(arg SwapArg) Ren {
		return arg.Content
	})...)
}
