package routine

import (
	"fmt"
	"time"
)

func DebugLongRunning(name string, f func()) {
	now := time.Now()
	done := make(chan struct{}, 1)
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				elapsed := time.Since(now).Milliseconds()
				fmt.Printf("function %s has not finished after %dms\n", name, elapsed)
			}
		}
	}()
	f()
	done <- struct{}{}
}
