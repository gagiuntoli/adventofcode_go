package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"container/heap"
)

type Direction uint8

const (
	Up    = 1
	Down  = 2
	Left  = 3
	Right = 4
)

var AllDirs = []Direction{Up, Down, Left, Right}

type State struct {
	i, j int
	dir Direction
	count int
}

type Item struct {
   state    State
   priority int    
   index    int   
}

type PriorityQueue []Item

func (piq PriorityQueue) Len() int {
	return len(piq)
}

func (piq PriorityQueue) Less(i, j int) bool {
	return piq[i].priority < piq[j].priority
}

func (piq PriorityQueue) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]
	piq[i].index = i
	piq[j].index = j
}

func (piq *PriorityQueue) Push(x interface{}) {
	n := len(*piq)
	item := x.(Item)
	item.index = n
	*piq = append(*piq, item)
}

func (piq *PriorityQueue) Pop() interface{} {
	old := *piq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*piq = old[0 : n-1]
	return item
}

func (piq *PriorityQueue) Update(item *Item, state State, priority int) {
	item.state = state
	item.priority = priority
	heap.Fix(piq, item.index)
}

func get_next_coord(i, j int, dir Direction) (int,int) {
	if dir == Up {
		return i - 1, j
	} else if dir == Down {
		return i + 1, j
	} else if dir == Right {
		return i, j + 1
	} else if dir == Left {
		return i, j - 1
	}
	panic("not reachable")
}

func get_min_heat(heat_map map[State]int, i, j int, m [][]int, min_movs, max_movs int) int {
	min_heat := 10000000
	for _, dir := range AllDirs {
		for movs := min_movs; movs <= max_movs; movs++ {
			if heat_map[State{i, j, dir, movs}] != 0 {
				min_heat = min(min_heat, heat_map[State{i, j, dir, movs}])
			}
		}
	}
	return min_heat
}


func minimal_heat_loss(m [][]int, min_movs, max_movs int) int {
	heat_map := make(map[State]int)

	s1 := State{0, 0, Right, 0}
	s2 := State{0, 0, Down, 0}
	queue := make(PriorityQueue, 0)
	heat_map[s1] = 0
	heat_map[s2] = 0
	heap.Push(&queue, Item{state: s1, priority: heat_map[s1]})
	heap.Push(&queue, Item{state: s2, priority: heat_map[s2]})

	for queue.Len() > 0 {
		val := heap.Pop(&queue).(Item)

		current := val.state

		dc := current.dir
		cc := current.count

		for _, dir := range AllDirs {
			move_cost := 1
			if dir == dc {
				move_cost += cc
			} 

			heat := heat_map[current]

			i := current.i
			j := current.j

			for ; move_cost <= max_movs; move_cost++ {
				in, jn := get_next_coord(i, j, dir)
				if in < 0 || jn < 0 || in >= len(m) || jn >= len(m[0]) {
					break
				}
				if dir != dc && cc < min_movs {
					continue
				}
				heat += m[in][jn]

				state := State{in, jn, dir, move_cost}

				if heat_map[state] == 0 || heat < heat_map[state] {
					heat_map[state] = heat
					heap.Push(&queue, Item{state: state, priority: heat})
				}

				i = in
				j = jn
			}
		}

	}

	return get_min_heat(heat_map, len(m)-1, len(m[0])-1, m, min_movs, max_movs)
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

	m := make([][]int, len(words))

	for i, word := range words {
		arr := make([]int, len(words[0]))
		for j, val := range word {
			num, _ := strconv.Atoi(string(val))
			arr[j] = num
		}
		m[i] = arr
	}

	solution1 := minimal_heat_loss(m, 1, 3)
	fmt.Println("solution 1:", solution1)

	solution2 := minimal_heat_loss(m, 4, 10)
	fmt.Println("solution 2:", solution2)
}
