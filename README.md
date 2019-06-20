# context: updatable deadline
golang context.Context implementation with an updatable deadline

## usage

```go
package main

import (
	"context"
	"time"
	
	igoctx "github.com/igolaizola/context"
)

func main() {
    ctx, cancel := igoctx.WithDeadline(context.Background())
    defer cancel()
    if err := ctx.SetDeadline(time.Now().Add(1 * time.Second)); err != nil {
    	panic(err)
    }
    <-ctx.Done()
}
```
