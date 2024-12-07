package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// set input file name
	inputFileName := getInputFileName()

	// read input file
	readFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// initialise safe counter
	safeReports := 0

	// iterate input lines checking for safe
	for fileScanner.Scan() {
		sl := getIntSlice(fileScanner.Text())
		if isSafeReport(sl) {
			safeReports++
			continue
		}
		for i, _ := range sl {
			dl := slices.Delete(slices.Clone(sl), i, i+1)
			if isSafeReport(dl) {
				safeReports++
				break
			}
		}
	}

	fmt.Println("safe reports count =", safeReports)
}

func getIntSlice(s string) []int {
	// get the string numbers from the string
	nums := strings.Fields(s)
	sl := make([]int, len(nums))

	for i, v := range nums {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		sl[i] = n
	}

	return sl
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

// Parse the numbers out of the string and check for montonic sequence with diff limits
func isSafeReport(sl []int) bool {
	const (
		maxDiff = 3
		minDiff = 1
	)

	var increasing, decreasing bool

	first, second := sl[0], sl[1]

	switch {
	case first < second:
		if second-first > maxDiff {
			return false
		}

		increasing = true
	case first > second:
		if first-second > maxDiff {
			return false
		}

		decreasing = true
	default: // not increasing or decreasing => not safe
		return false
	}

	prev := second

	for i := 2; i < len(sl); i++ {
		curr := sl[i]
		diff := 0

		switch {
		case increasing:
			diff = curr - prev
		case decreasing:
			diff = prev - curr
		}

		if diff > maxDiff || diff < minDiff {
			return false
		}

		prev = curr
	}

	return true
}
