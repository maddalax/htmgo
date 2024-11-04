package service

import (
	"log"
	"reflect"
	"sync"
)

type Lifecycle = string

var (
	Singleton Lifecycle = "singleton"
	Transient Lifecycle = "transient"
)

type Provider struct {
	cb        func() any
	lifecycle Lifecycle
}

type Locator struct {
	services map[string]Provider
	cache    map[string]any
	mutex    sync.RWMutex
}

// NewLocator creates a new locator to register services
// Usage:
//
// locator := service.NewLocator()
// service.Set[db.Queries](locator, service.Singleton, db.Provide)
//
// service.Get[db.Queries](locator)
func NewLocator() *Locator {
	return &Locator{
		services: make(map[string]Provider),
		cache:    make(map[string]any),
		mutex:    sync.RWMutex{},
	}
}

func (l *Locator) setCache(key string, value any) {
	l.cache[key] = value
}

func (l *Locator) clearCache(key string) {
	delete(l.cache, key)
}

func (l *Locator) getCache(key string) any {
	return l.cache[key]
}

// Get returns a service from the locator
// If the service is not found, log.Fatalf is called
// If the service is a singleton, it will be cached after first invocation
func Get[T any](locator *Locator) *T {
	locator.mutex.RLock()
	i := new(T)
	t := reflect.TypeOf(i).String()

	cached := locator.getCache(t)

	if cached != nil {
		locator.mutex.RUnlock()
		return cached.(*T)
	}

	entry, ok := locator.services[t]
	if !ok {
		log.Fatalf("%s is not registered in the service locator", t)
	}
	cb := entry.cb().(*T)
	locator.mutex.RUnlock()
	locator.mutex.Lock()
	if entry.lifecycle == Singleton {
		locator.setCache(t, cb)
	}
	locator.mutex.Unlock()
	return cb
}

// Set registers a service with the locator
// If the service is a singleton, it will be cached after first invocation of Get
func Set[T any](locator *Locator, lifecycle Lifecycle, value func() *T) {
	t := reflect.TypeOf(value)
	rt := t.Out(0)
	key := rt.String()
	locator.services[key] = Provider{
		cb: func() any {
			return value()
		},
		lifecycle: lifecycle,
	}
	locator.clearCache(key)
}
