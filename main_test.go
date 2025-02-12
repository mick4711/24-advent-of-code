package main

import (
	"maps"
	"testing"
)

func TestGetPathLength(t *testing.T) {
	tests := []struct {
		file string
		want int
	}{
		{ // example
			`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`, 41,
		},
		{ // 1 turn
			`#....
.....
.....
.....
^....`, 8,
		},
		{ // 2 turns
			`#....
....#
.....
.....
^....`, 10},
		{ // 3 turns
			`#....
....#
.....
.....
^..#.`, 11},
		{ // 4 turns
			`#....
....#
.....
.#...
^..#.`, 12},
		{ // multiple obstacles
			`..........
....#.....
.#......#.
..........
#.........
.#.....##.
.........#
..........
.^.#......
........#.`, 38},
		{ // bounce back
			`#....
.#...
.....
.....
^....`, 4},
		{ // bounce back after turn
			`#....
....#
...#.
.....
^....`, 7},
	}

	for _, test := range tests {
		got := getPathLength(test.file)
		if got != test.want {
			t.Errorf("getPathLength(), got:%v, want:%v", got, test.want)
		}
	}
}

type Wanted struct {
	marked map[Visit]Mark
	guard  Guard
}

func TestGuardWalk(t *testing.T) {
	tests := []struct {
		guard Guard
		lab   Lab
		want  Wanted
	}{
		{
			guard: Guard{row: 4, col: 0, path: 1, dir: Up},
			lab: Lab{
				obsRows: map[int][]int{0: {0}},
				obsCols: map[int][]int{0: {0}},
				marked:  map[Visit]Mark{{4, 0}: mark},
			},
			want: Wanted{
				marked: map[Visit]Mark{{1, 0}: mark, {2, 0}: mark, {3, 0}: mark, {4, 0}: mark},
				guard:  Guard{row: 1, col: 0, path: 4, dir: Right},
			},
		},
		{
			guard: Guard{row: 0, col: 0, path: 1, dir: Right},
			lab: Lab{
				obsRows: map[int][]int{0: {4}},
				obsCols: map[int][]int{4: {0}},
				marked:  map[Visit]Mark{{0, 0}: mark},
			},
			want: Wanted{
				marked: map[Visit]Mark{{0, 0}: mark, {0, 1}: mark, {0, 2}: mark, {0, 3}: mark},
				guard:  Guard{row: 0, col: 3, path: 4, dir: Down},
			},
		},
		{
			guard: Guard{row: 0, col: 4, path: 1, dir: Down},
			lab: Lab{
				obsRows: map[int][]int{4: {4}},
				obsCols: map[int][]int{4: {4}},
				marked:  map[Visit]Mark{{0, 4}: mark},
			},
			want: Wanted{
				marked: map[Visit]Mark{{0, 4}: mark, {1, 4}: mark, {2, 4}: mark, {3, 4}: mark},
				guard:  Guard{row: 3, col: 4, path: 4, dir: Left},
			},
		},
		{
			guard: Guard{row: 4, col: 4, path: 1, dir: Left},
			lab: Lab{
				obsRows: map[int][]int{4: {0}},
				obsCols: map[int][]int{0: {4}},
				marked:  map[Visit]Mark{{4, 4}: mark},
			},
			want: Wanted{
				marked: map[Visit]Mark{{4, 1}: mark, {4, 2}: mark, {4, 3}: mark, {4, 4}: mark},
				guard:  Guard{row: 4, col: 1, path: 4, dir: Up},
			},
		},
	}

	for _, test := range tests {
		test.guard.walk(test.lab)

		if !maps.Equal(test.lab.marked, test.want.marked) {
			t.Errorf("guard.walk, got:%v, want:%v", test.lab.marked, test.want.marked)
		}

		if test.guard != test.want.guard {
			t.Errorf("guard.walk, got:%v, want:%v", test.guard, test.want.guard)
		}
	}
}
