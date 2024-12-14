package main

import (
	// "fmt"

	"regexp"
	"strings"

	// "strings"
	"testing"
)

var splitter = regexp.MustCompile("\r?\n")

var testDataA = strings.Trim(`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`, "\n")

func TestPartOneA(t *testing.T) {
	lines := splitter.Split(testDataA, -1)
	input, err := ParseInput(lines)
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	value, err := PartOneSolution(input)

	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	expect := 480
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwoA(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testDataA, -1)
	input, err := ParseInput(lines)
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	value, err := PartTwoSolution(input)

	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	expect := 875318608908
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
