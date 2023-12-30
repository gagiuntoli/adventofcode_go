package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Vertex struct {
	i, j int
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

func is_node(grid []string, vertex Vertex) bool {
	i := vertex.i
	j := vertex.j

	count := 0
	//down
	if i < len(grid)-1 && grid[i+1][j] != '#' {
		count++
	}
	//right
	if j < len(grid[0])-1 && grid[i][j+1] != '#' {
		count++
	}
	//up
	if i > 0 && grid[i-1][j] != '#' {
		count++
	}
	//left
	if j > 0 && grid[i][j-1] != '#' {
		count++
	}

	return count > 2
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

		found := false
		for _, adj := range adjs {
			if !visited[adj] {
				current = adj
				found = true
				break
			}
		}
		if !found {
			return Vertex{0, 0}, -1
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
		if distance > 0 {
			connections = append(connections, next_node)
			distances = append(distances, distance)
		}
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

		if is_node(grid, current) {
			nodes = append(nodes, current)
		}

		for _, v := range adjs {
			if !visited[v] {
				queue = append(queue, v)
			}
		}
	}

	nodes = append(nodes, end)

	//----------------
	// Check connections
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

	start := Vertex{0, 1}
	end := Vertex{len(grid) - 1, len(grid[0]) - 2}
	connections, distances := compute_graph(grid, start, end)

	visited := make(map[Vertex]bool)
	solution1, _ := max_distance_graph(connections, distances, visited, start, end)

	new_grid := []string{}
	for _, row := range grid {
		for j := range row {
			if row[j] != '#' && row[j] != '.' {
				row = row[:j] + "." + row[j+1:]
			}
		}
		new_grid = append(new_grid, row)
	}

	connections, distances = compute_graph(new_grid, start, end)

	visited = make(map[Vertex]bool)
	solution2, _ := max_distance_graph(connections, distances, visited, start, end)

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
