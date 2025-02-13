package main

import (
	"maps"
	"slices"
	"testing"
)

func TestGetPathLength(t *testing.T) {
	tests := []struct {
		name       string
		file       string
		wantLength int
		wantNewObs int
	}{
		{
			"example",
			`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`, 41, 6,
		},
		{
			"1 turn",
			`#....
.....
.....
.....
^....`, 8, 0,
		},
		{
			"2 turns",
			`#....
....#
.....
.....
^....`, 10, 0,
		},
		{
			"3 turns",
			`#....
....#
.....
.....
^..#.`, 11, 0,
		},
		{
			"4 turns",
			`#....
....#
.....
.#...
^..#.`, 12, 0,
		},
		{
			"multiple obstacles",
			`..........
....#.....
.#......#.
..........
#.........
.#.....##.
.........#
..........
.^.#......
........#.`, 38, 0,
		},
		{
			"bounce back",
			`#....
.#...
.....
.....
^....`, 4, 0,
		},
		{
			"bounce back after turn",
			`#....
....#
...#.
.....
^....`, 7, 0,
		},
	}

	for _, test := range tests {
		gotLength, gotNewObs := getPathLength(test.file)
		if gotLength != test.wantLength {
			t.Errorf("getPathLength() length, %v got:%v, want:%v", test.name, gotLength, test.wantLength)
		}

		if gotNewObs != test.wantNewObs {
			t.Errorf("getPathLength() new obs, %v got:%v, want:%v", test.name, gotNewObs, test.wantNewObs)
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
		test.guard.walk(&test.lab)

		if !maps.Equal(test.lab.marked, test.want.marked) {
			t.Errorf("guard.walk, got:%v, want:%v", test.lab.marked, test.want.marked)
		}

		if test.guard != test.want.guard {
			t.Errorf("guard.walk, got:%v, want:%v", test.guard, test.want.guard)
		}
	}
}

func TestNewObstacle(t *testing.T) {
	tests := []struct {
		direction string
		guard     Guard
		lab       Lab
		want      Lab
	}{
		{
			direction: "Up",
			guard:     Guard{row: 6, col: 0, path: 1, dir: Up},
			lab: Lab{
				obsRows: map[int][]int{0: {0}, 3: {4}},
				obsCols: map[int][]int{0: {0}, 4: {3}},
				marked:  map[Visit]Mark{{6, 0}: mark},
				obsHits: []ObsHit{{row: 3, col: 4, dir: Right}},
			},
			want: Lab{
				obsHits: []ObsHit{{row: 3, col: 4, dir: Right}, {row: 0, col: 0, dir: Up}},
				newObs:  1,
			},
		},
		{
			direction: "Right",
			guard:     Guard{row: 3, col: 0, path: 1, dir: Right},
			lab: Lab{
				obsRows: map[int][]int{6: {2}, 3: {4}},
				obsCols: map[int][]int{2: {6}, 4: {3}},
				marked:  map[Visit]Mark{{3, 0}: mark},
				obsHits: []ObsHit{{row: 6, col: 2, dir: Down}},
			},
			want: Lab{
				obsHits: []ObsHit{{row: 6, col: 2, dir: Down}, {row: 3, col: 4, dir: Right}},
				newObs:  1,
			},
		},
		{
			direction: "Down",
			guard:     Guard{row: 0, col: 4, path: 1, dir: Down},
			lab: Lab{
				obsRows: map[int][]int{6: {4}, 4: {1}},
				obsCols: map[int][]int{1: {4}, 4: {6}},
				marked:  map[Visit]Mark{{0, 4}: mark},
				obsHits: []ObsHit{{row: 4, col: 1, dir: Left}},
			},
			want: Lab{
				obsHits: []ObsHit{{row: 4, col: 1, dir: Left}, {row: 6, col: 4, dir: Down}},
				newObs:  1,
			},
		},
		{
			direction: "Left",
			guard:     Guard{row: 6, col: 4, path: 1, dir: Left},
			lab: Lab{
				obsRows: map[int][]int{6: {0}, 0: {2}},
				obsCols: map[int][]int{0: {6}, 2: {0}},
				marked:  map[Visit]Mark{{6, 4}: mark},
				obsHits: []ObsHit{{row: 0, col: 2, dir: Up}},
			},
			want: Lab{
				obsHits: []ObsHit{{row: 0, col: 2, dir: Up}, {row: 6, col: 0, dir: Left}},
				newObs:  1,
			},
		},
	}
	for _, test := range tests {
		test.guard.walk(&test.lab)

		if test.lab.newObs != test.want.newObs {
			t.Errorf("new obstruction found, dir = %v: got:%v, want:%v", test.direction, test.lab.newObs, test.want.newObs)
		}

		if !slices.Contains(test.lab.obsHits, test.want.obsHits[1]) {
			t.Errorf("new obstruction hit added, dir = %v: got:%v, want:%v", test.direction, test.lab.obsHits, test.want.obsHits)
		}
	}
}
