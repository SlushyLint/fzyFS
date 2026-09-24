package main

import "testing"

func TestRunWithoutQueryReturnsUsage(t *testing.T) {
	if got := run([]string{}); got != 1 {
		t.Fatalf("run([]string{}) = %d, want 1", got)
	}
}
