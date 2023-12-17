package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func array2string(array []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}

type Key struct {
	str string
	seq string
}

var cache map[Key]int

func possibilities(str string, seq []int) int {
	cacheValue, ok := cache[Key{str, array2string(seq)}]
	if ok {
		return cacheValue
	}

	if len(str) == 0 && len(seq) > 0 {
		cache[Key{str, array2string(seq)}] = 0
		return 0
	}

	if len(str) >= 0 && len(seq) == 0 {
		if strings.Contains(str, "#") {
			cache[Key{str, array2string(seq)}] = 0
			return 0
		} else {
			cache[Key{str, array2string(seq)}] = 1
			return 1
		}
	}

	seq_v := seq[0]
	if seq_v > len(str) {
		cache[Key{str, array2string(seq)}] = 0
		return 0
	}

	if str[0] == '?' {
		str1 := "#" + str[1:]
		str2 := str[1:]

		count1 := possibilities(str1, seq)
		count2 := possibilities(str2, seq)

		returnValue := count1 + count2
		cache[Key{str, array2string(seq)}] = returnValue

		return returnValue

	} else if str[0] == '#' {
		// check if it is long enought for the first group
		if !strings.Contains(str[:seq_v], ".") {
			if seq_v == len(str) {
				cache[Key{str, array2string(seq)}] = 1
				return 1
			} else if str[seq_v] != '#' {
				return possibilities(str[seq_v+1:], seq[1:])
			} else {
				cache[Key{str, array2string(seq)}] = 0
				return 0
			}
		} else {
			cache[Key{str, array2string(seq)}] = 0
			return 0
		}
	}

	// it is '.'
	return possibilities(str[1:], seq)
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

	solution1 := 0
	solution2 := 0
	for _, word := range words {
		word_s := strings.Split(word, " ")
		str := word_s[0]
		sequence_str := strings.Split(word_s[1], ",")
		sequence := []int{}
		for _, val := range sequence_str {
			num, _ := strconv.Atoi(val)
			sequence = append(sequence, num)
		}

		seq5 := []int{}
		for i := 0; i < 5; i++ {
			for _, v := range sequence {
				seq5 = append(seq5, v)
			}
		}

		cache = make(map[Key]int)
		solution1 += possibilities(str+".", sequence)

		solution2 += possibilities(str+"?"+str+"?"+str+"?"+str+"?"+str+".", seq5)

	}
	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
