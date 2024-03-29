package wait

import (
	"context"
	"sync"
	"time"
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

// WaitTimeout waits works group and return error by timeout if some goroutine dont finished
func (g *Group) WaitTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{})
	go func() {
		g.wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}
