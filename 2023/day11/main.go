package main

import (
	"fmt"
	"os"
	"strings"
	//utils "github.com/gagiuntoli/adventofcode_go/utils"
)

func col_empty(universe []string, col int) bool {
	for row := range universe {
		if universe[row][col] != '.' {
			return false
		}
	}
	return true
}

func row_empty(universe []string, row int) bool {
	return universe[row] == strings.Repeat(".", len(universe[0]))
}

func Abs[T int](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}

type Point struct {
	i, j int
}

func distance(universe []string, expansion int, point1, point2 Point) int {
	dist := Abs(point2.i-point1.i) + Abs(point1.j-point2.j)

	empty_spaces := 0
	var istart, iend int
	if point1.i <= point2.i {
		istart, iend = point1.i+1, point2.i
	} else {
		istart, iend = point2.i, point1.i+1
	}
	for i := istart; i < iend; i++ {
		if row_empty(universe, i) {
			empty_spaces += 1
		}
	}

	if point1.j <= point2.j {
		istart, iend = point1.j+1, point2.j
	} else {
		istart, iend = point2.j, point1.j+1
	}
	for i := istart; i < iend; i++ {
		if col_empty(universe, i) {
			empty_spaces += 1
		}
	}
	return dist + empty_spaces*expansion - empty_spaces
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

	universe := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	galaxies := []Point{}
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[i]); j++ {
			if universe[i][j] == '#' {
				galaxies = append(galaxies, Point{i, j})
			}
		}
	}

	solution1 := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			dist := distance(universe, 2, galaxies[i], galaxies[j])
			solution1 += dist
		}
	}

	solution2 := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			dist := distance(universe, 1000000, galaxies[i], galaxies[j])
			solution2 += dist
		}
	}

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
