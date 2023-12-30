package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Direction uint8

const (
	Up    = 1
	Down  = 2
	Left  = 3
	Right = 4
)

type Vertex struct {
	i, j int
}

type Item struct {
	vertex   Vertex
	priority int
	index    int
}

func get_next_nodes(grid []string, vertex Vertex) []Vertex {
	i := vertex.i
	j := vertex.j

	nodes := []Vertex{}
	//down
	if i < len(grid)-1 && (grid[i+1][j] == '.' || grid[i+1][j] == 'v') {
		nodes = append(nodes, Vertex{i + 1, j})
	}
	//right
	if j < len(grid[0])-1 && (grid[i][j+1] == '.' || grid[i][j+1] == '>') {
		nodes = append(nodes, Vertex{i, j + 1})
	}
	//up
	if i > 0 && (grid[i-1][j] == '.' || grid[i-1][j] == '^') {
		nodes = append(nodes, Vertex{i - 1, j})
	}
	//left
	if j > 0 && (grid[i][j-1] == '.' || grid[i][j-1] == '<') {
		nodes = append(nodes, Vertex{i, j - 1})
	}

	return nodes
}

func max_distance(grid []string, visited map[Vertex]bool, from, to Vertex) int {
	visited[from] = true

	if from == to {
		return 0
	}

	nodes := get_next_nodes(grid, from)

	dist := 0
	for _, node := range nodes {
		if !visited[node] {
			new_visited := make(map[Vertex]bool)
			for k, v := range visited {
				new_visited[k] = v
			}
			dist = max(dist, max_distance(grid, new_visited, node, to)+1)
		}
	}
	return dist
}

func max_distance_graph(connections map[Vertex][]Vertex, distances map[[2]Vertex]int, visited map[Vertex]bool, from, to Vertex) (int, bool) {
	visited[from] = true

	if from == to {
		return 0, true
	}

	dist := 0
	reached_destination := false
	for _, node := range connections[from] {
		if !visited[node] {
			new_visited := make(map[Vertex]bool)
			for k, v := range visited {
				new_visited[k] = v
			}
			dist_, reached := max_distance_graph(connections, distances, new_visited, node, to)
			if reached {
				dist = max(dist, dist_+distances[[2]Vertex{from, node}])
				reached_destination = reached
			}
		}
	}

	return dist, reached_destination
}

func BFS(grid []string, v Vertex, nodes []Vertex, visited map[Vertex]bool) (Vertex, int) {
	distance := 1
	current := v

	for {
		visited[current] = true

		if slices.Contains(nodes, current) {
			break
		}

		adjs := get_next_nodes(grid, current)

		if len(adjs) == 0 && len(adjs) > 2 {
			panic("adjs should be <= 2 and > 0")
		}

		if !visited[adjs[0]] {
			current = adjs[0]
		} else if !visited[adjs[1]] {
			current = adjs[1]
		}

		distance++
	}

	return current, distance
}

func get_connections_and_distances(grid []string, nodes []Vertex, node Vertex) ([]Vertex, []int) {
	connections := []Vertex{}
	distances := []int{}
	adjs := get_next_nodes(grid, node)

	for _, v := range adjs {
		// we should find only 1 in this direction
		visited := make(map[Vertex]bool)
		visited[node] = true
		next_node, distance := BFS(grid, v, nodes, visited)
		connections = append(connections, next_node)
		distances = append(distances, distance)
	}

	return connections, distances
}

func compute_graph(grid []string, start, end Vertex) (map[Vertex][]Vertex, map[[2]Vertex]int) {
	visited := make(map[Vertex]bool)
	queue := []Vertex{start}
	nodes := []Vertex{start}

	//----------------
	// Search for nodes

	for len(queue) > 0 {

		current := queue[0]
		visited[current] = true
		queue = queue[1:]

		adjs := get_next_nodes(grid, current)

		if len(adjs) > 2 {
			// it is a node
			nodes = append(nodes, current)
		}

		for _, v := range adjs {
			if !visited[v] {
				queue = append(queue, v)
			}
		}
	}

	nodes = append(nodes, end)

	new_grid := []string{}
	for _, row := range grid {
		new_grid = append(new_grid, row)
	}
	for _, n := range nodes {
		i := n.i
		j := n.j
		new_grid[i] = new_grid[i][:j] + "O" + new_grid[i][j+1:]
	}

	// fmt.Println("NODES:")
	// for _, row := range new_grid {
	// 	fmt.Println(row)
	// }
	// fmt.Println()

	//----------------

	distanceMap := make(map[[2]Vertex]int)     // distance from i to j
	connectionMap := make(map[Vertex][]Vertex) // i connections
	for _, node := range nodes {
		connections, distances := get_connections_and_distances(grid, nodes, node)

		connectionMap[node] = connections

		for i := 0; i < len(connections); i++ {
			distanceMap[[2]Vertex{node, connections[i]}] = distances[i]
			distanceMap[[2]Vertex{connections[i], node}] = distances[i]
		}
	}

	//for k, d := range distanceMap {
	//	fmt.Printf("distance %+v -> %+v\n", k, d)
	//}
	//for k, d := range connectionMap {
	//	fmt.Printf("%+2v -> %+2v\n", k, d)
	//}

	return connectionMap, distanceMap
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

	grid := strings.Split(strings.Trim(string(dat), "\n"), "\n")

	visited := make(map[Vertex]bool)
	start := Vertex{0, 1}
	end := Vertex{len(grid) - 1, len(grid[0]) - 2}
	solution1 := max_distance(grid, visited, start, end)

	new_grid := []string{}
	for _, row := range grid {
		for j := range row {
			if row[j] != '#' && row[j] != '.' {
				row = row[:j] + "." + row[j+1:]
			}
		}
		new_grid = append(new_grid, row)
	}

	connections, distances := compute_graph(new_grid, start, end)

	visited = make(map[Vertex]bool)
	solution2, _ := max_distance_graph(connections, distances, visited, start, end)

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
