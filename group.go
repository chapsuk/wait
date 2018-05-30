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

// AddMany running call group.Add count times
func (g *Group) AddMany(count int, f func()) {
	for i := 0; i < count; i++ {
		g.Add(f)
	}
}

// AddManyWithContext running call group.AddWithContext count times
func (g *Group) AddManyWithContext(ctx context.Context, count int, f func(context.Context)) {
	for i := 0; i < count; i++ {
		g.AddWithContext(ctx, f)
	}
}

// Wait blocks until all added functions will be completed
func (g *Group) Wait() {
	g.wg.Wait()
}
