package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	pathLength := 0
	// store obs in row and col maps
	// find caret and starting direction up
	const (
		up = iota
		right
		down
		left
	)
	lines := strings.Split(file, "\n")
	obsRows := make(map[int][]int, len(lines))
	obsCols := make(map[int][]int, len(lines[0]))
	startRow := 0
	startCol := 0
	dir := up

	for row, line := range lines {
		for col, cell := range line {
			if string(cell) == "#" {
				_, ok := obsRows[row]
				if !ok {
					obsRows[row] = []int{col}
				} else {
					obsRows[row] = append(obsRows[row], col)
				}

				_, ok = obsCols[col]
				if !ok {
					obsCols[col] = []int{row}
				} else {
					obsCols[col] = append(obsCols[col], row)
				}

				pathLength++
			}

			if string(cell) == "^" {
				startRow = row
				startCol = col
			}
		}
	}

	fmt.Println(obsRows)
	fmt.Println(obsCols)
	fmt.Println(startRow)
	fmt.Println(startCol)
	fmt.Println(dir)
	// find 1st obstruction in column
	// change direction right
	// find obstr in row
	// change direction down
	obs, ok := obsCols[startCol]
	if ok {
		for i := len(obs) - 1; i >= 0; i-- {
			if obs[i] < startRow {
				// turn right and find obs in row
				fmt.Println("hit obs:", obs[i], startCol)
			}
		}
	}

	return pathLength
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
