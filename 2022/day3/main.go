package main

import (
	"fmt"
	"os"
	"strings"
	"errors"
)

func findRepeatedCharTwo(a string, b string) (byte, error) {
	ab := []byte(a)
	bb := []byte(b)
	for _, ac := range ab {
		for _, bc := range bb {
			if ac == bc {
				return ac, nil
			}
		}
	}
	return 0, errors.New("no repeated character")
}

func findRepeatedCharThree(a string, b string, c string) (byte, error) {
	ab := []byte(a)
	bb := []byte(b)
	cb := []byte(c)
	for _, ac := range ab {
		for _, bc := range bb {
			if ac == bc {
				for _, cc := range cb {
					if ac == cc {
						return ac, nil
					}
				}
			}
		}
	}
	return 0, errors.New("no repeated character")
}

func computePoints(c byte) int {
	if byte('A') <= c && c <= byte('Z') {
		return int(c) - int(byte('A')) + 27
	} else if byte('a') <= c && c <= byte('z') {
		return int(c) - int(byte('a')) + 1
	} else {
		panic("Invalid character")
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

	words := strings.Split(string(dat), "\n")

	points1 := 0
	points2 := 0
	arr := []string{}
	for _, line := range words {
		if len(line) > 0 {
			word1 := line[:len(line)/2]
			word2 := line[len(line)/2:]
			c, err := findRepeatedCharTwo(word1, word2)
			if err == nil {
				points1 += computePoints(c)
			} else {
				fmt.Println("no repetition found for", word1, word2)
			}
			arr = append(arr, line)
			if len(arr) == 3 {
				c, err := findRepeatedCharThree(arr[0], arr[1], arr[2])
				if err == nil {
					points2 += computePoints(c)
				} else {
					fmt.Println("no repetition found for", arr[0], arr[1], arr[2])
				}
				arr = []string{}
			}
		}
	}
	fmt.Println("solution 1:", points1)
	fmt.Println("solution 2:", points2)
}
