package main

import (
	"fmt"
	"os"
	"strings"
)

func check_horizontal_reflections(m []string) []int {
	rows := []int{}
	for row := 0; row < len(m)-1; row++ {
		found := true
		for i := 0; row+1+i < len(m) && row-i >= 0; i++ {
			if m[row+1+i] != m[row-i] {
				found = false
				break
			}
		}
		if found {
			rows = append(rows, row+1)
		}
	}
	return rows
}

func reverse_string(str string) string {
	strn := ""
	for i := range str {
		strn += string(str[len(str)-1-i])
	}
	return strn
}

func check_vertical_reflections(m []string) []int {
	cols := []int{}
	for col := 0; col < len(m[0])-1; col++ {
		width := min(len(m[0])-1-col, col+1)
		found := true
		for i := 0; i < len(m); i++ {
			if reverse_string(m[i][col-width+1:col+1]) != m[i][col+1:col+width+1] {
				found = false
				break
			}
		}
		if found {
			cols = append(cols, col+1)
		}
	}
	return cols
}

func find_reflections_1(m []string) (int, int) {
	rows := check_horizontal_reflections(m)
	if len(rows) > 0 {
		return rows[0], 0
	}
	cols := check_vertical_reflections(m)
	if len(cols) > 0 {
		return 0, cols[0]
	}
	return 0, 0
}

func swap(symb string) string {
	if symb == "#" {
		return "."
	} else {
		return "#"
	}
}

func find_reflections_2(m []string) (int, int) {
	row_orig, col_orig := find_reflections_1(m)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			line := m[i]
			symb := string(line[j])
			m[i] = line[:j] + swap(symb) + line[j+1:]
			rows := check_horizontal_reflections(m)
			cols := check_vertical_reflections(m)
			m[i] = line[:j] + symb + line[j+1:]

			if len(rows) > 0 {
				for _, row := range rows {
					if row != row_orig {
						return row, 0
					}
				}
			}
			if len(cols) > 0 {
				for _, col := range cols {
					if col != col_orig {
						return 0, col
					}
				}
			}
		}
	}
	return 0, 0
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

	words := strings.Split(string(dat), "\n\n")

	solution1 := 0
	solution2 := 0
	for _, word := range words {
		m := strings.Split(strings.Trim(word, "\n"), "\n")

		hor, ver := find_reflections_1(m)
		solution1 += ver + 100*hor

		hor, ver = find_reflections_2(m)
		solution2 += ver + 100*hor
	}

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
