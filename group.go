package wait

import "sync"

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

func (g *Group) Wait() {
	g.wg.Wait()
}
