package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_visible_right(heights []int, w int, h int, i int, j int) bool {
	for k := j + 1; k < w; k++ {
		if heights[w * i + k] >= heights[w * i + j] {
			return false
		}
	}
	return true
}

func is_visible_left(heights []int, w int, h int, i int, j int) bool {
	for k := j - 1; k >= 0; k-- {
		if heights[w * i + k] >= heights[w * i + j] {
			return false
		}
	}
	return true
}

func is_visible_top(heights []int, w int, h int, i int, j int) bool {
	for k := i - 1; k >= 0; k-- {
		if heights[w * k + j] >= heights[w * i + j] {
			return false
		}
	}
	return true
}

func is_visible_down(heights []int, w int, h int, i int, j int) bool {
	for k := i + 1; k < h; k++ {
		if heights[w * k + j] >= heights[w * i + j] {
			return false
		}
	}
	return true
}

func is_visible(heights []int, w int, h int, i int, j int) bool {
	return is_visible_right(heights, w, h, i, j) ||
		is_visible_left(heights, w, h, i, j) ||
		is_visible_top(heights, w, h, i, j) ||
		is_visible_down(heights, w, h, i, j)
}

func count_visibles(heights []int, w int, h int) int {
	count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if is_visible(heights, w, h, i, j) {
				count += 1
			}
		}
	}
	return count
}

func visible_distance_right(heights []int, w int, h int, i int, j int) int {
	k := j + 1
	for ; k < w && heights[w * i + k] < heights[w * i + j]; k++ {}
	if k == w {
		return k - j - 1
	} else {
		return k - j
	}
}

func visible_distance_left(heights []int, w int, h int, i int, j int) int {
	k := j - 1
	for ; k >= 0 && heights[w * i + k] < heights[w * i + j]; k-- {}
	if k == -1 {
		return j
	} else {
		return j - k
	}
}

func visible_distance_top(heights []int, w int, h int, i int, j int) int {
	k := i - 1
	for ; k >= 0 && heights[w * k + j] < heights[w * i + j]; k-- {}
	if k == -1 {
		return i
	} else {
		return i - k
	}
}

func visible_distance_down(heights []int, w int, h int, i int, j int) int {
	k := i + 1
	for ; k < h && heights[w * k + j] < heights[w * i + j]; k++ {}
	if k == h {
		return k - i - 1
	} else {
		return k - i
	}
}

func compute_max_visibility(heights []int, w int, h int) int {
	max_visibility := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			product :=  visible_distance_top(heights, w, h, i, j) * 
			visible_distance_down(heights, w, h, i, j) * 
			visible_distance_right(heights, w, h, i, j) * 
			visible_distance_left(heights, w, h, i, j)

			if product > max_visibility {
				max_visibility = product
			}
		}
	}
	return max_visibility
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

	h := len(words)
	w := len(words[0])
	heights := make([]int, h * w)
	for i, word := range words {
		for j, val := range word {
			ival, _ := strconv.Atoi(string(val))
			heights[i * w + j] = ival
		}
	}

	fmt.Println("solution 1:", count_visibles(heights, w, h))
	fmt.Println("solution 2:", compute_max_visibility(heights, w, h))
}
