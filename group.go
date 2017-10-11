package wait

import (
	"context"
	"sync"
)

type Group struct {
	wg sync.WaitGroup
}

func (g *Group) Add(f func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f()
	}()
}

func (g *Group) AddWithContext(ctx context.Context, f func(context.Context)) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f(ctx)
	}()
}

func (g *Group) Wait() {
	g.wg.Wait()
}
