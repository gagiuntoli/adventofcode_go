package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	// "errors"
)

func pairsFullyOverlap(a1 int, a2 int, b1 int, b2 int) bool {
	return (a1 >= b1 && a2 <= b2) || (b1 >= a1 && b2 <= a2)
}

func pairsOverlap(a1 int, a2 int, b1 int, b2 int) bool {
	return (a1 <= b1 && a2 >= b1) || (a1 <= b2 && a2 >= b2) || (b1 <= a2 && b2 >= a2)
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

	words := strings.Split(string(dat), "\n")

	fully_overlaped_pairs := 0
	overlaped_pairs := 0
	for _, line := range words {
		if len(line) > 0 {
			pair := strings.Split(line, ",")
			pair1 := strings.Split(pair[0], "-")
			pair2 := strings.Split(pair[1], "-")
			a1, _ := strconv.Atoi(pair1[0])
			a2, _ := strconv.Atoi(pair1[1])
			b1, _ := strconv.Atoi(pair2[0])
			b2, _ := strconv.Atoi(pair2[1])

			if pairsFullyOverlap(a1, a2, b1, b2) {
				fully_overlaped_pairs += 1 
			}
			if pairsOverlap(a1, a2, b1, b2) {
				overlaped_pairs += 1 
			}
		}
	}
	fmt.Println("solution 1:", fully_overlaped_pairs)
	fmt.Println("solution 2:", overlaped_pairs)
}
