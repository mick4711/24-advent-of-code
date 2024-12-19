package main

import (
	"testing"
)

func TestGetPathLength(t *testing.T) {
	tests := []struct {
		file string
		want int
	}{
		{
			`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			41,
		},
	}

	for _, test := range tests {
		got := getPathLength(test.file)
		if got != test.want {
			t.Errorf("getPathLength(), got:%v, want:%v", got, test.want)
		}
	}
}
