package main

import (
	"os"
	"strings"
	"testing"

	utils "github.com/gagiuntoli/adventofcode_go/utils"
)

func TestMainFunc(t *testing.T) {
	os.Args = []string{"./main", "./input.dat"}

	solutions := strings.Split(utils.CaptureStdout(main), "\n")
	utils.CheckSolution(t, solutions, "397643", "132392981697081")
}
