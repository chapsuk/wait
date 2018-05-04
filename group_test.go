package wait

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
)

func do()                               { fmt.Print("do\n") }
func doWithArgs(i, j int)               { fmt.Printf("doWith args: %d %d\n", i, j) }
func doWithContext(ctx context.Context) { fmt.Printf("doWithContext\n") }

func TestExamplesFromReadme(t *testing.T) {
	wg := Group{}

	wg.Add(do)
	wg.Add(func() {
		doWithArgs(1, 2)
	})
	wg.AddWithContext(context.TODO(), doWithContext)
	wg.Wait()
}

func TestAddWithMultipleFuncs(t *testing.T) {
	value1 := int32(0)
	value2 := int32(0)
	value3 := int32(0)

	wg := Group{}

	wg.Add(func() {
		atomic.StoreInt32(&value1, 1)
	})
	wg.Add(func() {
		atomic.StoreInt32(&value2, 1)
	})
	wg.Add(func() {
		atomic.StoreInt32(&value3, 1)
	})

	wg.Wait()

	if value1 != 1 || value2 != 1 || value3 != 1 {
		t.Error()
	}
}

func TestAddWithContextFuncWaitForDone(t *testing.T) {
	wg := Group{}

	ctx, cancel := context.WithCancel(context.Background())
	wg.AddWithContext(ctx, func(ctx context.Context) {
		<-ctx.Done()
	})

	go cancel()
	wg.Wait()
}
