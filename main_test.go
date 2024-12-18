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

// TODO use this to track visited locations.
/*

type Visit struct {
	x int
	y int
}
type visited interface{}

func main() {

	var x visited
	fmt.Println("hello world")
	visitsMap := make(map[Visit]visited, 10)
	visitsMap[Visit{1, 2}] = x
	visitsMap[Visit{2, 3}] = nil
	visitsMap[Visit{2, 3}] = nil
	visitsMap[Visit{2, 3}] = nil
	visitsMap[Visit{7, 10}] = nil
	fmt.Println(len(visitsMap))
	fmt.Println(visitsMap)

	visits := []Visit{{1, 2}, {2, 3}, {7, 10}}
	if slices.Contains(visits, Visit{2, 3}) {
		fmt.Println("Y")
	}

}

*/