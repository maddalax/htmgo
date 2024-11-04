package h

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerialize(t *testing.T) {
	t.Parallel()
	data := map[string]any{
		"hello": "world",
		"foo":   "bar",
	}
	serialized := JsonSerializeOrEmpty(data)
	assert.Equal(t, `{"foo":"bar","hello":"world"}`, serialized)
}

func TestSerializeNil(t *testing.T) {
	t.Parallel()
	serialized := JsonSerializeOrEmpty(nil)
	assert.Equal(t, "", serialized)
}

func TestSerializeInvalid(t *testing.T) {
	t.Parallel()
	serialized := JsonSerializeOrEmpty(func() {})
	assert.Equal(t, "", serialized)
}
