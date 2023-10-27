package main

import (
	"fmt"
	"os"
	"strconv"

	"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name           string
	files          []File
	subdirectories []Directory
	parent         *Directory
	size           int
}

func computeDirectorySize(directory Directory) int {
	size := 0
	for _, file := range directory.files {
		size += file.size
	}
	for _, sub_directory := range directory.subdirectories {
		size += computeDirectorySize(sub_directory)
	}
	return size
}

func compute_directory_sizes(dir Directory, dir_sizes map[string]int, path string) {
	dir_sizes[path] = computeDirectorySize(dir)
	for _, sub_directory := range dir.subdirectories {
		compute_directory_sizes(sub_directory, dir_sizes, path+"/"+sub_directory.name)
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

	words := strings.Split(string(dat), "\n")

	root := Directory{name: "root"}
	current := &root
	for _, word := range words {
		if len(word) > 0 {
			parts := strings.Split(word, " ")
			if parts[0] == "$" {
				if parts[1] == "cd" {
					if parts[2] == "/" {
						current = &root
					} else if parts[2] == ".." {
						current = current.parent
					} else {
						for i, dir := range current.subdirectories {
							if dir.name == parts[2] {
								current = &current.subdirectories[i]
							}
						}
					}
				}
			} else if parts[0] == "dir" {
				new_directory := Directory{name: parts[1], parent: current}
				current.subdirectories = append(current.subdirectories, new_directory)
			} else {
				size, _ := strconv.Atoi(parts[0])
				file := File{name: parts[1], size: size}
				current.files = append(current.files, file)
			}
		}
	}

	dir_sizes := map[string]int{}
	compute_directory_sizes(root, dir_sizes, "/")

	solution1 := 0
	for _, size := range dir_sizes {
		if size < 100000 {
			solution1 += size
		}
	}

	smallest_size := 100000000
	for _, size := range dir_sizes {
		if 70000000-dir_sizes["/"]+size > 30000000 && size < smallest_size {
			smallest_size = size
		}
	}

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", smallest_size)
}
