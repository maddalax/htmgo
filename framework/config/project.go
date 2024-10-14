package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type ProjectConfig struct {
	Tailwind    bool     `yaml:"tailwind"`
	WatchIgnore []string `yaml:"watch_ignore"`
	WatchFiles  []string `yaml:"watch_files"`
}

func DefaultProjectConfig() *ProjectConfig {
	return &ProjectConfig{
		Tailwind: true,
		WatchIgnore: []string{
			"node_modules", ".git", ".idea", "assets/dist",
		},
		WatchFiles: []string{
			"**/*.go", "**/*.html", "**/*.css", "**/*.js", "**/*.json", "**/*.yaml", "**/*.yml", "**/*.md",
		},
	}
}

func (cfg *ProjectConfig) EnhanceWithDefaults() *ProjectConfig {
	defaultCfg := DefaultProjectConfig()
	if len(cfg.WatchFiles) == 0 {
		cfg.WatchFiles = defaultCfg.WatchFiles
	}
	if len(cfg.WatchIgnore) == 0 {
		cfg.WatchIgnore = defaultCfg.WatchIgnore
	}
	return cfg
}

func FromConfigFile(workingDir string) *ProjectConfig {
	defaultCfg := DefaultProjectConfig()
	names := []string{"htmgo.yaml", "htmgo.yml", "_htmgo.yaml", "_htmgo.yml"}
	for _, name := range names {
		filePath := path.Join(workingDir, name)
		if _, err := os.Stat(filePath); err == nil {
			cfg := &ProjectConfig{}
			bytes, err := os.ReadFile(filePath)
			if err == nil {
				err = yaml.Unmarshal(bytes, cfg)
				if err == nil {
					return cfg.EnhanceWithDefaults()
				}
			}
		}
	}
	return defaultCfg
}
