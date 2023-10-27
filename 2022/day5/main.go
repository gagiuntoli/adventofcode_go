package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

	words := strings.Split(string(dat), "\n")

	stacks := [15][]byte{}
	i := 0
	for ; len(words[i]) > 0; i++ {
		if strings.ContainsAny(words[i], "[") {
			words[i] = strings.Replace(words[i], "[", " ", -1)
			words[i] = strings.Replace(words[i], "]", " ", -1)
			for j, w := range words[i] {
				if (j-1)%4 == 0 {
					if w != ' ' {
						stacks[(j-1)/4] = append(stacks[(j-1)/4], byte(w))
					}
				}
			}
		}
	}
	for _, stack := range stacks {
		slices.Reverse(stack)
	}

	stacks1 := [15][]byte{}
	stacks2 := [15][]byte{}
	for index := 0; index < len(stacks); index++ {
		for _, elem := range stacks[index] {
			stacks1[index] = append(stacks1[index], elem)
			stacks2[index] = append(stacks2[index], elem)
		}
	}

	i += 1
	for ; i < len(words); i++ {
		if len(words[i]) > 0 {
			words[i] = strings.Replace(words[i], "move", "", -1)
			words[i] = strings.Replace(words[i], "from", "", -1)
			words[i] = strings.Replace(words[i], "to", "", -1)
			numbers := strings.Split(words[i], "  ")
			qty, _ := strconv.Atoi(strings.Trim(numbers[0], " "))
			from, _ := strconv.Atoi(strings.Trim(numbers[1], " "))
			to, _ := strconv.Atoi(strings.Trim(numbers[2], " "))
			from -= 1
			to -= 1
			stacks2[to] = append(stacks2[to], stacks2[from][len(stacks2[from])-qty:]...)
			stacks2[from] = stacks2[from][:len(stacks2[from])-qty]
			for ; qty > 0; qty-- {
				value := stacks1[from][len(stacks1[from])-1]
				stacks1[from] = stacks1[from][:len(stacks1[from])-1]
				stacks1[to] = append(stacks1[to], value)
			}
		}
	}

	solution1 := []byte{}
	for _, stack := range stacks1 {
		if len(stack) > 0 {
			solution1 = append(solution1, stack[len(stack)-1])
		}
	}
	solution2 := []byte{}
	for _, stack := range stacks2 {
		if len(stack) > 0 {
			solution2 = append(solution2, stack[len(stack)-1])
		}
	}
	fmt.Println("solution 1:", string(solution1))
	fmt.Println("solution 2:", string(solution2))
}
