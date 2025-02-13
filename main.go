package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
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
	Visit struct {
		row int
		col int
	}
	Mark   interface{}
	ObsHit struct {
		row int
		col int
		dir Dir
	}
	Lab struct {
		obsRows map[int][]int
		obsCols map[int][]int
		marked  map[Visit]Mark
		obsHits []ObsHit
		newObs  int
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
	mark   Mark
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
	lab, guard := initialise(lines)
	// guard.walk finds obs in current dir, update row, col , path , dir, return finished as true or false
	for {
		finished := guard.walk(&lab)
		if finished {
			break
		}
	}

	return len(lab.marked)
}

func (guard *Guard) walk(lab *Lab) bool {
	if guard.dir == Up {
		// get obstacles in the upward column
		obs, ok := lab.obsCols[guard.col]
		// obstacles found
		if ok {
			// reverse iterate through the ordered row obstacle row vals for this col
			for i := len(obs) - 1; i >= 0; i-- {
				// if the row is below our current position, skip it
				if obs[i] >= guard.row {
					continue
				}
				// we've got the next obstacle above current pos, get the distance
				guard.path += guard.row - obs[i] - 1

				// mark the path elements
				for j := obs[i] + 1; j < guard.row; j++ {
					lab.marked[Visit{j, guard.col}] = mark
					// get 1st obstacle to the right in 1 below curr row
					checkNewObs(lab, j+1, guard)
				}

				// update obstacles hit, guards new row and new direction
				lab.obsHits = append(lab.obsHits, ObsHit{row: obs[i], col: guard.col, dir: Up})
				guard.row = obs[i] + 1
				guard.dir = Right

				// not finished
				return false
			}
			// no obstacles above our current position, mark all rows in current upwards to 0
			for i := 0; i < guard.row; i++ {
				lab.marked[Visit{i, guard.col}] = mark
			}
			// reached top exit lab
			return true
		}
		// no obstacles found, mark all rows in current upwards to 0
		for i := 0; i < guard.row; i++ {
			lab.marked[Visit{i, guard.col}] = mark
		}

		// reached top exit lab
		return true
	}

	if guard.dir == Right {
		obs, ok := lab.obsRows[guard.row]
		if ok {
			for i := 0; i < len(obs); i++ {
				if guard.col >= obs[i] {
					continue
				}

				guard.path += obs[i] - guard.col - 1

				for j := guard.col + 1; j < obs[i]; j++ {
					lab.marked[Visit{guard.row, j}] = mark
				}

				guard.col = obs[i] - 1
				guard.dir = Down

				return false
			}

			for i := guard.col + 1; i < maxCol; i++ {
				lab.marked[Visit{guard.row, i}] = mark
			}

			return true
		}

		for i := guard.col + 1; i < maxCol; i++ {
			lab.marked[Visit{guard.row, i}] = mark
		}

		return true
	}

	if guard.dir == Down {
		obs, ok := lab.obsCols[guard.col]
		if ok {
			for i := 0; i < len(obs); i++ {
				if guard.row >= obs[i] {
					continue
				}

				guard.path += obs[i] - guard.row - 1

				for j := guard.row + 1; j < obs[i]; j++ {
					lab.marked[Visit{j, guard.col}] = mark
				}

				guard.row = obs[i] - 1
				guard.dir = Left

				return false
			}

			for i := guard.row; i < maxRow; i++ {
				lab.marked[Visit{i, guard.col}] = mark
			}

			return true
		}

		for i := guard.row; i < maxRow; i++ {
			lab.marked[Visit{i, guard.col}] = mark
		}

		return true
	}

	if guard.dir == Left {
		obs, ok := lab.obsRows[guard.row]
		if ok {
			for i := len(obs) - 1; i >= 0; i-- {
				if obs[i] >= guard.col {
					continue
				}

				guard.path += guard.col - obs[i] - 1

				for j := obs[i] + 1; j < guard.col; j++ {
					lab.marked[Visit{guard.row, j}] = mark
				}

				guard.col = obs[i] + 1
				guard.dir = Up

				return false
			}

			for i := 0; i < guard.col; i++ {
				lab.marked[Visit{guard.row, i}] = mark
			}

			return true
		}

		for i := 0; i < guard.col; i++ {
			lab.marked[Visit{guard.row, i}] = mark
		}

		return true
	}

	return true
}

func checkNewObs(lab *Lab, currRow int, guard *Guard) {
	obsRow, ok := lab.obsRows[currRow]
	if ok {
		for i := 0; i < len(obsRow); i++ {
			if obsRow[i] < guard.col {
				continue // to next obstruction in this row
			}
			// check if hit already in dir Right
			if slices.Contains(lab.obsHits, ObsHit{row: currRow, col: obsRow[i], dir: Right}) {
				lab.newObs++
			}
		}
	}
}

func initialise(lines []string) (lab Lab, guard Guard) {
	// TODO replace named return params with var declarations
	maxRow = len(lines)
	maxCol = len(lines[0])
	lab.obsRows = make(map[int][]int, maxCol)
	lab.obsCols = make(map[int][]int, maxRow)
	lab.marked = make(map[Visit]Mark, maxCol*maxRow)
	lab.obsHits = []ObsHit{}

	for row, line := range lines {
		for col, cell := range line {
			// populate row and col obstacle slices with col & row vals in sequence
			// we get a slice of obstacles (col value) in each row in col sequence
			// and a slice of obstacles (row value) in each col in row sequence
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
				lab.marked[Visit{row, col}] = mark
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
