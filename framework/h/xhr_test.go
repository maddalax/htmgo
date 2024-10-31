package h

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	attr := Div(Get("/path", "load"))
	assert.Equal(t, `<div hx-get="/path" hx-trigger="load"></div>`, Render(attr))
}

func TestGetPartial(t *testing.T) {
	partial := func(ctx *RequestContext) *Partial {
		return &Partial{}
	}
	attr := Div(GetPartial(partial, "load"))
	expected := Render(Div(Get(GetPartialPath(partial), "load")))
	assert.Equal(t, expected, Render(attr))
}

func TestGetPartialWithQs(t *testing.T) {
	partial := func(ctx *RequestContext) *Partial {
		return &Partial{}
	}
	qs := NewQs("param", "value")
	attr := Div(GetPartialWithQs(partial, qs, "load"))
	expectedPath := Render(Div(Get(GetPartialPathWithQs(partial, qs), "load")))
	assert.Equal(t, expectedPath, Render(attr))
}

func TestPost(t *testing.T) {
	attr := Div(Post("/path", "submit"))
	assert.Equal(t, `<div hx-post="/path" hx-trigger="submit"></div>`, Render(attr))
}

func TestPostOnClick(t *testing.T) {
	attr := Div(PostOnClick("/path"))
	assert.Equal(t, `<div hx-post="/path" hx-trigger="click"></div>`, Render(attr))
}

func TestPostPartialOnClick(t *testing.T) {
	partial := func(ctx *RequestContext) *Partial {
		return &Partial{}
	}
	attr := Div(PostPartialOnClick(partial))
	expected := Render(Div(PostOnClick(GetPartialPath(partial))))
	assert.Equal(t, expected, Render(attr))
}

func TestPostPartialWithQs(t *testing.T) {
	partial := func(ctx *RequestContext) *Partial {
		return &Partial{}
	}
	qs := NewQs("key", "value")
	attr := Div(PostPartialWithQs(partial, qs, "click"))
	expected := Render(Div(Post(GetPartialPathWithQs(partial, qs), "click")))
	assert.Equal(t, expected, Render(attr))
}

func TestPostPartialOnClickQs(t *testing.T) {
	partial := func(ctx *RequestContext) *Partial {
		return &Partial{}
	}
	qs := NewQs("key", "value")
	attr := Div(PostPartialOnClickQs(partial, qs))
	expected := Render(Div(PostOnClick(GetPartialPathWithQs(partial, qs))))
	assert.Equal(t, expected, Render(attr))
}

func TestGetWithQs(t *testing.T) {
	qs := NewQs("param1", "value1", "param2", "value2")
	attr := Div(GetWithQs("/path", qs, "load"))
	expected := `<div hx-get="/path?param1=value1&amp;param2=value2" hx-trigger="load"></div>`
	assert.Equal(t, expected, Render(attr))
}

func TestPostWithQs(t *testing.T) {
	qs := NewQs("param1", "value1", "param2", "value2")
	attr := Div(PostWithQs("/path", qs, "submit"))
	expected := `<div hx-post="/path?param1=value1&amp;param2=value2" hx-trigger="submit"></div>`
	assert.Equal(t, expected, Render(attr))
}

func TestPostPartial(t *testing.T) {
	partial := func(ctx *RequestContext) *Partial {
		return &Partial{}
	}
	attr := Div(PostPartial(partial, "submit"))
	expected := Render(Div(Post(GetPartialPath(partial), "submit")))
	assert.Equal(t, expected, Render(attr))
}
