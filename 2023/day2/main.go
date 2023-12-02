package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func is_valid_game_1(game string) (bool, int) {
	m := regexp.MustCompile(`^Game \d+: `)

	game = m.ReplaceAllString(game, "")
	rounds := strings.Split(game, ";")

	is_valid := true

	min_red := 0
	min_green := 0
	min_blue := 0
	for ri, round := range rounds {
		red := 0
		green := 0
		blue := 0

		round_s := strings.ReplaceAll(round, ",", "")

		values := strings.Split(round_s, " ")
		for i, color := range values {
			if i > 0 {
				val, _ := strconv.Atoi(values[i-1])
				if color == "red" {
					red += val
				}else if color == "green" {
					green += val
				}else if color == "blue" {
					blue += val
				}
			}
		}

		if ri == 0 {
			min_red = red
			min_green = green
			min_blue = blue
		} else {
			min_red = max(red, min_red)
			min_green = max(green, min_green)
			min_blue = max(blue, min_blue)
		}
		if !(red <= 12 && green <= 13 && blue <= 14) {
			is_valid = false
		} 
	}
	power := min_red * min_green * min_blue
	return is_valid, power

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
	for i, game := range words {
		is_valid_game, power := is_valid_game_1(game)

		if is_valid_game {
			solution1 += i+1
		}
		solution2 += power
	}

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
