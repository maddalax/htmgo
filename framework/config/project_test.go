package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultProjectConfig(t *testing.T) {
	cfg := DefaultProjectConfig()
	assert.Equal(t, true, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}

func TestNoConfigFileUsesDefault(t *testing.T) {
	cfg := FromConfigFile("testdata")
	assert.Equal(t, true, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}

func TestPartialConfigMerges(t *testing.T) {
	os.Mkdir("testdata", 0755)
	defer os.RemoveAll("testdata")
	os.WriteFile("testdata/htmgo.yaml", []byte("tailwind: false"), 0644)
	cfg := FromConfigFile("testdata")
	assert.Equal(t, false, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}
