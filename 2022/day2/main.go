package main

import (
	"fmt"
	"os"
	"strings"
)


func pointsOne(a byte, b byte) int {
	if a == 'A' && b == 'X' { // Rock
		return 1 + 3
	} else if a == 'B' && b == 'Y' { // Paper
		return 2 + 3
	} else if a == 'C' && b == 'Z' { // Scissors
		return 3 + 3
	} else if a == 'A' && b == 'Y' {
		return 2 + 6
	} else if a == 'A' && b == 'Z' {
		return 3 + 0
	} else if a == 'B' && b == 'X' {
		return 1 + 0
	} else if a == 'B' && b == 'Z' {
		return 3 + 6
	} else if a == 'C' && b == 'X' {
		return 1 + 6
	} else if a == 'C' && b == 'Y' {
		return 2 + 0
	}
	return 0
}

func pointsTwo(a byte, b byte) int {
	if a == 'A' {
		if b == 'X' {
			return 3 + 0
		} else if b == 'Y' {
			return 1 + 3
		} else if b == 'Z' {
			return 2 + 6
		}
	} else if a == 'B' { 
		if b == 'X' {
			return 1 + 0
		} else if b == 'Y' {
			return 2 + 3
		} else if b == 'Z' {
			return 3 + 6
		}
	} else if a == 'C' { // Scissors
		if b == 'X' {
			return 2 + 0
		} else if b == 'Y' {
			return 3 + 3
		} else if b == 'Z' {
			return 1 + 6
		}
	}
	return 0
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

	points1 := 0
	points2 := 0
	for _, line := range words {
		values := strings.Split(line, " ")
		if len(values) == 2 {
			v1 := []byte(values[0])[0]
			v2 := []byte(values[1])[0]
			points1 += pointsOne(v1, v2)
			points2 += pointsTwo(v1, v2)
		}
	}
	fmt.Println("solution 1:", points1)
	fmt.Println("solution 2:", points2)
}
