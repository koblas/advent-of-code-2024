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
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
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
	expect := 14
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
	expect := 34
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
