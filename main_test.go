package main

import (
	"testing"
)

func TestGetListDistance(t *testing.T) {
	listDistance := getListDistance()
	if listDistance != 11 {
		t.Errorf("list distance, got:%v, want:%v", listDistance, 13)
	}
}

/*
func TestGetPartNo41(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"", 0},
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	sum := 0
	for _, test := range tests {
		got := getCardScore41(test.input)
		if got != test.want {
			t.Errorf("getCardScore41(%v), got:%v, want:%v", test.input, got, test.want)
		}
		sum += got
	}
	if sum != 13 {
		t.Errorf("sum of card scores, got:%v, want:%v", sum, 13)
	}

}

func TestCardCount42(t *testing.T) {
	var tests = []struct {
		input string
		want  map[int]int
	}{
		{"", map[int]int{}},
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1}},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", map[int]int{1: 1, 2: 2, 3: 3, 4: 3, 5: 1}},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", map[int]int{1: 1, 2: 2, 3: 4, 4: 7, 5: 5}},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", map[int]int{1: 1, 2: 2, 3: 4, 4: 8, 5: 13}},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", map[int]int{1: 1, 2: 2, 3: 4, 4: 8, 5: 14}},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", map[int]int{1: 1, 2: 2, 3: 4, 4: 8, 5: 14, 6: 1}},
	}

	currentCardCounts := map[int]int{}
	for _, test := range tests {
		got := cardCount42(test.input, currentCardCounts)
		currentCardCounts = got
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("cardCount42(%v), got:%v, want:%v", test.input, got, test.want)
		}
	}

	sum := 0
	for _, v := range currentCardCounts {
		sum += v
	}
	if sum != 30 {
		t.Errorf("sum of card counts, got:%v, want:%v", sum, 30)
	}

}
*/
