package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	utils "github.com/gagiuntoli/adventofcode_go/utils"
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

func count_visible_trees(heights []int, w int, h int) int {
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

func count_visible_trees_part(wg *sync.WaitGroup, heights []int, w int, h int, w1 int, w2 int, h1 int, h2 int, result *int) {
	defer wg.Wait()
	count := 0
	for i := h1; i < h2; i++ {
		for j := w1; j < w2; j++ {
			if is_visible(heights, w, h, i, j) {
				count += 1
			}
		}
	}
	*result = count
}


/* This function was developed to test the Go concurrency only */
func compute_max_visibility_part(wg *sync.WaitGroup, heights []int, w int, h int, w1 int, w2 int, h1 int, h2 int, result *int) {

	defer wg.Done()
	max_visibility := 0
	for i := h1; i < h2; i++ {
		for j := w1; j < w2; j++ {
			product := visible_distance_top(heights, w, h, i, j) *
				visible_distance_down(heights, w, h, i, j) *
				visible_distance_right(heights, w, h, i, j) *
				visible_distance_left(heights, w, h, i, j)

			if product > max_visibility {
				max_visibility = product
			}
		}
	}
	*result = max_visibility
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

	fmt.Println("solution 1:", count_visible_trees(heights, w, h))
	fmt.Println("solution 2:", compute_max_visibility(heights, w, h))

	var wg sync.WaitGroup
	jobs := 8

	results1 := make([]int, jobs)

	results2 := make([]int, jobs)
	for i := range results2 {
		wg.Add(1)
		h1 := 0
		h2 := h
		w1 := i * w / jobs
		w2 := (i + 1) * w / jobs
		go count_visible_trees_part(&wg, heights, w, h, w1, w2, h1, h2, &results1[i])
		go compute_max_visibility_part(&wg, heights, w, h, w1, w2, h1, h2, &results2[i])
	}
	wg.Wait()

	solution1 := utils.ArraySum(results1)
	fmt.Println("solution 1 (parallel):", solution1)

	solution2 := utils.ArrayMax(results2)
	fmt.Println("solution 2 (parallel):", solution2)
}
