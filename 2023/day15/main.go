package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hash(str string) int {
	current := 0
	for _, v := range str {
		current += int(v)
		current *= 17
		current %= 256
	}

	return current
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

	words := strings.Split(strings.Trim(string(dat), "\n"), ",")

	type Lens struct {
		Label string
		Focal int
	}

	type Box struct {
		Lens []Lens
	}

	boxes := [256]Box{}

	solution1 := 0
	for _, word := range words {
		solution1 += hash(word)

		var label string
		var focal int
		if strings.Contains(word, "-") {
			label = word[:len(word)-1]
			focal = -1
		} else {
			s := strings.Split(word, "=")
			label = s[0]
			focal, _ = strconv.Atoi(s[1])
		}
		box_num := hash(label)

		found := false
		for i, lens := range boxes[box_num].Lens {
			if lens.Label == label {
				found = true
				if focal == -1 {
					boxes[box_num].Lens = append(boxes[box_num].Lens[:i], boxes[box_num].Lens[i+1:]...)
				} else {
					boxes[box_num].Lens[i].Focal = focal
				}
				break
			}
		}
		if !found && focal != -1 {
			boxes[box_num].Lens = append(boxes[box_num].Lens, Lens{label, focal})

		}

	}

	solution2 := 0
	for box_num := 0; box_num < len(boxes); box_num++ {
		for i, lens := range boxes[box_num].Lens {
			solution2 += (box_num + 1) * (i + 1) * lens.Focal
		}
	}

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
