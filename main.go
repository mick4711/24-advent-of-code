package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	// TODO FindAllString vs re.FindAllStringIndex
	// TODO https://goplay.tools/snippet/QbEL5EngBgY

	ans := getMulResult(string(b))
	fmt.Println("answer =", ans)
}

// Parse the numbers out of the string and check for montonic sequence with diff limits
func getMulResult(s string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	muls := re.FindAllString(s, -1)
	res := 0

	for _, mul := range muls {
		ops := strings.Split(mul[4:len(mul)-1], ",")

		op1, err := strconv.Atoi(ops[0])
		if err != nil {
			panic(err)
		}

		op2, err := strconv.Atoi(ops[1])
		if err != nil {
			panic(err)
		}

		res += op1 * op2
	}

	return res
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
