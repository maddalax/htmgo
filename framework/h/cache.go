package h

import (
	"sync"
	"time"
)

type CachedNode struct {
	cb         func() *Element
	mutex      sync.Mutex
	expiration time.Time
	html       string
}

type GetElementFunc func() *Element
type GetElementFuncT[T any] func(T) *Element
type GetElementFuncT2[T any, T2 any] func(T, T2) *Element
type GetElementFuncT3[T any, T2 any, T3 any] func(T, T2, T3) *Element
type GetElementFuncT4[T any, T2 any, T3 any, T4 any] func(T, T2, T3, T4) *Element

func Cached(duration time.Duration, cb GetElementFunc) func() *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			cb:         cb,
			html:       "",
			expiration: time.Now().Add(duration),
		},
	}
	return func() *Element {
		return element
	}
}

func CachedT[T any](duration time.Duration, cb GetElementFuncT[T]) func(T) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			html:       "",
			expiration: time.Now().Add(duration),
			mutex:      sync.Mutex{},
		},
	}
	return func(data T) *Element {
		element.meta.(*CachedNode).cb = func() *Element {
			return cb(data)
		}
		return element
	}
}

func CachedT2[T any, T2 any](duration time.Duration, cb GetElementFuncT2[T, T2]) func(T, T2) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			html:       "",
			expiration: time.Now().Add(duration),
		},
	}
	return func(data T, data2 T2) *Element {
		element.meta.(*CachedNode).cb = func() *Element {
			return cb(data, data2)
		}
		return element
	}
}

func CachedT3[T any, T2 any, T3 any](duration time.Duration, cb GetElementFuncT3[T, T2, T3]) func(T, T2, T3) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			html:       "",
			expiration: time.Now().Add(duration),
		},
	}
	return func(data T, data2 T2, data3 T3) *Element {
		element.meta.(*CachedNode).cb = func() *Element {
			return cb(data, data2, data3)
		}
		return element
	}
}

func CachedT4[T any, T2 any, T3 any, T4 any](duration time.Duration, cb GetElementFuncT4[T, T2, T3, T4]) func(T, T2, T3, T4) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			html:       "",
			expiration: time.Now().Add(duration),
		},
	}
	return func(data T, data2 T2, data3 T3, data4 T4) *Element {
		element.meta.(*CachedNode).cb = func() *Element {
			return cb(data, data2, data3, data4)
		}
		return element
	}
}

func (c *CachedNode) ClearCache() {
	c.html = ""
}

func (c *CachedNode) Render(ctx *RenderContext) {
	c.mutex.Lock()
	if c.expiration.Before(time.Now()) {
		c.html = ""
	}

	if c.html != "" {
		ctx.builder.WriteString(c.html)
	} else {
		c.html = Render(c.cb())
		ctx.builder.WriteString(c.html)
	}
	c.mutex.Unlock()
}
