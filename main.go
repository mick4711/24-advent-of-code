package main

import (
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

	b, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	ans := getMiddleSum(string(b))
	fmt.Println("ans =", ans)
}

func getMiddleSum(file string) int {
	middleSum := 0
	// split out rules and updates
	rnu := strings.Split(file, "\n\n")

	// iter thru rules make bkwd map
	bkwds := getBkwdRules(rnu[0])

	// iter thru updates check violations
	for _, update := range strings.Split(rnu[1], "\n") {
		ns := getNumericSlice(update)

		if !isViolation(ns, bkwds) {
			middleSum += ns[(len(ns)-1)/2]
		}
	}

	return middleSum
}

// a violation is when a forward check encounters a bkwds entry
func isViolation(ns []int, bkwds map[int][]int) bool {
	isViolation := false

	for i, n := range ns {
		precedings, ok := bkwds[n]
		if !ok {
			continue
		}

		// check if any trailing n's are in bkwds
		for j := i; j < len(ns); j++ {
			if slices.Contains(precedings, ns[j]) {
				isViolation = true
			}
		}
	}

	return isViolation
}

func getNumericSlice(update string) []int {
	nums := strings.Split(update, ",")
	ns := []int{}

	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		ns = append(ns, n)
	}

	return ns
}

func getBkwdRules(rules string) map[int][]int {
	bkwds := make(map[int][]int, len(rules))

	for _, v := range strings.Split(rules, "\n") {
		r := strings.Split(v, "|")

		aft, err := strconv.Atoi(r[1])
		if err != nil {
			panic(err)
		}

		bef, err := strconv.Atoi(r[0])
		if err != nil {
			panic(err)
		}

		pres, ok := bkwds[aft]
		if !ok {
			bkwds[aft] = []int{bef}
		} else {
			bkwds[aft] = append(pres, bef)
		}
	}

	return bkwds
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
