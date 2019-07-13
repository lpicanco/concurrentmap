# Concurrent Map
Go concurrent map implementation

[![GoDoc](https://godoc.org/github.com/lpicanco/concurrentmap?status.svg)](https://godoc.org/github.com/lpicanco/concurrentmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/lpicanco/concurrentmap)](https://goreportcard.com/report/github.com/lpicanco/concurrentmap)
[![GoCover](http://gocover.io/_badge/github.com/lpicanco/concurrentmap)](http://gocover.io/github.com/lpicanco/concurrentmap)

## How to use

Simple usage:

```go
import (
	"fmt"

	"github.com/lpicanco/concurrentmap"
)

func main() {
	m := concurrentmap.New()
	m.Put(42, "answer")

	value, found := m.Get(42)
	if found {
		fmt.Printf("Value: %v\n", value)
	}

	fmt.Printf("Map size: %v\n", m.Size())

	m.Remove(42)
}
```
