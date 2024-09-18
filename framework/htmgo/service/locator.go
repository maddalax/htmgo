package service

import (
	"log"
	"reflect"
	"sync"
)

type Lifecycle = string

var (
	Singleton Lifecycle = "singleton"
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

func (l *Locator) getCache(key string) any {
	return l.cache[key]
}

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

func Set[T any](locator *Locator, lifecycle Lifecycle, value func() *T) {
	t := reflect.TypeOf(value)
	rt := t.Out(0)
	locator.services[rt.String()] = Provider{
		cb: func() any {
			return value()
		},
		lifecycle: lifecycle,
	}
}
