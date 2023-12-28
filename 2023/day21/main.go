package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gagiuntoli/adventofcode_go/utils"
)

type Point struct {
	i, j int
}

func get_adjs(p Point, grid []string) []Point {
	i := p.i
	j := p.j

	res := []Point{}

	if i > 0 && grid[i-1][j] != '#' && grid[i-1][j] != 'O' {
		res = append(res, Point{i - 1, j})
	}
	if j > 0 && grid[i][j-1] != '#' && grid[i][j-1] != 'O' {
		res = append(res, Point{i, j - 1})
	}
	if i < len(grid)-1 && grid[i+1][j] != '#' && grid[i+1][j] != 'O' {
		res = append(res, Point{i + 1, j})
	}
	if j < len(grid[0])-1 && grid[i][j+1] != '#' && grid[i][j+1] != 'O' {
		res = append(res, Point{i, j + 1})
	}

	return res
}

func get_S_coor(grid []string) (int, int) {
	for i, row := range grid {
		index := strings.Index(row, "S")
		if index >= 0 {
			return i, index
		}
	}
	return -1, -1
}

func evolve_grid_for(grid, ref_grid []string, i, j int) {
	adjs := get_adjs(Point{i, j}, ref_grid)
	for _, p := range adjs {
		grid[p.i] = grid[p.i][:p.j] + "O" + grid[p.i][p.j+1:]
	}
}

func count_points(grid []string) int {
	res := 0
	for _, row := range grid {
		res += strings.Count(row, "O")
	}
	return res
}

func compute_points(grid []string, steps int) int {
	for s := 0; s < steps; s++ {
		old_grid := []string{}
		for _, row := range grid {
			old_grid = append(old_grid, row)
		}

		for i, row := range old_grid {
			for j, val := range row {
				if val == 'O' {
					evolve_grid_for(grid, old_grid, i, j)
					grid[i] = grid[i][:j] + "." + grid[i][j+1:]
				}
			}
		}
	}

	return count_points(grid)
}

func compute_quadratic_coeffients(xs [3]int, ys [3]int) [4]int {
	x1, x2, x3 := xs[0], xs[1], xs[2]
	y1, y2, y3 := ys[0], ys[1], ys[2]

	a1, b1, c1 := x1*x1, x1, 1
	a2, b2, c2 := x2*x2, x2, 1
	a3, b3, c3 := x3*x3, x3, 1

	det := a1*(b2*c3-b3*c2) - b1*(a2*c3-c2*a3) + c1*(a2*b3-a3*b2)
	d1 := y1*(b2*c3-b3*c2) - b1*(y2*c3-c2*y3) + c1*(y2*b3-y3*b2)
	d2 := a1*(y2*c3-y3*c2) - y1*(a2*c3-c2*a3) + c1*(a2*y3-a3*y2)
	d3 := a1*(b2*y3-b3*y2) - b1*(a2*y3-y2*a3) + y1*(a2*b3-a3*b2)

	return [4]int{d1, d2, d3, det}
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

	grid_orig := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	grid := []string{}
	for _, row := range grid_orig {
		grid = append(grid, row)
	}

	is, js := get_S_coor(grid_orig)

	grid[is] = grid[is][:js] + "O" + grid[is][js+1:]
	solution1 := compute_points(grid, 64)
	fmt.Println("solution 1:", solution1)

	W := len(grid)
	grid_orig[is] = grid_orig[is][:js] + "." + grid_orig[is][js+1:]

	grid = []string{}
	reps := 7
	for i := 0; i < reps; i++ {
		for _, row := range grid_orig {
			new_row := ""
			for j := 0; j < reps; j++ {
				new_row += row
			}
			grid = append(grid, new_row)
		}
	}
	grid[reps/2*W+is] = grid[reps/2*W+js][:reps/2*W+js] + "O" + grid[reps/2*W+is][reps/2*W+js+1:]

	y1 := compute_points(grid, W/2)
	y2 := compute_points(grid, W)
	y3 := compute_points(grid, W)

	x1 := W / 2
	x2 := x1 + W
	x3 := x2 + W

	coeffs := compute_quadratic_coeffients([3]int{x1, x2, x3}, [3]int{y1, y2, y3})
	a, b, c, det := uint64(-coeffs[0]), uint64(-coeffs[1]), uint64(-coeffs[2]), uint64(-coeffs[3])

	gdc := utils.GCD(a, det)
	a = a / gdc
	gdc = utils.GCD(b, det)
	b = b / gdc
	gdc = utils.GCD(c, det)
	c = c / gdc
	det = det / gdc

	steps := uint64(26501365)
	solution2 := (a*steps*steps + b*steps + c) / det

	fmt.Println("solution 2:", solution2)
}
