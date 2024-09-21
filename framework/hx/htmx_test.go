package hx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStringTrigger(t *testing.T) {
	trigger := "click once, htmx:click throttle:5, load delay:10"
	tgr := NewStringTrigger(trigger)
	assert.Equal(t, len(tgr.events), 3)
	assert.Equal(t, tgr.events[0].event, "click")
	assert.Equal(t, tgr.events[0].modifiers[0].Modifier(), "once")
	assert.Equal(t, tgr.events[1].event, "click")
	assert.Equal(t, tgr.events[1].modifiers[0].Modifier(), "throttle:5")
	assert.Equal(t, tgr.events[2].event, "load")
	assert.Equal(t, tgr.events[2].modifiers[0].Modifier(), "delay:10")
	assert.Equal(t, "click once, click throttle:5, load delay:10", tgr.ToString())
}
