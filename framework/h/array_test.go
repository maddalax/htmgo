package h

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUnique(t *testing.T) {
	t.Parallel()
	slice := []string{"a", "b", "b", "c", "d", "d", "x"}
	unique := Unique(slice, func(item string) string {
		return item
	})
	assert.Equal(t, []string{"a", "b", "c", "d", "x"}, unique)
}

func TestFilter(t *testing.T) {
	t.Parallel()
	slice := []string{"a", "b", "b", "c", "d", "d", "x"}
	filtered := Filter(slice, func(item string) bool {
		return item == "b"
	})
	assert.Equal(t, []string{"b", "b"}, filtered)
}

func TestMap(t *testing.T) {
	t.Parallel()
	slice := []string{"a", "b", "c"}
	mapped := Map(slice, func(item string) string {
		return strings.ToUpper(item)
	})
	assert.Equal(t, []string{"A", "B", "C"}, mapped)
}

func TestGroupBy(t *testing.T) {
	t.Parallel()

	type Item struct {
		Name string
		Job  string
	}

	items := []Item{
		{Name: "Alice", Job: "Developer"},
		{Name: "Bob", Job: "Designer"},
		{Name: "Charlie", Job: "Developer"},
		{Name: "David", Job: "Designer"},
		{Name: "Eve", Job: "Developer"},
		{Name: "Frank", Job: "Product Manager"},
	}

	grouped := GroupBy(items, func(item Item) string {
		return item.Job
	})

	assert.Equal(t, 3, len(grouped))
	assert.Equal(t, 3, len(grouped["Developer"]))
	assert.Equal(t, 2, len(grouped["Designer"]))
	assert.Equal(t, 1, len(grouped["Product Manager"]))
}

func TestGroupByOrdered(t *testing.T) {
	t.Parallel()

	type Item struct {
		Name string
		Job  string
	}

	items := []Item{
		{Name: "Alice", Job: "Developer"},
		{Name: "Bob", Job: "Designer"},
		{Name: "Charlie", Job: "Developer"},
		{Name: "David", Job: "Designer"},
		{Name: "Eve", Job: "Developer"},
		{Name: "Frank", Job: "Product Manager"},
	}

	grouped := GroupByOrdered(items, func(item Item) string {
		return item.Job
	})

	keys := []string{"Developer", "Designer", "Product Manager"}
	assert.Equal(t, keys, grouped.Keys())

	devs, ok := grouped.Get("Developer")
	assert.True(t, ok)
	assert.Equal(t, 3, len(devs))
	assert.Equal(t, "Alice", devs[0].Name)
	assert.Equal(t, "Charlie", devs[1].Name)
	assert.Equal(t, "Eve", devs[2].Name)
}

func TestFind(t *testing.T) {
	t.Parallel()
	slice := []string{"a", "b", "c"}
	found := Find(slice, func(item *string) bool {
		return *item == "b"
	})
	assert.Equal(t, "b", *found)
}
