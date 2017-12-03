package wait

import (
	"context"
	"sync"
)

// Group is wrapper over sync.WaitGroup
type Group struct {
	wg sync.WaitGroup
}

// Add function calling argument function in a separate goroutine with sync.WaitGroup control
func (g *Group) Add(f func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f()
	}()
}

// AddWithContext function calling argument function with context
// in a separate goroutine with sync.WaitGroup control
func (g *Group) AddWithContext(ctx context.Context, f func(context.Context)) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f(ctx)
	}()
}

// Wait blocks until all added functions will be completed
func (g *Group) Wait() {
	g.wg.Wait()
}
