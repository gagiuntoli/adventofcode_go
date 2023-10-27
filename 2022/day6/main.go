package main

import (
	"fmt"
	"os"
	// "slices"
	// "strconv"
	"strings"
)

func allDifferentCharacters(word string) bool {
	for i := 0; i < len(word)-1; i++ {
		for j := i + 1; j < len(word); j++ {
			if word[i] == word[j] {
				return false
			}
		}
	}
	return true
}

func findFirstMarkerPosition(line string, size int) int {
	for i := size; i < len(line); i++ {
		if allDifferentCharacters(line[i-size : i]) {
			return i
		}
	}
	return -1
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

	line := strings.Split(string(dat), "\n")[0]

	fmt.Println("solution 1:", findFirstMarkerPosition(line, 4))
	fmt.Println("solution 2:", findFirstMarkerPosition(line, 14))
}
