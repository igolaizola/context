[![Go Report Card](https://goreportcard.com/badge/github.com/igolaizola/context)](https://goreportcard.com/report/github.com/igolaizola/context)
[![GoDoc](https://godoc.org/github.com/igolaizola/context/go?status.svg)](https://godoc.org/github.com/igolaizola/context)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

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
