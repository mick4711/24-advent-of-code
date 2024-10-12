package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	currentCardCounts := map[int]int{}
	for fileScanner.Scan() {
		cardCounts := cardCount42(fileScanner.Text(), currentCardCounts)
		currentCardCounts = cardCounts
	}

	sum := 0
	for _, v := range currentCardCounts {
		sum += v
	}

	fmt.Println("final sum =", sum)
}

func cardCount42(s string, currentCardCounts map[int]int) map[int]int {
	game := strings.Split(s, ":")
	if len(game) != 2 {
		return currentCardCounts
	}

	gameNumber, err := strconv.Atoi(strings.Fields(game[0])[1])
	if err != nil {
		return currentCardCounts
	}

	//count initial card for this game number
	_, ok := currentCardCounts[gameNumber]
	if !ok {
		currentCardCounts[gameNumber] = 1
	} else {
		currentCardCounts[gameNumber] += 1
	}

	// get winning points, convert to game wins and bump winning card counts
	points := getCardScore41(s)
	score := 0
	if points > 0 {
		score = int(math.Log2(float64(points))) + 1
	}

	// for each card currently held
	for i := 0; i < currentCardCounts[gameNumber]; i++ {
		// for each number up to winning score
		for j := gameNumber + 1; j <= gameNumber+score; j++ {
			currentCardCounts[j] += 1
		}
	}

	return currentCardCounts
}

func getCardScore41(s string) int {
	// filter out game number and get lists of winning numbers and player number
	game := strings.Split(s, ":")
	if len(game) != 2 {
		return 0
	}

	numbers := strings.Split(game[1], "|")
	if len(numbers) != 2 {
		return 0
	}

	winners := make([]int, 0, len(numbers[0]))
	for _, n := range strings.Fields(numbers[0]) {
		num, _ := strconv.Atoi(string(n))
		winners = append(winners, num)
	}

	players := make([]int, 0, len(numbers[1]))
	for _, n := range strings.Fields(numbers[1]) {
		num, _ := strconv.Atoi(string(n))
		players = append(players, num)
	}

	// calculate scores
	score := 0
	for _, winner := range winners {
		if slices.Contains(players, winner) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}
