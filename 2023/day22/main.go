package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Brick struct {
	x1, y1, z1 int
	x2, y2, z2 int
	name       string
}

func intersect_xy(b1, b2 Brick) bool {
	if ((b1.x1 <= b2.x1 && b2.x1 <= b1.x2) || (b1.x1 <= b2.x2 && b2.x2 <= b1.x2)) && ((b2.y1 <= b1.y1 && b1.y1 <= b2.y2) || (b2.y1 <= b1.y2 && b1.y2 <= b2.y2)) {
		return true
	}
	if ((b2.x1 <= b1.x1 && b1.x1 <= b2.x2) || (b2.x1 <= b1.x2 && b1.x2 <= b2.x2)) && ((b1.y1 <= b2.y1 && b2.y1 <= b1.y2) || (b1.y1 <= b2.y2 && b2.y2 <= b1.y2)) {
		return true
	}
	return false
}

func move_down(bricks []Brick, i int) (Brick, int) {
	brick := bricks[i]
	height := brick.z2 - brick.z1

	indeces := []int{}
	bricks_z2 := []Brick{}

	for k := i + 1; k < len(bricks); k++ {
		indeces = append(indeces, k)
		bricks_z2 = append(bricks_z2, bricks[k])
	}

	// We sort putting first the highers
	sort.Slice(bricks_z2, func(i, j int) bool {
		return bricks_z2[i].z2 > bricks_z2[j].z2
	})

	for k := 0; k < len(bricks_z2); k++ {
		if intersect_xy(bricks_z2[k], brick) {
			brick.z1 = bricks_z2[k].z2 + 1
			brick.z2 = brick.z1 + height
			return brick, indeces[k] - 1
		}
	}

	brick.z1 = 1
	brick.z2 = brick.z1 + height
	return brick, len(bricks) - 1
}

func supported_by(bricks []Brick, i int) []int {
	supported := []int{}
	brick := bricks[i]
	for k := 0; k < len(bricks); k++ {
		if k != i {
			if bricks[k].z2+1 == brick.z1 {
				if intersect_xy(bricks[k], brick) {
					supported = append(supported, k)
				}
			}
		}
	}
	return supported
}

func compute_reaction(supports map[int][]int, supported_by_ map[int][]int, destroyed map[int]bool, start int) int {
	if len(supports[start]) == 0 {
		return 0
	}

	destroyed[start] = true

	res := 0
	for _, s := range supports[start] {
		all_destroyed := true
		for _, m := range supported_by_[s] {
			if !destroyed[m] {
				all_destroyed = false
			}
		}
		if all_destroyed {
			res += compute_reaction(supports, supported_by_, destroyed, s) + 1
		}
	}

	return res
}

func count_desintegrables(bricks []Brick) (int, int) {
	is_important := make(map[int]bool)
	supports := make(map[int][]int)
	supported_by_ := make(map[int][]int)

	for i := range bricks {
		supported := supported_by(bricks, i)
		supported_by_[i] = supported
		if len(supported) == 1 {
			is_important[supported[0]] = true
		}
		for _, s := range supported {
			supports[s] = append(supports[s], i)
		}
	}

	chain_reaction := 0
	for key := range is_important {
		destroyed := make(map[int]bool)
		chain_reaction += compute_reaction(supports, supported_by_, destroyed, key)
	}

	return len(bricks) - len(is_important), chain_reaction
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

	bricks := []Brick{}
	for _, line := range words {
		line_s := strings.Split(line, "~")
		e1 := strings.Split(line_s[0], ",")
		e2 := strings.Split(line_s[1], ",")
		x1, _ := strconv.Atoi(e1[0])
		y1, _ := strconv.Atoi(e1[1])
		z1, _ := strconv.Atoi(e1[2])
		x2, _ := strconv.Atoi(e2[0])
		y2, _ := strconv.Atoi(e2[1])
		z2, _ := strconv.Atoi(e2[2])
		bricks = append(bricks, Brick{x1, y1, z1, x2, y2, z2, ""})
	}

	// Sort by z1
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].z1 > bricks[j].z1
	})

	for i := len(bricks) - 1; i >= 0; i-- {
		new_brick, new_i := move_down(bricks, i)

		bricks = append(bricks[:i], bricks[i+1:]...)
		bricks = append(bricks[:new_i+1], bricks[new_i:]...)
		bricks[new_i] = new_brick
	}

	solution1, solution2 := count_desintegrables(bricks)
	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
