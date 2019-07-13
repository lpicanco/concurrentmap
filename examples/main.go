package main

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
