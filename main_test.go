package main

import (
	"testing"
)

func TestIsSafeReport(t *testing.T) {
	tests := []struct {
		report string
		want   bool
	}{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", false},
		{"8 6 4 4 1", false},
		{"1 3 6 7 9", true},
	}

	for _, test := range tests {
		got := isSafeReport(test.report)
		if got != test.want {
			t.Errorf("isSafeReport(%v), got:%v, want:%v", test.report, got, test.want)
		}
	}
}
