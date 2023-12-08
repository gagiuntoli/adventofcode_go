package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// captureStdout calls a function f and returns its stdout side-effect as string
func captureStdout(f func()) string {
	// return to original state afterwards
	// note: defer evaluates (and saves) function ARGUMENT values at definition
	// time, so the original value of os.Stdout is preserved before it is changed
	// further into this function.
	defer func(orig *os.File) {
		os.Stdout = orig
	}(os.Stdout)

	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := io.ReadAll(r)

	return string(out)
}

func check_solutions(t *testing.T, solutions []string, expected1, expected2 string) {
	sol1 := strings.Trim(strings.Split(solutions[0], ":")[1], " ")
	sol2 := strings.Trim(strings.Split(solutions[1], ":")[1], " ")
	if sol1 != expected1 {
		t.Errorf("main() = %v, want %v", sol1, expected1)
	}
	if sol2 != expected2 {
		t.Errorf("main() = %v, want %v", sol2, expected2)
	}
}

func TestMainFunc(t *testing.T) {
	os.Args = []string{"./main", "./input.dat"}

	solutions := strings.Split(captureStdout(main), "\n")
	check_solutions(t, solutions, "16409", "11795205644011")
}
