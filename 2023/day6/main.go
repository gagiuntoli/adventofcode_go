package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compute_distance(hold int, duration int) int {
	speed := hold
	remaining_time := duration - hold
	return speed * remaining_time
}

func compute_possibilities(time int, distance int) int {
	possibilities := 0
	for t := 0; t <= time; t++ {
		res_distance := compute_distance(t, time)
		if res_distance > distance {
			possibilities += 1
		}
	}
	return possibilities
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
	times := strings.Fields(words[0])
	distances := strings.Fields(words[1])

	solution1 := 1
	total_time_str := ""
	total_distance_str := ""
	for i := 1; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		total_time_str += times[i]
		total_distance_str += distances[i]
		solution1 *= compute_possibilities(time, distance)
	}

	total_time, _ := strconv.Atoi(total_time_str)
	total_distance, _ := strconv.Atoi(total_distance_str)
	solution2 := compute_possibilities(total_time, total_distance)

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
