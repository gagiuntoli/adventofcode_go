package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

//	"github.com/gagiuntoli/adventofcode_go/utils"
)

type Instruction struct {
	dir Direction
	dist int
	color string
}

type Direction uint8

const (
	R = 0
	D = 1
	L = 2
	U = 3
	None = 4
)

func string2dir(dir string) Direction {
	if dir == "U" {
		return U
	} else if dir == "D" {
		return D
	} else if dir == "R" {
		return R
	} else if dir == "L" {
		return L
	}
	panic("direction not recognized")
}

func get_next_point2(i, j, dist int, dir Direction) (int, int) {
	if dir == R {
		return i, j + dist
	} else if dir == L {
		return i, j - dist
	} else if dir == U {
		return i - dist, j
	} else if dir == D {
		return i + dist, j
	}
	panic("invalid direction")
}

type Point struct {
	i, j int
}

func area(points []Point) int {
	area := 0
	points = append(points, points[0])
	for i := 0; i < len(points) - 1; i++ {
		y1 := points[i].i
		x1 := points[i].j
		y2 := points[i+1].i
		x2 := points[i+1].j
		area += x1 * y2 - y1 * x2
	}
	return area / 2
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

	i1 := 0
	j1 := 0
	i2 := 0
	j2 := 0

	points1 := []Point{}
	points2 := []Point{}
	P1 := 0
	P2 := 0

	for _, word := range words {
		points1 = append(points1, Point{i2,j2})
		points2 = append(points2, Point{i1,j1})

		word_s := strings.Split(word, " ")
		dir := word_s[0]
		dist, _ := strconv.Atoi(word_s[1])
		color := word_s[2]
		dist2, _ := strconv.ParseInt(color[2:len(color)-2], 16, 32)
		dir2, _ := strconv.Atoi(color[len(color)-2:len(color)-1])

		P1 += dist
		P2 += int(dist2)
		i2, j2 = get_next_point2(i2, j2, dist, string2dir(dir))
		i1, j1 = get_next_point2(i1, j1, int(dist2), Direction(dir2))
	}

	A1 := area(points1)
	I1 := A1 + 1 - P1 / 2
	solution1 := I1 + P1

	A2 := area(points2)
	I2 := A2 + 1 - P2 / 2
	solution2 := I2 + P2

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
