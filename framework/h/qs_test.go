package h

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertHas(t *testing.T, qs *Qs, key string, value string) {
	str := qs.ToString()
	if value == "" {
		assert.Contains(t, str, key)
		assert.NotContains(t, str, key+"=")
	} else {
		assert.Contains(t, str, key+"="+value)
	}
}

func TestQs(t *testing.T) {
	t.Parallel()
	qs := NewQs("a", "b", "c")
	assertHas(t, qs, "a", "b")
	assertHas(t, qs, "c", "")

	qs2 := NewQs("a", "b", "c", "d")
	assertHas(t, qs2, "a", "b")
	assertHas(t, qs2, "c", "d")

	qs2.Add("e", "f")
	assertHas(t, qs2, "a", "b")
	assertHas(t, qs2, "c", "d")
	assertHas(t, qs2, "e", "f")

	qs2.Remove("e")
	assert.NotContains(t, qs2.ToString(), "e")
}

func TestSetQsOnUrl(t *testing.T) {
	t.Parallel()
	qs := NewQs("a", "b", "c", "d")
	set := SetQueryParams("https://example.com/path", qs)
	assert.Equal(t, "https://example.com/path?a=b&c=d", set)
}

func TestSetQsOnUrlWithDelete(t *testing.T) {
	t.Parallel()
	qs := NewQs("a", "b2", "c", "")
	set := SetQueryParams("https://example.com/path?a=b&c=d", qs)
	assert.Equal(t, "https://example.com/path?a=b2", set)
}
