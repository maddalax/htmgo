package h

import (
	"time"

	"github.com/maddalax/htmgo/framework/h/cache"
)

// A single key to represent the cache entry for non-per-key components.
const _singleCacheKey = "__htmgo_single_cache_key__"

type CachedNode struct {
	cb       func() *Element
	isByKey  bool
	duration time.Duration
	cache    cache.Store[any, string]
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

// CacheOption defines a function that configures a CachedNode.
type CacheOption func(*CachedNode)

// WithStore allows providing a custom cache implementation for a cached component.
func WithStore(store cache.Store[any, string]) CacheOption {
	return func(c *CachedNode) {
		c.cache = store
	}
}

// DefaultCacheProvider is a package-level function that creates a default cache instance.
// Initially, this uses a TTL-based map cache, but could be swapped for an LRU cache later.
// Advanced users can override this for the entire application.
var DefaultCacheProvider = func() cache.Store[any, string] {
	return cache.NewTTLStore[any, string]()
}

// Cached caches the given element for the given duration. The element is only rendered once, and then cached for the given duration.
// Please note this element is globally cached, and not per unique identifier / user.
// Use CachedPerKey to cache elements per unique identifier.
func Cached(duration time.Duration, cb GetElementFunc, opts ...CacheOption) func() *Element {
	node := &CachedNode{
		cb:       cb,
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

	return func() *Element {
		return element
	}
}

// CachedPerKey caches the given element for the given duration. The element is only rendered once per key, and then cached for the given duration.
// The element is cached by the unique identifier that is returned by the callback function.
func CachedPerKey[K comparable](duration time.Duration, cb GetElementFuncWithKey[K], opts ...CacheOption) func() *Element {
	node := &CachedNode{
		isByKey:  true,
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

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

// CachedPerKeyT caches the given element for the given duration. The element is only rendered once per key, and then cached for the given duration.
// The element is cached by the unique identifier that is returned by the callback function.
func CachedPerKeyT[K comparable, T any](duration time.Duration, cb GetElementFuncTWithKey[K, T], opts ...CacheOption) func(T) *Element {
	node := &CachedNode{
		isByKey:  true,
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

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

// CachedPerKeyT2 caches the given element for the given duration. The element is only rendered once per key, and then cached for the given duration.
// The element is cached by the unique identifier that is returned by the callback function.
func CachedPerKeyT2[K comparable, T any, T2 any](duration time.Duration, cb GetElementFuncT2WithKey[K, T, T2], opts ...CacheOption) func(T, T2) *Element {
	node := &CachedNode{
		isByKey:  true,
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

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

// CachedPerKeyT3 caches the given element for the given duration. The element is only rendered once per key, and then cached for the given duration.
// The element is cached by the unique identifier that is returned by the callback function.
func CachedPerKeyT3[K comparable, T any, T2 any, T3 any](duration time.Duration, cb GetElementFuncT3WithKey[K, T, T2, T3], opts ...CacheOption) func(T, T2, T3) *Element {
	node := &CachedNode{
		isByKey:  true,
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

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

// CachedPerKeyT4 caches the given element for the given duration. The element is only rendered once per key, and then cached for the given duration.
// The element is cached by the unique identifier that is returned by the callback function.
func CachedPerKeyT4[K comparable, T any, T2 any, T3 any, T4 any](duration time.Duration, cb GetElementFuncT4WithKey[K, T, T2, T3, T4], opts ...CacheOption) func(T, T2, T3, T4) *Element {
	node := &CachedNode{
		isByKey:  true,
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

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

// CachedT caches the given element for the given duration. The element is only rendered once, and then cached for the given duration.
// Please note this element is globally cached, and not per unique identifier / user.
// Use CachedPerKey to cache elements per unique identifier.
func CachedT[T any](duration time.Duration, cb GetElementFuncT[T], opts ...CacheOption) func(T) *Element {
	node := &CachedNode{
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

	return func(data T) *Element {
		node.cb = func() *Element {
			return cb(data)
		}
		return element
	}
}

// CachedT2 caches the given element for the given duration. The element is only rendered once, and then cached for the given duration.
// Please note this element is globally cached, and not per unique identifier / user.
// Use CachedPerKey to cache elements per unique identifier.
func CachedT2[T any, T2 any](duration time.Duration, cb GetElementFuncT2[T, T2], opts ...CacheOption) func(T, T2) *Element {
	node := &CachedNode{
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

	return func(data T, data2 T2) *Element {
		node.cb = func() *Element {
			return cb(data, data2)
		}
		return element
	}
}

// CachedT3 caches the given element for the given duration. The element is only rendered once, and then cached for the given duration.
// Please note this element is globally cached, and not per unique identifier / user.
// Use CachedPerKey to cache elements per unique identifier.
func CachedT3[T any, T2 any, T3 any](duration time.Duration, cb GetElementFuncT3[T, T2, T3], opts ...CacheOption) func(T, T2, T3) *Element {
	node := &CachedNode{
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

	return func(data T, data2 T2, data3 T3) *Element {
		node.cb = func() *Element {
			return cb(data, data2, data3)
		}
		return element
	}
}

// CachedT4 caches the given element for the given duration. The element is only rendered once, and then cached for the given duration.
// Please note this element is globally cached, and not per unique identifier / user.
// Use CachedPerKey to cache elements per unique identifier.
func CachedT4[T any, T2 any, T3 any, T4 any](duration time.Duration, cb GetElementFuncT4[T, T2, T3, T4], opts ...CacheOption) func(T, T2, T3, T4) *Element {
	node := &CachedNode{
		duration: duration,
	}

	for _, opt := range opts {
		opt(node)
	}

	if node.cache == nil {
		node.cache = DefaultCacheProvider()
	}

	element := &Element{
		tag:  CachedNodeTag,
		meta: node,
	}

	return func(data T, data2 T2, data3 T3, data4 T4) *Element {
		node.cb = func() *Element {
			return cb(data, data2, data3, data4)
		}
		return element
	}
}

// ClearCache clears the cached HTML of the element. This is called automatically by the framework.
func (c *CachedNode) ClearCache() {
	c.cache.Purge()
}

// ClearExpired is deprecated and does nothing. Cache expiration is now handled by the Store implementation.
func (c *CachedNode) ClearExpired() {
	// No-op for backward compatibility
}

func (c *CachedNode) Render(ctx *RenderContext) {
	if c.isByKey {
		panic("CachedPerKey should not be rendered directly")
	} else {
		// For simple cached components, we use a single key
		// Use GetOrCompute for atomic check-and-set
		html := c.cache.GetOrCompute(_singleCacheKey, func() string {
			return Render(c.cb())
		}, c.duration)
		ctx.builder.WriteString(html)
	}
}

func (c *ByKeyEntry) Render(ctx *RenderContext) {
	key := c.key
	parentMeta := c.parent.meta.(*CachedNode)

	// Use GetOrCompute for atomic check-and-set
	html := parentMeta.cache.GetOrCompute(key, func() string {
		return Render(c.cb())
	}, parentMeta.duration)
	ctx.builder.WriteString(html)
}
