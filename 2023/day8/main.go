package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gagiuntoli/adventofcode_go/utils"
)

type Direction struct {
	left  string
	right string
}

func follow_instructions(directions map[string]Direction, instructions string, start string) (int, string) {
	var destiny string

	for i, c := range instructions {
		destiny = apply_step(directions, byte(c), start)
		start = destiny

		if destiny == "ZZZ" {
			return i + 1, destiny
		}
	}
	return len(instructions), destiny
}

func apply_step(directions map[string]Direction, step byte, start string) string {
	if step == 'L' {
		return directions[start].left
	} else if step == 'R' {
		return directions[start].right
	} else {
		err := fmt.Errorf("Instruction not recognized %b", step)
		panic(err)
	}
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

	parts := strings.Split(strings.Trim(string(dat), "\n"), "\n\n")
	instructions := strings.Trim(parts[0], " ")
	map_directions := strings.Split(parts[1], "\n")

	directions := make(map[string]Direction)

	starting_nodes := []string{}

	for _, line := range map_directions {
		parts := strings.Split(line, "=")
		start := strings.Trim(parts[0], " ")
		destinies := strings.Split(strings.Trim(parts[1], " ()"), ",")
		left := strings.Trim(destinies[0], " ")
		right := strings.Trim(destinies[1], " ")
		directions[start] = Direction{left, right}
		if start[len(start)-1] == 'A' {
			starting_nodes = append(starting_nodes, start)
		}
	}

	solution1 := 0
	start := "AAA"
	for {
		steps, destiny := follow_instructions(directions, instructions, start)
		solution1 += steps
		start = destiny
		if destiny == "ZZZ" {
			break
		}
	}

	destiny_nodes := make([]string, len(starting_nodes))

	type Check struct {
		column int
		a      string
		c      byte
	}
	checkpoints := make(map[Check]bool)
	checkpoint_distances := make([]uint64, len(starting_nodes))

	for k := 0; ; k++ {
		step := instructions[k%len(instructions)]
		for i, start := range starting_nodes {
			destiny := apply_step(directions, step, start)
			destiny_nodes[i] = destiny
			if destiny[len(destiny)-1] == 'Z' {
				if checkpoint_distances[i] == 0 {
					check := Check{i, destiny, step}
					_, ok := checkpoints[check]
					if !ok {
						checkpoint_distances[i] = uint64(k + 1)
					}
				}
			}
			starting_nodes[i] = destiny
		}

		every := true

		for _, c := range checkpoint_distances {
			if c == 0 {
				every = false
				break
			}
		}

		if every {
			break
		}
	}

	solution2 := utils.LCM(checkpoint_distances)

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
