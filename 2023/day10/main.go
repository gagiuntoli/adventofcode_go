package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	utils "github.com/gagiuntoli/adventofcode_go/utils"
)

type Coordinate struct {
	i, j int
}

const (
	Up = 0
	Down = 1
	Left = 2
	Right = 3
)

func find_S(map_ []string) (int, int) {
	for i, line := range map_ {
		for j := range line {
			if line[j] == 'S' {
				return i, j
			}
		}
	}
	panic("S not found")
}

func find_next_position(map_ []string, visited map[Coordinate]bool, i, j int) (int, int, int) {
	type Movement struct {
		Coordinate
		Direction int
		PossibleConnections []byte
	}

	movementsJ := []Movement{
		{Coordinate{i - 1, j}, Up, []byte{'|', 'F', '7', 'S'}},
		{Coordinate{i, j - 1}, Left, []byte{'-', 'L', 'F', 'S'}},
	}

	movements7 := []Movement{
		{Coordinate{i + 1, j}, Down, []byte{'|', 'L', 'J', 'S'}},
		{Coordinate{i, j - 1}, Left, []byte{'-', 'L', 'F', 'S'}},
	}

	movementsL := []Movement{
		{Coordinate{i - 1, j}, Up, []byte{'|', 'F', '7', 'S'}},
		{Coordinate{i, j + 1}, Right, []byte{'-', '7', 'J', 'S'}},
	}

	movementsF := []Movement{
		{Coordinate{i + 1, j}, Down, []byte{'|', 'L', 'J', 'S'}},
		{Coordinate{i, j + 1}, Right, []byte{'-', '7', 'J', 'S'}},
	}

	movementsH := []Movement{
		{Coordinate{i, j + 1}, Right, []byte{'-', '7', 'J', 'S'}},
		{Coordinate{i, j - 1}, Left, []byte{'-', 'L', 'F', 'S'}},
	}

	movementsV := []Movement{
		{Coordinate{i - 1, j}, Up, []byte{'|', 'F', '7', 'S'}},
		{Coordinate{i + 1, j}, Down, []byte{'|', 'L', 'J', 'S'}},
	}

	movementsS := []Movement{
		{Coordinate{i - 1, j}, Up, []byte{'|', 'F', '7', 'S'}},
		{Coordinate{i + 1, j}, Down, []byte{'|', 'L', 'J', 'S'}},
		{Coordinate{i, j + 1}, Right, []byte{'-', '7', 'J', 'S'}},
		{Coordinate{i, j - 1}, Left, []byte{'-', 'L', 'F', 'S'}},
	}
	
	var movements []Movement

	if map_[i][j] == 'J' {
		movements = movementsJ
	} else if map_[i][j] == '7' {
		movements = movements7
	} else if map_[i][j] == 'L' {
		movements = movementsL
	} else if map_[i][j] == 'F' {
		movements = movementsF
	} else if map_[i][j] == '-' {
		movements = movementsH
	} else if map_[i][j] == '|' {
		movements = movementsV
	} else if map_[i][j] == 'S' {
		movements = movementsS
	}

	for _, movement := range movements {
		coordinate := movement.Coordinate
		in := coordinate.i
		jn := coordinate.j
		
		if in >= 0 && in < len(map_) && jn >= 0 && jn < len(map_[0]) && !visited[Coordinate{in, jn}] {
			if slices.Contains(movement.PossibleConnections, map_[in][jn]) {
				return in, jn, movement.Direction
			}
		}
	}

	for _, movement := range movements {
		coordinate := movement.Coordinate
		in := coordinate.i
		jn := coordinate.j
		if in >= 0 && in < len(map_) && jn >= 0 && jn < len(map_[0]) {
			if map_[in][jn] == 'S' {
				return in, jn, movement.Direction
			}
		}
	}
	panic("could not find next movement")
}

