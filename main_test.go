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
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},                      // clean
		{"xpqdo()c$mul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},              // preceding do()
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)xvdon't()aq£$undo()?mul(8,5))", 48},         // trailing don't
		{"xpqdo()c$mul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)xvdon't()aq£$undo()?mul(8,5))", 48}, // preceding do() and trailing don't
	}

	for _, test := range tests {
		got := getEnabledMuls(test.memory)
		if got != test.want {
			t.Errorf("getEnabledMuls(%v), got:%v, want:%v", test.memory, got, test.want)
		}
	}
}
