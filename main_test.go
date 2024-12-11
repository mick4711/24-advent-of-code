package main

import (
	"testing"
)

// TODO add tests for nested dos and donts
func TestIsSafeReport(t *testing.T) {
	tests := []struct {
		memory string
		want   int
	}{
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},
	}

	for _, test := range tests {
		got := getMulResult(test.memory)
		if got != test.want {
			t.Errorf("getMulResult(%v), got:%v, want:%v", test.memory, got, test.want)
		}
	}
}