func mark_direction(map_ []string, i, j int, directionMap map[Coordinate][4]int, direction int) {
	width := len(map_[0])
	height := len(map_)

	if direction == Up {
		for k := j - 1; k >= 0; k-- {
			arr := directionMap[Coordinate{i, k}]
			arr[Up]--
		directionMap[Coordinate{i, k}] = arr
			if map_[i][k] != '.' {
				break
			}
		}
		for k := j + 1; k < width; k++ {
			arr := directionMap[Coordinate{i, k}]
			arr[Up]++
		directionMap[Coordinate{i, k}] = arr
			if map_[i][k] != '.' {
				break
			}
		}
	} else if direction == Down {
		for k := j - 1; k >= 0; k-- {
			arr := directionMap[Coordinate{i, k}]
			arr[Down]++
		directionMap[Coordinate{i, k}] = arr
			if map_[i][k] != '.' {
				break
			}
		}
		for k := j + 1; k < width; k++ {
			arr := directionMap[Coordinate{i, k}]
			arr[Down]--
		directionMap[Coordinate{i, k}] = arr
			if map_[i][k] != '.' {
				break
			}
		}
	} else if direction == Right {
		for k := i + 1; k < height; k++ {
			arr := directionMap[Coordinate{k, j}]
			arr[Right]++
		directionMap[Coordinate{k, j}] = arr
			if map_[k][j] != '.' {
				break
			}
		}
		for k := i - 1; k >= 0; k-- {
			arr := directionMap[Coordinate{k, j}]
			arr[Right]--
		directionMap[Coordinate{k, j}] = arr
			if map_[k][j] != '.' {
				break
			}
		}
	} else if direction == Left {
		for k := i + 1; k < height; k++ {
			arr := directionMap[Coordinate{k, j}]
			arr[Left]--
		directionMap[Coordinate{k, j}] = arr
			if map_[k][j] != '.' {
				break
			}
		}
		for k := i - 1; k >= 0; k-- {
			arr := directionMap[Coordinate{k, j}]
			arr[Left]++
		directionMap[Coordinate{k, j}] = arr
			if map_[k][j] != '.' {
				break
			}
		}
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

	map_ := strings.Split(strings.Trim(string(dat), "\n"), "\n")
	clean_map := make([]string, len(map_))
	for i := range map_ {
		clean_map[i] = strings.Repeat(".", len(map_[0]))
	}

	directionMap := make(map[Coordinate][4]int)
	for i := 0; i < len(map_); i++ {
		for j := 0; j < len(map_[0]); j++ {
			directionMap[Coordinate{i, j}] = [4]int{0, 0, 0, 0}
		}
	}

	solution1 := 0
	visited := make(map[Coordinate]bool)
	is, js := find_S(map_)
	i, j := is, js
	var direction, prev_direction, prev_i, prev_j int
	for steps := 0; ; steps++ {
		visited[Coordinate{i, j}] = true
		clean_map[i] = clean_map[i][:j] + string(map_[i][j]) + clean_map[i][j+1:]
		i, j, direction = find_next_position(map_, visited, i, j)
		if map_[i][j] == 'S' {
			solution1 = (steps + 1) / 2
			break
		}
		prev_direction = direction
		prev_i, prev_j = i, j
	}

	visited = make(map[Coordinate]bool)
	i, j = is, js
	for steps := 0; ; steps++ {
		visited[Coordinate{i, j}] = true
		i, j, direction = find_next_position(map_, visited, i, j)
		if steps == 0 {
			prev_direction = direction
			prev_i, prev_j = i, j
		} else if prev_direction != direction {
			mark_direction(clean_map, prev_i, prev_j, directionMap, direction)
		}
		mark_direction(clean_map, i, j, directionMap, direction)
		if map_[i][j] == 'S' {
			break
		}
		prev_direction = direction
		prev_i, prev_j = i, j
	}

	clockwise := 0
	anticlockwise := 0
	for i := range(map_) {
		for j := range(map_[0]) {
			if clean_map[i][j] == '.' {
				arr := directionMap[Coordinate{i,j}]
				if utils.AllEqual(arr[:]) {
					if arr[0] > 0 {
						clockwise += 1
					} else if arr[0] < 0 {
						anticlockwise += 1
					}
				}
			}
		}
	}

	// For some cases could be the clockwise and for other the anticlowise
	// fmt.Println(clockwise)
	// fmt.Println(anticlockwise)
	solution2 := anticlockwise

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
