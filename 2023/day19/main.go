package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x, m, a, s int
}

type Rule struct {
	part string
	relation string
	value int
	result string
}

func part_sum(p Part) int {
	return p.x + p.m + p.a + p.s
}

func get_result(rule Rule, part Part) string {
	if rule.part == "" {
		return rule.result
	} else {
		var val int
		if rule.part == "x" {
			val = part.x
		} else if rule.part == "m" {
			val = part.m
		} else if rule.part == "a" {
			val = part.a
		} else if rule.part == "s" {
			val = part.s
		}
		if rule.relation == ">" {
			if val > rule.value {
				return rule.result
			} else {
				return "NEXT"
			}
		} else if rule.relation == "<" {
			if val < rule.value {
				return rule.result
			} else {
				return "NEXT"
			}
		}
	}
	panic("false rule or part values")
}

func is_acepted(rulesMap map[string][]Rule, part Part) bool {
	rules := rulesMap["in"]
	index := 0
	for {
		result := get_result(rules[index], part)
		if result == "A" {
			return true
		} else if result == "R" {
			return false
		} else if result == "NEXT" {
			index++
		} else {
			rules = rulesMap[result]
			index = 0
		}
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

	words := strings.Split(strings.Trim(string(dat), "\n"), "\n\n")
	rules_str := strings.Split(words[0], "\n")
	parts_str := strings.Split(words[1], "\n")

	rules := make(map[string][]Rule)
	parts := []Part{}

	for _, part := range parts_str {
		part = string(strings.ReplaceAll(part, "{", " "))
		part = string(strings.ReplaceAll(part, "}", " "))
		part_spl := strings.Split(part, ",")
		x, _ := strconv.Atoi(strings.Trim(part_spl[0], " ")[2:])
		m, _ := strconv.Atoi(strings.Trim(part_spl[1], " ")[2:])
		a, _ := strconv.Atoi(strings.Trim(part_spl[2], " ")[2:])
		s, _ := strconv.Atoi(strings.Trim(part_spl[3], " ")[2:])
		parts = append(parts, Part{x, m, a, s})
	}

	for _, rule := range rules_str {
		rule = string(strings.ReplaceAll(rule, "{", " "))
		rule = string(strings.ReplaceAll(rule, "}", " "))
		rule_spl := strings.Split(rule, " ")
		name := strings.Trim(rule_spl[0], " ")
		rule_spl2 := strings.Split(rule_spl[1], ",")
		rule_arr := []Rule{}

		for _, rule := range rule_spl2 {
			rule = strings.Trim(rule, " ")
			index := strings.Index(rule, ":")
			if index >= 0 {
				part := rule[:1]
				value, _ := strconv.Atoi(rule[2:index])
				result := rule[index+1:]
				var relation string
				if strings.Contains(rule, ">") {
					relation = ">"
				} else if strings.Contains(rule, "<") {
					relation = "<"
				} else {
					panic("invalid rule")
				}
				rule_arr = append(rule_arr, Rule{part, relation, value, result})

			} else {
				rule_arr = append(rule_arr, Rule{part: "", relation: "", value: -1, result: rule})
			}
		}

		rules[name] = rule_arr
	}

	solution1 := 0
	for _, part := range parts {
		if is_acepted(rules, part) {
			solution1 += part_sum(part)
		}
	}

	solution2 := 0
	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
