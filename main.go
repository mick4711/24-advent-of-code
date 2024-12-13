package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	searchTerm    = "XMAS"
	searchTermLen = len(searchTerm)
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
	re := regexp.MustCompile(`X`)
	// split file into lines
	lines := strings.Split(file, "\n")
	for i, s := range lines {
		// find 'x' in each lines
		xs := re.FindAllStringIndex(s, -1)
		for _, v := range xs {
			ans += findW(s, v[0])
			ans += findE(s, v[0])
			ans += findS(i, v[0], lines)
			ans += findN(i, v[0], lines)
			ans += findNW(i, v[0], lines)
			ans += findNE(i, v[0], lines)
			ans += findSW(i, v[0], lines)
			ans += findSE(i, v[0], lines)
		}

		// search along 8 directions for 'mas'
		fmt.Println("ans =", i, ans)
	}

	return ans
}

func findW(s string, col int) int {
	if col+searchTermLen > len(s) {
		return 0
	}

	var word string
	for i := 0; i < searchTermLen; i++ {
		word += string(s[col+i])
	}

	if word == searchTerm {
		return 1
	}

	return 0
}

func findE(s string, col int) int {
	if col < searchTermLen-1 {
		return 0
	}

	var word string
	for i := 0; i < searchTermLen; i++ {
		word += string(s[col-i])
	}

	if word == searchTerm {
		return 1
	}

	return 0
}

func findS(row, col int, lines []string) int {
	if row > len(lines)-searchTermLen {
		return 0
	}

	var word string
	for i := 0; i < searchTermLen; i++ {
		word += string(lines[row+i][col])
	}

	if word == searchTerm {
		return 1
	}

	return 0
}

func findN(row, col int, lines []string) int {
	if row < searchTermLen-1 {
		return 0
	}

	var word string
	for i := 0; i < searchTermLen; i++ {
		word += string(lines[row-i][col])
	}

	if word == searchTerm {
		return 1
	}

	return 0
}

func findNW(row, col int, lines []string) int {
	if row < searchTermLen-1 {
		return 0
	}

	if col+searchTermLen > len(lines[0]) {
		return 0
	}

	var word string
	for i := 0; i < searchTermLen; i++ {
		word += string(lines[row-i][col+i])
	}

	if word == searchTerm {
		return 1
	}

	return 0
}

func findNE(row, col int, lines []string) int {
	if row < searchTermLen-1 {
		return 0
	}

	if col < searchTermLen-1 {
		return 0
	}

	var word string
	for i := 0; i < searchTermLen; i++ {
		word += string(lines[row-i][col-i])
	}

	if word == searchTerm {
		return 1
	}

	return 0
}

func findSW(row, col int, lines []string) int {
	if row > len(lines)-searchTermLen {
		return 0
	}

	if col+searchTermLen > len(lines[0]) {
		return 0
	}

	var word string
	for i := 0; i < searchTermLen; i++ {
		word += string(lines[row+i][col-i])
	}

	if word == searchTerm {
		return 1
	}

	return 0
}

func findSE(row, col int, lines []string) int {
	if row > len(lines)-searchTermLen {
		return 0
	}

	if col < searchTermLen-1 {
		return 0
	}

	var word string
	for i := 0; i < searchTermLen; i++ {
		word += string(lines[row+i][col+i])
	}

	if word == searchTerm {
		return 1
	}

	return 0
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
