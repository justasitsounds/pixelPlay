package main

import (
	"fmt"
	"testing"
)

func TestSmoothing(t *testing.T) {
	limit := 4
	for i := 0; i <= limit; i++ {
		fmt.Printf("smoothing(%v, %v) = %v\n", i, limit, smoothing(i, limit))
	}
	t.FailNow()
}
