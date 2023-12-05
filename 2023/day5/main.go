package main

import (
	"fmt"
	"os"
	"strconv"

	// "strconv"
	"strings"
)

type Converter struct {
	SrcName  string
	DestName string
	Destinations  []int
	Sources       []int
	Ranges        []int
}

func parse_converter(str string) Converter {
	lines := strings.Split(strings.Trim(str, "\n"), "\n")
	names := strings.Split(strings.Split(lines[0], " ")[0], "-")
	src_name := names[0]
	dest_name := names[2]

	sources := []int{}
	destinations := []int{}
	ranges := []int{}

	for i := 1; i < len(lines); i++ {
		values := strings.Split(lines[i], " ")
		v0, _ := strconv.Atoi(values[0])
		v1, _ := strconv.Atoi(values[1])
		v2, _ := strconv.Atoi(values[2])
		destinations = append(destinations, v0)
		sources = append(sources, v1)
		ranges = append(ranges, v2)
	}

	return Converter {
		SrcName: src_name,
		DestName: dest_name,
		Destinations: destinations,
		Sources: sources,
		Ranges: ranges,
	}
}

type convert func(converters []Converter, src string, seed int) (int, string)

func convert_1(converters []Converter, src string, seed int) (int, string) {
	var dest_name string
	for _, converter := range converters {
		if converter.SrcName == src {
			dest_name = converter.DestName
			for i := 0; i < len(converter.Sources); i++ {
				if converter.Sources[i] <= seed && seed <= converter.Sources[i] + converter.Ranges[i] {
					return converter.Destinations[i] + seed - converter.Sources[i], dest_name
				}
			}
		}
	}
	return seed, dest_name
}

func convert_2(converters []Converter, src string, seed int) (int, string) {
	var src_name string
	for _, converter := range converters {
		if converter.DestName == src {
			src_name = converter.SrcName
			for i := 0; i < len(converter.Destinations); i++ {
				if converter.Destinations[i] <= seed && seed <= converter.Destinations[i] + converter.Ranges[i] {
					return converter.Sources[i] + seed - converter.Destinations[i], src_name
				}
			}
		}
	}
	return seed, src_name
}

func find_dest(converters []Converter, source string, dest string, seed int, conv convert) int {
	var new_dest string
	val := seed
	for {
		val, new_dest = conv(converters, source, val)
		if new_dest == dest {
			break
		}
		source = new_dest
	}
	return val
}

func belongs_to_seeds(seeds []int, seed int) bool {
	for i := 0; i < len(seeds); i+=2 {
		if seeds[i] <= seed && seed <= seeds[i] + seeds[i+1] {
			return true
		}
	}
	return false
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

	words := strings.Split(strings.Trim(string(dat), "\n"), "\n\n")

	seeds_s := strings.Split(words[0][7:], " ")
	seeds := []int{}
	for _, s := range seeds_s {
		val, _ := strconv.Atoi(s)
		seeds = append(seeds, val)
	}

	var converters []Converter
	for i := 1; i < len(words); i++ {
		converters = append(converters, parse_converter(words[i]))
	}

	var min_location int
	for i, seed := range seeds {
		val := find_dest(converters, "seed", "location", seed, convert_1)
		if i == 0{
			min_location = val
		} else {
			min_location = min(min_location, val)
		}
	}

	fmt.Println("solution 1:", min_location)

	// If I had a chance to find a min location (due to the high seed spectrum) is by going from back to front
	var min_dest_index int
	for _, converter := range converters {
		if converter.DestName == "location" {
			for i, dest_val := range converter.Destinations {
				if i == 0 {
					min_dest_index = 0
				} else {
					if dest_val < converter.Destinations[min_dest_index] {
						min_dest_index = i
					}
				}
			}
		}
	}

	solution2 := -1
	r := converters[len(converters)-1].Ranges[min_dest_index]
	start := converters[len(converters)-1].Destinations[min_dest_index]
	for i := start; i < start + r; i++ {
		dest := find_dest(converters, "location", "seed", i, convert_2)
		if belongs_to_seeds(seeds, dest) {
			solution2 = i
			break
		}
	}

	// This strategy is for inputa.dat only
	//for i := 0; i < len(seeds); i+=2 {
	//	for j := 0; j < seeds[i+1]; j++ {
	//		seed := seeds[i] + j
	//		val := find_dest(converters, "seed", "location", seed, convert_1)
	//		if i == 0 && j == 0 {
	//			min_location = val
	//		} else {
	//			min_location = min(min_location, val)
	//		}

	//	}
	//}
	//solution2 = min_location

	fmt.Println("solution 2:", solution2)
}
