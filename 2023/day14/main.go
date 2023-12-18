package main

import (
	"fmt"
	"os"
	"strings"
)

func roll_north(m []string) []string {
	m_new := m
	for j := 0; j < len(m[0]); j++ {
		for i := 1; i < len(m); i++ {
			if m_new[i][j] == 'O' {
				row := i - 1
				for ; row >= 0 && m_new[row][j] == '.'; row-- {
				}

				if row+1 != i {
					m_new[row+1] = m_new[row+1][:j] + "O" + m_new[row+1][j+1:]
					m_new[i] = m_new[i][:j] + "." + m_new[i][j+1:]
				}
			}
		}
	}
	return m_new
}

func roll_south(m []string) []string {
	m_new := m
	for j := 0; j < len(m[0]); j++ {
		for i := len(m) - 2; i >= 0; i-- {
			if m_new[i][j] == 'O' {
				row := i + 1
				for ; row < len(m) && m_new[row][j] == '.'; row++ {
				}

				if row-1 != i {
					m_new[row-1] = m_new[row-1][:j] + "O" + m_new[row-1][j+1:]
					m_new[i] = m_new[i][:j] + "." + m_new[i][j+1:]
				}
			}
		}
	}
	return m_new
}

func roll_east(m []string) []string {
	m_new := m
	for i := 0; i < len(m); i++ {
		for j := len(m[0]) - 2; j >= 0; j-- {
			if m_new[i][j] == 'O' {
				col := j + 1
				for ; col < len(m[0]) && m_new[i][col] == '.'; col++ {
				}

				if col-1 != j {
					m_new[i] = m_new[i][:j] + "." + m_new[i][j+1:col-1] + "O" + m_new[i][col:]
				}
			}
		}
	}
	return m_new
}

func roll_west(m []string) []string {
	m_new := m
	for i := 0; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {
			if m_new[i][j] == 'O' {
				col := j - 1
				for ; col >= 0 && m_new[i][col] == '.'; col-- {
				}

				if col+1 != j {
					m_new[i] = m_new[i][:col+1] + "O" + m_new[i][col+2:j] + "." + m_new[i][j+1:]
				}
			}
		}
	}
	return m_new
}

func m2string(m []string) string {
	res := ""
	for _, row := range m {
		res += row
	}
	return res
}

func cycle(m []string, total_cycles uint64) []string {
	cache := make(map[string]uint64)
	start := int64(-1)
	cycle := uint64(1)

	for i := uint64(0); i < total_cycles; i++ {
		mstr := m2string(m)
		index, ok := cache[mstr]
		if start != -1 {
			if index == uint64(start) {
				for j := uint64(0); j < (total_cycles-uint64(start))%cycle; j++ {
					m = one_cycle(m)
				}
				return m

			} else {
				cycle++
			}
		} else {
			if ok {
				start = int64(index)
			} else {
				cache[mstr] = i
			}
		}

		m = one_cycle(m)
	}
	return m
}

func one_cycle(m []string) []string {
	m = roll_north(m)
	m = roll_west(m)
	m = roll_south(m)
	m = roll_east(m)
	return m
}

func compute_force(m []string) int {
	force := 0

	for i := 0; i < len(m); i++ {
		force += (len(m) - i) * strings.Count(m[i], "O")
	}

	return force
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

	mn := roll_north(m)
	solution1 := compute_force(mn)

	m = cycle(m, 1000000000)
	solution2 := compute_force(m)

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
