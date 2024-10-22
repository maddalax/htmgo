package h

import (
	"flag"
	"log/slog"
	"sync"
	"time"
)

type CachedNode struct {
	cb              func() *Element
	isByKey         bool
	byKeyCache      map[any]*Entry
	byKeyExpiration map[any]time.Time
	mutex           sync.Mutex
	duration        time.Duration
	expiration      time.Time
	html            string
}

type Entry struct {
	expiration time.Time
	html       string
}

type GetElementFunc func() *Element
type GetElementFuncT[T any] func(T) *Element
type GetElementFuncT2[T any, T2 any] func(T, T2) *Element
type GetElementFuncT3[T any, T2 any, T3 any] func(T, T2, T3) *Element
type GetElementFuncT4[T any, T2 any, T3 any, T4 any] func(T, T2, T3, T4) *Element

type GetElementFuncWithKey[K comparable] func() (K, GetElementFunc)
type GetElementFuncTWithKey[K comparable, T any] func(T) (K, GetElementFunc)
type GetElementFuncT2WithKey[K comparable, T any, T2 any] func(T, T2) (K, GetElementFunc)
type GetElementFuncT3WithKey[K comparable, T any, T2 any, T3 any] func(T, T2, T3) (K, GetElementFunc)
type GetElementFuncT4WithKey[K comparable, T any, T2 any, T3 any, T4 any] func(T, T2, T3, T4) (K, GetElementFunc)

func startExpiredCacheCleaner(node *CachedNode) {
	isTests := flag.Lookup("test.v") != nil
	go func() {
		for {
			if isTests {
				time.Sleep(time.Second)
			} else {
				time.Sleep(time.Minute)
			}
			node.ClearExpired()
		}
	}()
}

func Cached(duration time.Duration, cb GetElementFunc) func() *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			cb:       cb,
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
	return func() *Element {
		return element
	}
}

func CachedPerKey[K comparable](duration time.Duration, cb GetElementFuncWithKey[K]) func() *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			isByKey:  true,
			cb:       nil,
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
	return func() *Element {
		key, componentFunc := cb()
		return &Element{
			tag: CachedNodeByKeyEntry,
			meta: &ByKeyEntry{
				key:    key,
				parent: element,
				cb:     componentFunc,
			},
		}
	}
}

type ByKeyEntry struct {
	key    any
	cb     func() *Element
	parent *Element
}

func CachedPerKeyT[K comparable, T any](duration time.Duration, cb GetElementFuncTWithKey[K, T]) func(T) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			isByKey:  true,
			cb:       nil,
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
	return func(data T) *Element {
		key, componentFunc := cb(data)
		return &Element{
			tag: CachedNodeByKeyEntry,
			meta: &ByKeyEntry{
				key:    key,
				parent: element,
				cb:     componentFunc,
			},
		}
	}
}

func CachedPerKeyT2[K comparable, T any, T2 any](duration time.Duration, cb GetElementFuncT2WithKey[K, T, T2]) func(T, T2) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			isByKey:  true,
			cb:       nil,
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
	return func(data T, data2 T2) *Element {
		key, componentFunc := cb(data, data2)
		return &Element{
			tag: CachedNodeByKeyEntry,
			meta: &ByKeyEntry{
				key:    key,
				parent: element,
				cb:     componentFunc,
			},
		}
	}
}

func CachedPerKeyT3[K comparable, T any, T2 any, T3 any](duration time.Duration, cb GetElementFuncT3WithKey[K, T, T2, T3]) func(T, T2, T3) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			isByKey:  true,
			cb:       nil,
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
	return func(data T, data2 T2, data3 T3) *Element {
		key, componentFunc := cb(data, data2, data3)
		return &Element{
			tag: CachedNodeByKeyEntry,
			meta: &ByKeyEntry{
				key:    key,
				parent: element,
				cb:     componentFunc,
			},
		}
	}
}

