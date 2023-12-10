package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "github.com/gagiuntoli/adventofcode_go/utils"
)

func compute_difference(sequence []int) []int {
	var difference []int
	for i := 1; i < len(sequence); i++ {
		difference = append(difference, sequence[i]-sequence[i-1])
	}
	return difference
}

func find_next_element(sequence []int) int {
	differences := sequence
	last_value_stack := []int{}

	for {
		last_value_stack = append(last_value_stack, differences[len(differences)-1])

		if utils.All(differences, 0) {
			break
		}

		differences = compute_difference(differences)
	}

	return utils.ArraySum(last_value_stack)
}

func find_first_element(sequence []int) int {
	differences := sequence
	first_value_stack := []int{}

	for {
		first_value_stack = append(first_value_stack, differences[0])

		if utils.All(differences, 0) {
			break
		}

		differences = compute_difference(differences)
	}

	for i := len(first_value_stack) - 1; i > 0; i-- {
		first_value_stack[i-1] -= first_value_stack[i]
	}

	return first_value_stack[0]
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
		numbers := strings.Split(word, " ")
		sequence := []int{}
		for _, number_s := range numbers {
			number, _ := strconv.Atoi(number_s)
			sequence = append(sequence, number)
		}
		solution1 += find_next_element(sequence)
		solution2 += find_first_element(sequence)
	}

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
