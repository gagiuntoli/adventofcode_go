package main

import (
	"fmt"
	"os"
	"strings"

	utils "github.com/gagiuntoli/adventofcode_go/utils"
)

type Direction int

const (
	Up = 1
	Down = 2
	Left = 3
	Right = 4
)

type Position struct {
	i, j int
}

type Beam struct {
	position Position
	dir  Direction
}

func next_positions(m []string, beam Beam) []Beam {
	dir := beam.dir
	i := beam.position.i
	j := beam.position.j
	square := m[i][j]

	if dir == Right {
		if square == '.' || square == '-' {
			return []Beam{{Position{i,j+1}, dir}}
		} else if square == '\\' {
			return []Beam{{Position{i+1,j}, Down}}
		} else if square == '/' {
			return []Beam{{Position{i-1,j}, Up}}
		} else if square == '|' {
			return []Beam{{Position{i-1,j}, Up}, {Position{i+1,j}, Down}}
		}
	} else if dir == Left {
		if square == '.' || square == '-' {
			return []Beam{{Position{i,j-1}, dir}}
		} else if square == '\\' {
			return []Beam{{Position{i-1,j}, Up}}
		} else if square == '/' {
			return []Beam{{Position{i+1,j}, Down}}
		} else if square == '|' {
			return []Beam{{Position{i-1,j}, Up}, {Position{i+1,j}, Down}}
		}
	} else if dir == Up {
		if square == '.' || square == '|' {
			return []Beam{{Position{i-1,j}, dir}}
		} else if square == '\\' {
			return []Beam{{Position{i,j-1}, Left}}
		} else if square == '/' {
			return []Beam{{Position{i,j+1}, Right}}
		} else if square == '-' {
			return []Beam{{Position{i,j-1}, Left}, {Position{i,j+1}, Right}}
		}
	} else if dir == Down {
		if square == '.' || square == '|' {
			return []Beam{{Position{i+1,j}, dir}}
		} else if square == '\\' {
			return []Beam{{Position{i,j+1}, Right}}
		} else if square == '/' {
			return []Beam{{Position{i,j-1}, Left}}
		} else if square == '-' {
			return []Beam{{Position{i,j-1}, Left}, {Position{i,j+1}, Right}}
		}
	}
	panic("Should not be reached")
}

func get_energizided_cells(m []string, initial_beam Beam) int {
	type Key struct {
		position Position
		dir Direction
	}
	visited := make(map[Key]bool)
	beams := []Beam{initial_beam}

	for ; len(beams) > 0; {
		beams_tmp := []Beam{}
		repeated := []bool{}
		for _, beam := range beams {
			i := beam.position.i
			j := beam.position.j
			dir := beam.dir
			repeated_val := visited[Key{Position{i, j}, dir}]
			if !repeated_val {
				beams_tmp = append(beams_tmp, beam)
			}
			repeated = append(repeated, repeated_val)
			visited[Key{Position{i, j}, dir}] = true
		}
		if utils.All(repeated, true) {
			break
		}
		beams = beams_tmp

		beams_tmp = []Beam{}
		for _, beam := range beams {
			i := beam.position.i
			j := beam.position.j
			if i >= 0 && j >= 0 && i < len(m) && j < len(m[0]) {
			      beams_tmp = append(beams_tmp, beam)
			}
		}
		beams = beams_tmp


		beams_tmp = []Beam{}
		for _, beam := range beams {
			new_beams := next_positions(m, beam)

			beams_tmp = append(beams_tmp, new_beams...)
		}
		beams = beams_tmp
	}


	enegized_cells := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			for _, dir := range []Direction{Up, Down, Right, Left} {
				if visited[Key{Position{i, j}, dir}] {
					enegized_cells += 1
					break
				}
			}
		}
	}
	return enegized_cells
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

	m := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	solution1 := get_energizided_cells(m, Beam{Position{0, 0}, Right})
	fmt.Println("solution 1:", solution1)

	solution2 := 0
	for i := 0; i < len(m); i++ {
		solution2 = max(solution2, get_energizided_cells(m, Beam{Position{i, 0}, Right}))
		solution2 = max(solution2, get_energizided_cells(m, Beam{Position{i, len(m[0])-1}, Left}))
	}
	for j := 0; j < len(m[0]); j++ {
		solution2 = max(solution2, get_energizided_cells(m, Beam{Position{0, j}, Down}))
		solution2 = max(solution2, get_energizided_cells(m, Beam{Position{len(m)-1, j}, Up}))
	}
	fmt.Println("solution 2:", solution2)
}
