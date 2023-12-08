package utils

import (
	"os"
	"testing"
	"strings"
	"io"
)

func ArrayMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	minVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < minVal {
			minVal = arr[i]
		}
	}
	return minVal
}

func ArrayMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	return maxVal
}

func ArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// captureStdout calls a function f and returns its stdout side-effect as string
func CaptureStdout(f func()) string {
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

func CheckSolution(t *testing.T, solutions []string, expected1, expected2 string) {
	sol1 := strings.Trim(strings.Split(solutions[0], ":")[1], " ")
	sol2 := strings.Trim(strings.Split(solutions[1], ":")[1], " ")
	if sol1 != expected1 {
		t.Errorf("main() = %v, want %v", sol1, expected1)
	}
	if sol2 != expected2 {
		t.Errorf("main() = %v, want %v", sol2, expected2)
	}
}
