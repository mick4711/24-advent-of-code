package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
		if isSafeReport(fileScanner.Text()) {
			safeReports++
		}
	}

	fmt.Println("safe reports count =", safeReports)
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
func isSafeReport(s string) bool {
	const (
		maxDiff = 3
		minDiff = 1
	)
	// get the string numbers from the string
	nums := strings.Fields(s)

	first, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}

	second, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	var increasing, decreasing bool

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

	for i := 2; i < len(nums); i++ {
		curr, err := strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}

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
