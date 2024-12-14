package main

import (
	"testing"
)

func TestGetXmasCount(t *testing.T) {
	tests := []struct {
		file string
		want int
	}{
		{
			"MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX", 9,
		},
	}

	for _, test := range tests {
		got := getXmasCount(test.file)
		if got != test.want {
			t.Errorf("getXmasCount(%v), got:%v, want:%v", test.file, got, test.want)
		}
	}
}
