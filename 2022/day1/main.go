package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "github.com/gagiuntoli/adventofcode_go/utils"
)

func findMinIndex(arr []int) int {
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[index] {
			index = i
		}
	}
	return index
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

	arr := []int{0, 0, 0}

	max_calories := 0
	calories := 0
	for _, elem := range words {
		tmp, _ := strconv.Atoi(elem)
		calories += tmp
		if elem == "" {
			if calories > max_calories {
				max_calories = calories
			}
			minIndex := findMinIndex(arr)
			if calories > arr[minIndex] {
				arr[minIndex] = calories
			}
			calories = 0
		}
	}
	fmt.Println("solution 1:", max_calories)
	fmt.Println("solution 2:", utils.ArraySum(arr))
}
