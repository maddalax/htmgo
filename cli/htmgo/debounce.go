package htmgo

import (
	"sync"
	"time"
)

// Debouncer is a struct that holds the debounce logic
type Debouncer struct {
	delay time.Duration
	timer *time.Timer
	mu    sync.Mutex
}

// NewDebouncer creates a new Debouncer with the specified delay
func NewDebouncer(delay time.Duration) *Debouncer {
	return &Debouncer{
		delay: delay,
	}
}

// Do calls the provided function after the delay, resetting the delay if called again
func (d *Debouncer) Do(f func()) {
	d.mu.Lock()
	defer d.mu.Unlock()

	// If there's an existing timer, stop it
	if d.timer != nil {
		d.timer.Stop()
	}

	// Create a new timer
	d.timer = time.AfterFunc(d.delay, f)
}
