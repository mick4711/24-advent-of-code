package main

import (
	"testing"
)

func TestParseNums(t *testing.T) {
	tests := []struct {
		line      string
		left      []int
		right     []int
		wantLeft  []int
		wantRight []int
	}{
		{"3    4", []int{}, []int{}, []int{3}, []int{4}},
		{"3    4", []int{1}, []int{2}, []int{1, 3}, []int{2, 4}},
	}

	for _, test := range tests {
		gotLeft, gotRight := parseNums(test.line, test.left, test.right)
		for i := range gotLeft {
			if gotLeft[i] != test.wantLeft[i] || gotRight[i] != test.wantRight[i] {
				t.Errorf("parseNums(%v, %v, %v), got:%v, %v, want:%v,%v", test.line, test.left, test.right,
					gotLeft[i], gotRight[i], test.wantLeft[i], test.wantRight[i])
			}
		}
	}
}

func TestGetListDistance(t *testing.T) {
	tests := []struct {
		left  []int
		right []int
		want  int
	}{
		{[]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 11},
	}

	for _, test := range tests {
		got := getListDistance(test.left, test.right)
		if got != test.want {
			t.Errorf("getListDistance(%v, %v), got:%v, want:%v", test.left, test.right, got, test.want)
		}
	}
}
