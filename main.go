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

	// read input file
	readFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// initialise slices to hold left and right int values
	left := []int{}
	right := []int{}

	// iterate input lines and parse out left and right int values into slices
	for fileScanner.Scan() {
		s := fileScanner.Text()
		left, right = parseNums(s, left, right)
	}

	sim := getSimilarityScore(left, right)

	fmt.Println("similarity score =", sim)
}

// Parse the left and right numbers out of the string and append them to the corresponding laft and right slices
func parseNums(s string, left, right []int) (leftRes, rightRes []int) {
	// get the 2 string numbers from the string
	nums := strings.Fields(s)

	// convert 1st string and append to left slice
	iLeft, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}

	left = append(left, iLeft)

	// convert 2nd string and append to right slice
	iRight, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	right = append(right, iRight)

	return left, right
}

// sort the left and right slices and accumulate the absolute difference values
func getListDistance(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	var dist int

	for i := range len(left) {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}

		dist += diff
	}

	return dist
}

func getSimilarityScore(left, right []int) int {
	// get the grouped counts for numbers in right into a map [int]int
	rightCounts := make(map[int]int)
	curr := right[0]

	for _, v := range right {
		if v != curr {
			curr = v
		}

		rightCounts[curr]++
	}

	// iterate through left and accumulate counts*key
	sim := 0
	for _, v := range left {
		sim += v * rightCounts[v]
	}

	return sim
}
