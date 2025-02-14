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
		start   Visit
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

	length, newObs := getPathLength(string(b))
	fmt.Println("length =", length, "new obs =", newObs)
}

func getPathLength(file string) (length, newObs int) {
	lines := strings.Split(file, "\n")
	lab, guard := initialise(lines)
	// guard.walk finds obs in current dir, update row, col , path , dir, return finished as true or false
	for {
		finished := guard.walk(&lab)
		if finished {
			break
		}
	}

	return len(lab.marked), lab.newObs
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
					checkNewObsRight(lab, j+1, guard)
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
			// get 1st obstacle to the right in 1 below curr row
			checkNewObsRight(lab, i+1, guard)
		}
		// reached top exit lab
		return true
	}

	if guard.dir == Right {
		obs, ok := lab.obsRows[guard.row]
		// obstacles found
		if ok {
			// iterate through the ordered col obstacle col vals for this row
			for i := 0; i < len(obs); i++ {
				// if the row is left of our current position, skip it
				if guard.col >= obs[i] {
					continue
				}
				// we've got the next obstacle right of current pos, get the distance
				guard.path += obs[i] - guard.col - 1

				// mark the path elements
				for j := guard.col + 1; j < obs[i]; j++ {
					lab.marked[Visit{guard.row, j}] = mark
					// get 1st obstacle below in 1 to left of curr col
					checkNewObsDown(lab, j-1, guard)
				}

				// update obstacles hit, guards new row and new direction
				lab.obsHits = append(lab.obsHits, ObsHit{row: guard.row, col: obs[i], dir: Right})
				guard.col = obs[i] - 1
				guard.dir = Down

				// not finished
				return false
			}
			// no obstacles right of our current position, mark all cols in current right to end
			for i := guard.col + 1; i < maxCol; i++ {
				lab.marked[Visit{guard.row, i}] = mark
				// get 1st obstacle below in 1 to left of curr col
				checkNewObsDown(lab, i-1, guard)
			}
			// reached right edge exit lab
			return true
		}
		// no obstacles found, mark all cols in current right to end
		for i := guard.col + 1; i < maxCol; i++ {
			lab.marked[Visit{guard.row, i}] = mark
		}
		// reached right edge exit lab
		return true
	}

	if guard.dir == Down {
		// get obstacles in the downward column
		obs, ok := lab.obsCols[guard.col]
		// obstacles found
		if ok {
			// iterate through the ordered row obstacle row vals for this col
			for i := 0; i < len(obs); i++ {
				// if the row is above our current position, skip it
				if guard.row >= obs[i] {
					continue
				}
				// we've got the next obstacle below current pos, get the distance
				guard.path += obs[i] - guard.row - 1

				// mark the path elements
				for j := guard.row + 1; j < obs[i]; j++ {
					lab.marked[Visit{j, guard.col}] = mark
					// get 1st obstacle to the left in 1 above curr row
					checkNewObsLeft(lab, j-1, guard)
				}

				// update obstacles hit, guards new row and new direction
				lab.obsHits = append(lab.obsHits, ObsHit{row: obs[i], col: guard.col, dir: Down})
				guard.row = obs[i] - 1
				guard.dir = Left

				// not finished
				return false
			}
			// no obstacles below our current position, mark all rows in current downward to bottom
			for i := guard.row; i < maxRow; i++ {
				lab.marked[Visit{i, guard.col}] = mark
				// get 1st obstacle to the left in 1 above curr row
				checkNewObsLeft(lab, i-1, guard)
			}
			// reached bottom exit lab
			return true
		}
		// no obstacles found, mark all rows in current downward to end
		for i := guard.row; i < maxRow; i++ {
			lab.marked[Visit{i, guard.col}] = mark
		}
		// reached bottom exit lab
		return true
	}

	if guard.dir == Left {
		obs, ok := lab.obsRows[guard.row]
		// obstacles found
		if ok {
			// iterate through the ordered col obstacle col vals for this row
			for i := len(obs) - 1; i >= 0; i-- {
				// if the row is right of our current position, skip it
				if obs[i] >= guard.col {
					continue
				}
				// we've got the next obstacle left of current pos, get the distance
				guard.path += guard.col - obs[i] - 1

				// mark the path elements
				for j := obs[i] + 1; j < guard.col; j++ {
					lab.marked[Visit{guard.row, j}] = mark
					// get 1st obstacle above in 1 to right of curr col
					checkNewObsUp(lab, j+1, guard)
				}

				// update obstacles hit, guards new row and new direction
				lab.obsHits = append(lab.obsHits, ObsHit{row: guard.row, col: obs[i], dir: Left})
				guard.col = obs[i] + 1
				guard.dir = Up

				// not finished
				return false
			}
			// no obstacles left of our current position, mark all cols in current left to end
			for i := 0; i < guard.col; i++ {
				lab.marked[Visit{guard.row, i}] = mark
				// get 1st obstacle above in 1 to right of curr col
				checkNewObsUp(lab, i+1, guard)
			}
			// reached left edge exit lab
			return true
		}
		// no obstacles found, mark all cols in current left to end
		for i := 0; i < guard.col; i++ {
			lab.marked[Visit{guard.row, i}] = mark
		}
		// reached left edge exit lab
		return true
	}
	// fallthrough for no/invalid direction
	return true
}

