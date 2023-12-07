package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func card_counter(cards string) map[string]int {
	count := map[string]int{}
	for _, c := range cards {
		count[string(c)] += 1
	}
	return count
}

func is_five_of_a_kind(card_counter map[string]int) bool {
	return len(card_counter) == 1
}

func is_four_of_a_kind(card_counter map[string]int) bool {
	if len(card_counter) == 2 {
		for _, value := range card_counter {
			return value == 1 || value == 4
		}
	}
	return false
}

func is_full_hause(card_counter map[string]int) bool {
	if len(card_counter) == 2 {
		for _, value := range card_counter {
			return value == 3 || value == 2
		}
	}
	return false
}

func is_three_of_a_kind(card_counter map[string]int) bool {
	if len(card_counter) == 3 {
		for _, value := range card_counter {
			if value == 3 {
				return true
			}
		}
	}
	return false
}

func is_two_pair(card_counter map[string]int) bool {
	if len(card_counter) == 3 {
		for _, value := range card_counter {
			if value == 2 {
				return true
			}
		}
	}
	return false
}

func is_one_pair(card_counter map[string]int) bool {
	if len(card_counter) == 4 {
		for _, value := range card_counter {
			if value == 2 {
				return true
			}
		}
	}
	return false
}

func is_high_card(card_counter map[string]int) bool {
	return len(card_counter) == 5
}

func compute_power(card_counter map[string]int) int {
	if is_high_card(card_counter) {
		return 0
	} else if is_one_pair(card_counter) {
		return 1
	} else if is_two_pair(card_counter) {
		return 2
	} else if is_two_pair(card_counter) {
		return 3
	} else if is_three_of_a_kind(card_counter) {
		return 4
	} else if is_full_hause(card_counter) {
		return 5
	} else if is_four_of_a_kind(card_counter) {
		return 6
	} else if is_five_of_a_kind(card_counter) {
		return 7
	}

	panic("card couldn't be identified")
}

func compare_cards_second(cards1 string, cards2 string, card_powers map[string]int) bool {
	for i := range cards1 {
		if cards1[i] != cards2[i] {
			return card_powers[string(cards1[i])] > card_powers[string(cards2[i])]
		}
	}
	return false
}

func use_jokers(card_counter map[string]int) map[string]int {
	Jcards := card_counter[string("J")]

	if Jcards != 0 {
		key_max := ""
		val_max := 0
		for key, val := range card_counter {
			if key != string("J") && val > val_max {
				key_max = key
				val_max = val
			}
		}
		card_counter[key_max] += Jcards
		delete(card_counter, string("J"))
	}

	return card_counter
}

func compute_points(hands []Hand) int {
	points := 0
	for i, hand := range hands {
		points += (i + 1) * hand.bid
	}
	return points
}

type Hand struct {
	cards string
	bid   int
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
	hands := []Hand{}

	for _, word := range words {
		word_s := strings.Split(word, " ")
		bid, _ := strconv.Atoi(word_s[1])
		hands = append(hands, Hand{cards: word_s[0], bid: bid})
	}

	card_powers := map[string]int{"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1}

	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i]
		h2 := hands[j]

		cc1 := card_counter(h1.cards)
		cc2 := card_counter(h2.cards)

		p1 := compute_power(cc1)
		p2 := compute_power(cc2)

		if p1 != p2 {
			return p1 < p2
		}

		return !compare_cards_second(h1.cards, h2.cards, card_powers)
	})

	solution1 := compute_points(hands)

	card_powers = map[string]int{"A": 13, "K": 12, "Q": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1}

	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i]
		h2 := hands[j]

		cc1 := card_counter(h1.cards)
		cc2 := card_counter(h2.cards)

		cc1 = use_jokers(cc1)
		cc2 = use_jokers(cc2)

		p1 := compute_power(cc1)
		p2 := compute_power(cc2)

		if p1 != p2 {
			return p1 < p2
		}

		return !compare_cards_second(h1.cards, h2.cards, card_powers)
	})

	solution2 := compute_points(hands)

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
