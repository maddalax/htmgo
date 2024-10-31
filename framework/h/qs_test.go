package h

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
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

func TestGetQueryParam(t *testing.T) {
	t.Parallel()
	req, _ := http.NewRequest("GET", "http://localhost/?foo=bar&baz=qux", nil)
	ctx := &RequestContext{Request: req}

	result := GetQueryParam(ctx, "foo")
	assert.Equal(t, "bar", result)

	result = GetQueryParam(ctx, "baz")
	assert.Equal(t, "qux", result)

	result = GetQueryParam(ctx, "missing")
	assert.Equal(t, "", result)

	ctx.currentBrowserUrl = "http://localhost/?current=value"

	result = GetQueryParam(ctx, "current")
	assert.Equal(t, "value", result)

	// url params should override browser url
	req.URL, _ = url.Parse("http://localhost/?foo=override")
	result = GetQueryParam(ctx, "foo")
	assert.Equal(t, "override", result)
}