func checkNewObsRight(lab *Lab, currRow int, guard *Guard) {
	obsRow, ok := lab.obsRows[currRow]
	if ok {
		for i := 0; i < len(obsRow); i++ {
			if obsRow[i] < guard.col {
				continue // to next obstruction in this row
			}
			// check if hit already in dir Right
			if slices.Contains(lab.obsHits, ObsHit{row: currRow, col: obsRow[i], dir: Right}) {
				if !(lab.start.row == currRow && lab.start.col == obsRow[i]) {
					lab.newObs++
				}
			}
		}
	}
}

func checkNewObsDown(lab *Lab, currCol int, guard *Guard) {
	obsCol, ok := lab.obsCols[currCol]
	if ok {
		for i := 0; i < len(obsCol); i++ {
			if obsCol[i] < guard.col {
				continue // to next obstruction in this col
			}
			// check if hit already in dir Down
			if slices.Contains(lab.obsHits, ObsHit{row: obsCol[i], col: currCol, dir: Down}) {
				if !(lab.start.row == obsCol[i] && lab.start.col == currCol) {
					lab.newObs++
				}
			}
		}
	}
}

func checkNewObsLeft(lab *Lab, currRow int, guard *Guard) {
	obsRow, ok := lab.obsRows[currRow]
	if ok {
		for i := len(obsRow) - 1; i >= 0; i-- {
			if obsRow[i] >= guard.col {
				continue // to next obstruction in this row
			}
			// check if hit already in dir Left
			if slices.Contains(lab.obsHits, ObsHit{row: currRow, col: obsRow[i], dir: Left}) {
				if !(lab.start.row == currRow && lab.start.col == obsRow[i]) {
					lab.newObs++
				}
			}
		}
	}
}

func checkNewObsUp(lab *Lab, currCol int, guard *Guard) {
	obsCol, ok := lab.obsCols[currCol]
	if ok {
		for i := len(obsCol) - 1; i >= 0; i-- {
			if obsCol[i] >= guard.col {
				continue // to next obstruction in this col
			}
			// check if hit already in dir Up
			if slices.Contains(lab.obsHits, ObsHit{row: obsCol[i], col: currCol, dir: Up}) {
				if !(lab.start.row == obsCol[i] && lab.start.col == currCol) {
					lab.newObs++
				}
			}
		}
	}
}

func initialise(lines []string) (Lab, Guard) {
	var lab Lab

	var guard Guard

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
				lab.start = Visit{row, col}
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
