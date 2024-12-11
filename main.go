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

	s := string(b)
	// get slice of do() indices
	reDo := regexp.MustCompile(`do\(\)`)
	dos := reDo.FindAllStringIndex(s, -1)
	// get slice of don't() indices
	reDont := regexp.MustCompile(`don't\(\)`)
	donts := reDont.FindAllStringIndex(s, -1)
	// set curr do index = 0
	pDo := -1
	currDo := []int{0, 0}
	// set curr dont index = 1st dont index
	pDont := 0
	currDont := donts[pDont]
	ans := 0
	// get mulResult of curr do to curr dont
	for {
		searchString := s[currDo[1]:currDont[0]]
		log.Println(searchString)
		ans += getMulResult(searchString)
		// find next index of do > curr dont and make it curr do
		pDo++
		if pDo < len(dos) {
			currDo = dos[pDo]
		} else {
			break
		}

		// TODO check for currDo > currDont
		// for currDo[pDo] < currDont[pDont] {
		// 	pDo++
		// 	if pDo > len(dos)-1 {
		// 		break
		// 	}
		// }

		// find next index of dont > curr do and make it curr dont
		pDont++
		if pDont < len(donts) {
			currDont = donts[pDont]
		} else {
			currDont = []int{len(s) - 1, len(s) - 1}
		}

		// TODO check for currDo > currDont
		// for currDo[pDo] < currDont[pDont] {
		// 	pDont++
		// 	if pDont > len(donts)-1 {
		// 		break
		// 	}
		// }
	}

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
