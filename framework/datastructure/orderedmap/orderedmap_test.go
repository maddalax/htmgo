package orderedmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	t.Parallel()
	om := New[string, int]()

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for index, letter := range alphabet {
		om.Set(letter, index)
	}

	assert.Equal(t, alphabet, om.Keys())

	c, ok := om.Get("c")
	assert.True(t, ok)
	assert.Equal(t, 2, c)

	for i, entry := range om.Entries() {
		if i == 5 {
			assert.Equal(t, "f", entry.Key)
		}
	}

	om.Delete("c")
	value, ok := om.Get("c")
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}
