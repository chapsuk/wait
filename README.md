# What

Sugared `sync.WaitGroup`

## Why

LGTM

## Example

```go
package main

import (
    "context"
    "fmt"

    "github.com/chapsuk/wait"
)

func do()                               { fmt.Print("do\n") }
func doWithArgs(i, j int)               { fmt.Printf("doWith args: %d %d\n", i, j) }
func doWithContext(ctx context.Context) { fmt.Printf("doWithContext\n") }

func main() {
    wg := wait.Group{}

    wg.Add(do)
    wg.Add(func() {
        doWithArgs(1, 2)
    })
    wg.AddWithContext(context.TODO(), doWithContext)
    wg.Wait()
}

```
