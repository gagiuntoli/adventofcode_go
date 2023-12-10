package main

import (
	"fmt"
	"os"
	"strings"
	//utils "github.com/gagiuntoli/adventofcode_go/utils"
)

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

	for _, word := range words {
		fmt.Println(word)
	}

	solution1 := 0
	solution2 := 0
	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
