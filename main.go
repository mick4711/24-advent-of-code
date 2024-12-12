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

	ans := getXmasCount(string(b))
	fmt.Println("ans =", ans)
}

func getXmasCount(file string) int {
	// split file into lines
	lines := strings.Split(file, "\n")
	for i, s := range lines {
		// find 'x' in each lines
		// search along 8 directions for 'mas'
		fmt.Println("line =", i, s)
	}

	return len(lines)
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
