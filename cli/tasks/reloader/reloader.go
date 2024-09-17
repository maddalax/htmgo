package reloader

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/maddalax/htmgo/cli/tasks/astgen"
	"github.com/maddalax/htmgo/cli/tasks/css"
	"github.com/maddalax/htmgo/cli/tasks/process"
	"github.com/maddalax/htmgo/cli/tasks/run"
	"log/slog"
	"strings"
	"sync"
)

type Change struct {
	name string
}

func NewChange(name string) *Change {
	return &Change{name: name}
}

func (c *Change) Name() string {
	return c.name
}

func (c *Change) HasAnyPrefix(prefix ...string) bool {
	for _, s := range prefix {
		if strings.HasPrefix(c.name, s) {
			return true
		}
	}
	return false
}

func (c *Change) HasAnySuffix(suffix ...string) bool {
	for _, s := range suffix {
		if strings.HasSuffix(c.name, s) {
			return true
		}
	}
	return false
}

func (c *Change) IsGenerated() bool {
	return c.HasAnySuffix("generated.go")
}

func (c *Change) IsGo() bool {
	return c.HasAnySuffix(".go")
}

type Tasks struct {
	AstGen bool
	Css    bool
	Run    bool
	Ent    bool
}

func OnFileChange(events []*fsnotify.Event) {
	tasks := Tasks{}

	for _, event := range events {
		c := NewChange(event.Name)

		slog.Debug("file changed", slog.String("file", c.Name()))

		if c.IsGenerated() {
			continue
		}

		if c.IsGo() && c.HasAnyPrefix("pages/", "partials/") {
			tasks.AstGen = true
		}

		if c.IsGo() {
			tasks.Css = true
			tasks.Run = true
		}

		if c.HasAnySuffix("tailwind.config.js", ".css") {
			tasks.Css = true
		}

		if c.HasAnyPrefix("ent/schema") {
			tasks.Ent = true
		}
	}

	deps := make([]func() any, 0)

	if tasks.AstGen {
		deps = append(deps, func() any {
			return astgen.GenAst(false)
		})
	}

	if tasks.Css {
		deps = append(deps, func() any {
			return css.GenerateCss(false)
		})
	}

	if tasks.Ent {
		deps = append(deps, func() any {
			run.EntGenerate()
			return nil
		})
	}

	wg := sync.WaitGroup{}

	for _, dep := range deps {
		wg.Add(1)
		go func(dep func() any) {
			defer wg.Done()
			err := dep()
			if err != nil {
				fmt.Println(err)
			}
		}(dep)
	}

	wg.Wait()

	if tasks.Run {
		process.KillAll()
		_ = run.Server(false)
	}
}
