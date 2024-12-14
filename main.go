package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// set input file name
	inputFileName := getInputFileName()

	b, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	ans := getXmasCount(string(b))
	fmt.Println("ans =", ans)
}

func getXmasCount(file string) int {
	ans := 0
	re := regexp.MustCompile(`A`)
	// split file into lines
	lines := strings.Split(file, "\n")
	rowLen := len(lines[0])
	rowCount := len(lines)

	for row, line := range lines {
		// eliminate edge rows
		if row < 1 || row >= rowCount-1 {
			continue
		}
		// find 'A' in each lines
		xs := re.FindAllStringIndex(line, -1)
		for _, v := range xs {
			// eliminate edge cols
			col := v[0]
			if col < 1 || col >= rowLen-1 {
				continue
			}

			// get the 4 corners
			tl := string(lines[row-1][col-1])
			tr := string(lines[row-1][col+1])
			bl := string(lines[row+1][col-1])
			br := string(lines[row+1][col+1])

			// search top left for M or S and bottom right for conjugate
			switch tl {
			case "M":
				if br != "S" {
					continue
				}
			case "S":
				if br != "M" {
					continue
				}
			default:
				continue
			}

			// search top right for M or S and bottom left for conjugate
			switch tr {
			case "M":
				if bl != "S" {
					continue
				}
			case "S":
				if bl != "M" {
					continue
				}
			default:
				continue
			}

			ans++
		}
	}

	return ans
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
