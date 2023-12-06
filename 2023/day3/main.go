package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	// "strconv"
	"strings"
)

func has_adjacent_symbol(engine []string, i int, j int) bool {
	positions := []struct {
		i, j int
	}{
		{i + 1, j},
		{i + 1, j + 1},
		{i + 1, j - 1},
		{i, j + 1},
		{i, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i - 1, j - 1},
	}
	for _, position := range positions {
		i = position.i
		j = position.j
		if i >= 0 && i < len(engine) && j >= 0 && j < len(engine[0]) {
			if !unicode.IsDigit(rune(engine[i][j])) && engine[i][j] != '.' {
				return true
			}
		}
	}
	return false
}

func parse_number(line string, i int) (int, int) {
	j := i

	for ; j < len(line) && unicode.IsDigit(rune(line[j])); j++ {
	}

	if j == i {
		return -1, -1
	} else {
		val, _ := strconv.Atoi(line[i:j])
		return val, j
	}
}

func compute_sum(engine []string) int {
	sum := 0
	for i, line := range engine {
		j := 0
		for j < len(line) {
			val, offset := parse_number(line, j)
			if offset == -1 {
				j++
			} else {
				for k := j; k < offset; k++ {
					if has_adjacent_symbol(engine, i, k) {
						sum += val
						break
					}
				}
				j = offset
			}
		}
	}
	return sum
}

func find_gears_sum(engine []string) int {
	sum := 0
	for i, line := range engine {
		for j, c := range line {
			if c == '*' {
				numbers := find_adjacent_numbers(engine, i, j)
				if len(numbers) == 2 {
					sum += numbers[0] * numbers[1]
				}
			}
		}
	}
	return sum
}

func find_numbers_in_positions(engine []string, positions []struct{ i, j int }) []int {
	numbers := []int{}
	for _, position := range positions {
		ii := position.i
		jj := position.j
		if ii >= 0 && ii < len(engine) && jj >= 0 && jj < len(engine[0]) {
			if unicode.IsDigit(rune(engine[ii][jj])) {
				k := jj - 1
				for ; k >= 0 && unicode.IsDigit(rune(engine[ii][k])); k-- {
				}
				k++

				val, _ := parse_number(engine[ii], k)
				if val != -1 {
					numbers = append(numbers, val)
				}
			}
		}
	}
	return numbers
}

func find_adjacent_numbers(engine []string, i int, j int) []int {
	numbers := []int{}

	positions := []struct{ i, j int }{{i, j - 1}, {i, j + 1}}

	tmp := find_numbers_in_positions(engine, positions)
	numbers = append(numbers, tmp...)

	positions = []struct{ i, j int }{{i + 1, j}}

	tmp = find_numbers_in_positions(engine, positions)
	numbers = append(numbers, tmp...)

	if len(tmp) == 0 {
		positions = []struct{ i, j int }{{i + 1, j + 1}, {i + 1, j - 1}}
		tmp = find_numbers_in_positions(engine, positions)
		numbers = append(numbers, tmp...)
	}

	positions = []struct{ i, j int }{{i - 1, j}}

	tmp = find_numbers_in_positions(engine, positions)
	numbers = append(numbers, tmp...)
	if len(tmp) == 0 {
		positions = []struct{ i, j int }{{i - 1, j + 1}, {i - 1, j - 1}}
		tmp = find_numbers_in_positions(engine, positions)
		numbers = append(numbers, tmp...)
	}
	return numbers
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

	engine := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	solution1 := compute_sum(engine)
	solution2 := find_gears_sum(engine)

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
