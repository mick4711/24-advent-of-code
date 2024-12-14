package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	ans := len(file)
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
