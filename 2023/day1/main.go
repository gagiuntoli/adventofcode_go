package main

import (
	"fmt"
	"os"
	// "strconv"
	"strings"

	"unicode"
)

func extractNumberPartA(word string) int {
	lastDigit := 'a'
	firstDigit := 'a'
	for _, char := range word {
		if unicode.IsDigit(char) {
			lastDigit = char - '0'
			if firstDigit == 'a' {
				firstDigit = lastDigit
			}
		}
	}
	return int(firstDigit)*10 + int(lastDigit)
}

func extractNumberPartB(word string) int {
	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	lastDigit := 'a'
	firstDigit := 'a'
	for i, char := range word {
		if unicode.IsDigit(char) {
			lastDigit = char - '0'
		} else {
			for j, digit := range digits {
				if i+len(digit) <= len(word) && digit == word[i:i+len(digit)] {
					lastDigit = rune(byte(j))
				}
			}
		}
		if firstDigit == 'a' {
			firstDigit = lastDigit
		}
	}
	return int(firstDigit)*10 + int(lastDigit)
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

	solution1 := 0
	solution2 := 0
	for _, word := range words {
		solution1 += extractNumberPartA(word)
		solution2 += extractNumberPartB(word)
	}

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
