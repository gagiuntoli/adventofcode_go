package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func count_nodes(connections map[string][]string, visited map[string]bool, start string) int {
	visited[start] = true
	count := 1
	for _, c := range connections[start] {
		if !visited[c] {
			count += count_nodes(connections, visited, c)	
		}
	}
	return count
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

	connections := make(map[string][]string)
	for _, line := range words {
		line_s := strings.Split(line, ": ")
		node := line_s[0]
		conns := strings.Split(strings.Trim(line_s[1], " "), " ")
		connections[node] = conns
	}

	// Nodes found with graphviz
	pairs := [][]string{
		{"lxb", "vcq"}, {"vcq", "lxb"},
		{"rnx", "ddj"}, {"ddj", "rnx"},
		{"mmr", "znk"}, {"znk", "mmr"},
	}

	keys := []string{}
	for k := range connections {
		keys = append(keys, k)
	}
	
	for _, node := range keys {
		for _, c := range connections[node] {
			conns := connections[c]
			if !slices.Contains(conns, node) {
				conns = append(conns, node)
			}
			connections[c] = conns
		}
	}
	
	keys = []string{}
	for k := range connections {
		keys = append(keys, k)
	}

	for _, key := range keys {
		conns := connections[key]
		for _, pair := range pairs {
			n1 := pair[0]
			n2 := pair[1]
			if key == n1 {
				index := slices.Index(conns, n2)
				if index > 0 {
					conns = append(conns[:index], conns[index+1:]...)
				}
			}
		}
		connections[key] = conns
	}

	visited := make(map[string]bool)
	start1 := "nnx"
	cluster1 := count_nodes(connections, visited, start1)

	start2 := "mck"
	cluster2 := count_nodes(connections, visited, start2)

	solution1 := cluster1 * cluster2

	solution2 := 0
	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
