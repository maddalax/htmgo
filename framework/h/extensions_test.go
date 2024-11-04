package h

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseExtensions(t *testing.T) {
	// Test when not in development
	os.Unsetenv("ENV")
	result := BaseExtensions()
	expected := "path-deps, response-targets, mutation-error, htmgo, sse"
	assert.Equal(t, expected, result)

	// Test when in development
	os.Setenv("ENV", "development")
	result = BaseExtensions()
	expected = "path-deps, response-targets, mutation-error, htmgo, sse, livereload"
	assert.Equal(t, expected, result)
}
