# What

Sugared `sync.WaitGroup`

## Why

LGTM

## Example

```go
import "github.com/chapsuk/wait"

wg := wait.Group{}

wg.Add(do.Func)
wg.Add(func() {
    do.FuncWithArgs(arg1, arg2)
})

wg.Wait()
```
