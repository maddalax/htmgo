package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultProjectConfig(t *testing.T) {
	t.Parallel()
	cfg := DefaultProjectConfig()
	assert.Equal(t, true, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}

func TestNoConfigFileUsesDefault(t *testing.T) {
	t.Parallel()
	cfg := FromConfigFile("testdata2")
	assert.Equal(t, true, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}

func TestPartialConfigMerges(t *testing.T) {
	t.Parallel()
	os.Mkdir("testdata", 0755)
	defer os.RemoveAll("testdata")
	os.WriteFile("testdata/htmgo.yaml", []byte("tailwind: false"), 0644)
	cfg := FromConfigFile("testdata")
	assert.Equal(t, false, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}

func TestShouldNotSetTailwindTrue(t *testing.T) {
	t.Parallel()
	os.Mkdir("testdata1", 0755)
	defer os.RemoveAll("testdata1")
	os.WriteFile("testdata1/htmgo.yaml", []byte("someValue: false"), 0644)
	cfg := FromConfigFile("testdata")
	assert.Equal(t, false, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}
