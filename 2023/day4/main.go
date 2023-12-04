package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"math"
	"regexp"
	"strings"
)

func get_card_matches(cards []int, winning []int) int {
	matches := 0
	for _, w := range winning {
		if slices.Contains(cards, w) {
			matches += 1
		}
	}
	return matches
}

func main() {
	if len(os.Args) < 2 {
		panic("The program requires the input file path as argument")
	}
	input := os.Args[1]
	dat, err := os.ReadFile(input)
	if err != nil {
		panic("Input file not found")
	}

	words := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	card_count := make(map[int]int)

	solution1 := 0

	for i, card := range words {
		m := regexp.MustCompile(`^Card \d*: `)

		card = m.ReplaceAllString(card, "")
		card = strings.ReplaceAll(card, "  ", " ")

		card_s := strings.Split(card, "|")
		winning_cards := []int{}
		for _, c := range strings.Split(strings.Trim(card_s[0], " "), " ") {
			val, _ := strconv.Atoi(c)
			winning_cards = append(winning_cards, val)
		}

		cards := []int{}
		for _, c := range strings.Split(strings.Trim(card_s[1], " "), " ") {
			val, _ := strconv.Atoi(c)
			cards = append(cards, val)
		}

		matches := get_card_matches(cards, winning_cards)

		var points int
		if matches == 0 {
			points = 0
		} else {
			points = int(math.Pow(2, float64(matches-1)))
		}

		copies := card_count[i+1]
		card_count[i+1] += 1
		for j := 0; j < matches; j++ {
			card_count[i+2+j] += 1 + copies
		}

		solution1 += points
	}

	solution2 := 0
	for _, num_cards := range card_count {
		solution2 += num_cards
	}

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