func CachedPerKeyT4[K comparable, T any, T2 any, T3 any, T4 any](duration time.Duration, cb GetElementFuncT4WithKey[K, T, T2, T3, T4]) func(T, T2, T3, T4) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			isByKey:  true,
			cb:       nil,
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
	return func(data T, data2 T2, data3 T3, data4 T4) *Element {
		key, componentFunc := cb(data, data2, data3, data4)
		return &Element{
			tag: CachedNodeByKeyEntry,
			meta: &ByKeyEntry{
				key:    key,
				parent: element,
				cb:     componentFunc,
			},
		}
	}
}

func CachedT[T any](duration time.Duration, cb GetElementFuncT[T]) func(T) *Element {
	element := &Element{
		tag: CachedNodeTag,
		meta: &CachedNode{
			html:     "",
			duration: duration,
			mutex:    sync.Mutex{},
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
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
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
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
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
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
			html:     "",
			duration: duration,
		},
	}
	startExpiredCacheCleaner(element.meta.(*CachedNode))
	return func(data T, data2 T2, data3 T3, data4 T4) *Element {
		element.meta.(*CachedNode).cb = func() *Element {
			return cb(data, data2, data3, data4)
		}
		return element
	}
}

func (c *CachedNode) ClearCache() {
	c.html = ""
	if c.byKeyCache != nil {
		for key := range c.byKeyCache {
			delete(c.byKeyCache, key)
		}
	}
	if c.byKeyExpiration != nil {
		for key := range c.byKeyExpiration {
			delete(c.byKeyExpiration, key)
		}
	}
}

func (c *CachedNode) ClearExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	deletedCount := 0
	if c.isByKey {
		if c.byKeyCache != nil && c.byKeyExpiration != nil {
			for key := range c.byKeyCache {
				expir, ok := c.byKeyExpiration[key]
				if ok && expir.Before(time.Now()) {
					delete(c.byKeyCache, key)
					delete(c.byKeyExpiration, key)
					deletedCount++
				}
			}
		}
	} else {
		now := time.Now()
		expiration := c.expiration
		if c.html != "" && expiration.Before(now) {
			c.html = ""
			deletedCount++
		}
	}

	if deletedCount > 0 {
		slog.Debug("Deleted expired cache entries", slog.Int("count", deletedCount))
	}
}

func (c *CachedNode) Render(ctx *RenderContext) {
	if c.isByKey {
		panic("CachedPerKey should not be rendered directly")
	} else {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		now := time.Now()
		expiration := c.expiration

		if expiration.IsZero() || expiration.Before(now) {
			c.html = ""
			c.expiration = now.Add(c.duration)
		}

		if c.html != "" {
			ctx.builder.WriteString(c.html)
		} else {
			c.html = Render(c.cb())
			ctx.builder.WriteString(c.html)
		}
	}
}

func (c *ByKeyEntry) Render(ctx *RenderContext) {
	key := c.key
	parentMeta := c.parent.meta.(*CachedNode)

	parentMeta.mutex.Lock()
	defer parentMeta.mutex.Unlock()

	if parentMeta.byKeyCache == nil {
		parentMeta.byKeyCache = make(map[any]*Entry)
	}

	if parentMeta.byKeyExpiration == nil {
		parentMeta.byKeyExpiration = make(map[any]time.Time)
	}

	var setAndWrite = func() {
		html := Render(c.cb())
		parentMeta.byKeyCache[key] = &Entry{
			expiration: parentMeta.expiration,
			html:       html,
		}
		ctx.builder.WriteString(html)
	}

	expEntry, ok := parentMeta.byKeyExpiration[key]

	if !ok {
		parentMeta.byKeyExpiration[key] = time.Now().Add(parentMeta.duration)
	} else {
		// key is expired
		if expEntry.Before(time.Now()) {
			parentMeta.byKeyExpiration[key] = time.Now().Add(parentMeta.duration)
			setAndWrite()
			return
		}
	}

	entry := parentMeta.byKeyCache[key]

	// not in cache
	if entry == nil {
		setAndWrite()
		return
	}

	// exists in cache and not expired
	ctx.builder.WriteString(entry.html)
}
