package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
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
	cfg := FromConfigFile("non-existing-dir")
	assert.Equal(t, true, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}

func TestPartialConfigMerges(t *testing.T) {
	t.Parallel()
	dir := writeConfigFile(t, "tailwind: false")
	cfg := FromConfigFile(dir)
	assert.Equal(t, false, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}

func TestShouldNotSetTailwindTrue(t *testing.T) {
	t.Parallel()
	dir := writeConfigFile(t, "someValue: true")
	cfg := FromConfigFile(dir)
	assert.Equal(t, false, cfg.Tailwind)
	assert.Equal(t, 4, len(cfg.WatchIgnore))
	assert.Equal(t, 8, len(cfg.WatchFiles))
}

func writeConfigFile(t *testing.T, content string) string {
	temp := os.TempDir()
	os.Mkdir(temp, 0755)
	err := os.WriteFile(path.Join(temp, "htmgo.yml"), []byte(content), 0644)
	assert.Nil(t, err)
	return temp
}
