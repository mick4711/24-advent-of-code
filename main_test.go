package main

import (
	"testing"
)

func TestIsSafeReport(t *testing.T) {
	tests := []struct {
		memory string
		want   int
	}{
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161},
	}

	for _, test := range tests {
		got := getMulResult(test.memory)
		if got != test.want {
			t.Errorf("getMulResult(%v), got:%v, want:%v", test.memory, got, test.want)
		}
	}
}
