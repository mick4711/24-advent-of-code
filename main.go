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
	// split out rules and updates
	rnu := strings.Split(file, "\n\n")
	rules := rnu[0]
	middleSum := 0

	// iter thru rules make fwd and bkwd maps
	bkwds := make(map[int][]int, len(rules))
	fwds := make(map[int][]int, len(rules))

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

		trls, ok := fwds[bef]
		if !ok {
			fwds[bef] = []int{aft}
		} else {
			fwds[bef] = append(trls, aft)
		}
	}

	// iter thru updates check fwd and bkwd violations
	updates := rnu[1]
violationCheck:
	for _, update := range strings.Split(updates, "\n") {
		nums := strings.Split(update, ",")
		ns := []int{}

		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			ns = append(ns, n)
		}
		end := len(ns)
		middle := ns[(end-1)/2]

		// a violation is
		// - when a reverse check encounters a fwds match
		// - a forward check encounters a bkwds
		// maybe bkwd checks are enough!
		// forward checks
		// check if any trailing n's are in bkwds
		for i, n := range ns {
			precedings, ok := bkwds[n]
			if !ok {
				continue
			}

			for j := i; j < end; j++ {
				if slices.Contains(precedings, ns[j]) {
					break violationCheck
				}
			}
		}

		middleSum += middle
	}

	return middleSum
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
