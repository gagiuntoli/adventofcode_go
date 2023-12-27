package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/gagiuntoli/adventofcode_go/utils"
)

type Mod struct {
	connections []string
	state bool // On/High = true | Off/Low = false
	kind int
	input_memory map[string]bool
}

type Pulse struct {
	source string
	dest string
	state bool // High = true | Low = false
}

const (
	Broadcaster = 0
	FlipFlop = 1
	Conjunction = 2
)

func press_button(mods map[string]Mod, ff_test []string) (int, int, []bool) {
	first_connections := mods["broadcaster"].connections

	pulses := []Pulse{}
	low, high := 1, 0
	got_low := make([]bool, len(ff_test))

	// we send low pulses to every broadcast connection
	for _, conn := range first_connections {
		pulses = append(pulses, Pulse{"broadcaster", conn, false})
		low++
	}


	for iter := 1; len(pulses) > 0; iter++ {
		pulse := pulses[0]
		pulses = pulses[1:]

		mod := mods[pulse.dest]

		for i, ff := range ff_test {
			//fmt.Println("testing", ff)
			if pulse.dest == ff && !pulse.state {
				got_low[i] = true	
			}
		}

		if len(mod.connections) == 0 {
		 	continue
		}

		if mod.kind == FlipFlop {
			// FlipFlop modules do nothing if a high pulse arrives
			if pulse.state {
				continue
			}

			for _, conn := range mod.connections {
				var new_pulse Pulse
				if mod.state {
					new_pulse = Pulse{pulse.dest, conn, false}
					low++
				} else {
					new_pulse = Pulse{pulse.dest, conn, true}
					high++
				}
				pulses = append(pulses, new_pulse)
			}
			mod.state = !mod.state
			mods[pulse.dest] = mod

		} else if mod.kind == Conjunction {

			mod.input_memory[pulse.source] = pulse.state
			mods[pulse.dest] = mod

			every_is_high := true
			for _, state := range mod.input_memory {
				if !state {
					every_is_high = false
					break
				}
			}

			for _, conn := range mod.connections {
				var new_pulse Pulse
				if every_is_high {
					new_pulse = Pulse{pulse.dest, conn, false}
					low++
				} else {
					new_pulse = Pulse{pulse.dest, conn, true}
					high++
				}
				pulses = append(pulses, new_pulse)
			}
		} else {
			panic("bad module connection")
		}
	}

	return low, high, got_low
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

	mods := make(map[string]Mod)

	cjmods := []string{}
	for _, word := range words {
		ws := strings.Split(word, "->")
		wss := strings.Split(ws[1], ",")
		str := []string{}
		for _, mod := range wss {
			str = append(str, strings.Trim(mod, " "))
		}
		name := strings.Trim(ws[0], " ")
		mod := Mod{connections: str, state: false, kind: -1}
		if name == "broadcaster" {
			mod.kind = Broadcaster
			mods[name] = mod
		} else if name[0] == '%' {
			mod.kind = FlipFlop
			mods[name[1:]] = mod
		} else if name[0] == '&' {
			mod.kind = Conjunction
			mod.input_memory = make(map[string]bool)
			mods[name[1:]] = mod
			cjmods = append(cjmods, name[1:])
		}
	}

	// Search Conjunction modules input
	for _, cjmod := range cjmods {
		for mod_name, mod := range mods {
			if slices.Contains(mod.connections, cjmod) {
				mods[cjmod].input_memory[mod_name] = false
			}
		}
	}

	low, high := 0, 0
	for i := 0; i < 1000; i++ {
		low_, high_, _ := press_button(mods, []string{})
		low += low_
		high += high_
	}
	solution1 := low * high

	// Reset mods to initial state
	for key1, mod := range mods {
		mod.state = false
		for key2 := range mod.input_memory {
			mod.input_memory[key2] = false
		}
		mods[key1] = mod
	}


	ffs := []string{"kh", "lz", "tg", "hn"}
	cycles := [4]int{0, 0, 0, 0}

	for i := 0; i < 10000; i++ {
		_, _, got_low := press_button(mods, ffs)

		for j, ok := range got_low {
			if ok && cycles[j] == 0 {
				cycles[j] = i + 1
			}
		}
	}

	solution2 := utils.LCM(cycles[:])

	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}
