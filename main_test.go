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
	if smoothing(0, limit) != 1.0 {
		t.Fatalf("expected %v but got %v", 1.0, smoothing(0, limit))
	}
	if smoothing(limit, limit) != 0 {
		t.Fatalf("expected %v but got %v", 0, smoothing(limit, limit))
	}
}
