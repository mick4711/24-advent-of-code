package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	Dir   int
	Guard struct {
		row  int
		col  int
		path int
		dir  Dir
	}
	Lab struct {
		obsRows map[int][]int
		obsCols map[int][]int
	}
)

const (
	Up Dir = iota
	Right
	Down
	Left
)

var (
	maxCol int
	maxRow int
)

func main() {
	// set input file name
	inputFileName := getInputFileName()

	b, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	ans := getPathLength(string(b))
	fmt.Println("ans =", ans)
}

func getPathLength(file string) int {
	lines := strings.Split(file, "\n")

	// initialise obstructions and guard
	lab, guard := initialise(lines)
	// guard.walk finds obs in current dir, update row, col , path , dir, return leave as true or false
	for {
		finished := guard.walk(lab)
		if finished {
			break
		}
	}

	return guard.path
}

func (guard *Guard) walk(lab Lab) bool {
	if guard.dir == Up {
		obs, ok := lab.obsCols[guard.col]
		if ok {
			for i := len(obs) - 1; i >= 0; i-- {
				if obs[i] >= guard.row {
					continue
				}

				fmt.Println("hit obs:", obs[i], guard.col)
				guard.path += guard.row - obs[i] - 1
				guard.row = obs[i] + 1
				guard.dir = Right

				return false
			}

			return true
		}

		return false
	}

	if guard.dir == Right {
		obs, ok := lab.obsRows[guard.row]
		if ok {
			for i := 0; i < len(obs); i++ {
				if guard.col >= obs[i] {
					continue
				}

				fmt.Println("hit obs:", guard.row, obs[i])
				guard.path += obs[i] - guard.col - 1
				guard.col = obs[i] - 1
				guard.dir = Down

				return false
			}

			return true
		}

		return false
	}

	if guard.dir == Down {
		obs, ok := lab.obsCols[guard.col]
		if ok {
			for i := 0; i < len(obs); i++ {
				if guard.row >= obs[i] {
					continue
				}

				fmt.Println("hit obs:", obs[i], guard.col)
				guard.path += obs[i] - guard.row - 1
				guard.row = obs[i] - 1
				guard.dir = Left

				return false
			}

			return true
		}

		return false
	}

	if guard.dir == Left {
		obs, ok := lab.obsRows[guard.row]
		if ok {
			for i := len(obs) - 1; i >= 0; i-- {
				if obs[i] >= guard.col {
					continue
				}

				fmt.Println("hit obs:", guard.row, obs[i])
				guard.path += guard.col - obs[i] - 1
				guard.col = obs[i] + 1
				guard.dir = Up

				return false
			}

			return true
		}

		return false
	}

	return true
}

func initialise(lines []string) (lab Lab, guard Guard) {
	maxCol = len(lines)
	maxRow = len(lines[0])
	lab.obsRows = make(map[int][]int, maxCol)
	lab.obsCols = make(map[int][]int, maxRow)

	for row, line := range lines {
		for col, cell := range line {
			if string(cell) == "#" {
				_, ok := lab.obsRows[row]
				if !ok {
					lab.obsRows[row] = []int{col}
				} else {
					lab.obsRows[row] = append(lab.obsRows[row], col)
				}

				_, ok = lab.obsCols[col]
				if !ok {
					lab.obsCols[col] = []int{row}
				} else {
					lab.obsCols[col] = append(lab.obsCols[col], row)
				}
			}

			// find caret and starting direction up
			if string(cell) == "^" {
				guard = Guard{row: row, col: col, path: 1, dir: Up}
			}
		}
	}

	return lab, guard
}

func getInputFileName() string {
	var inputFile string

	testFlag := flag.Bool("test", true, "test flag")
	flag.Parse()
	log.Printf("Start - testFlag = %v\n", *testFlag)

	switch {
	case *testFlag:
		inputFile = "input_test.txt"
	default:
		inputFile = "input.txt"
	}

	return inputFile
}
