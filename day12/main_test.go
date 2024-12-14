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
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE
`, "\n")

var testDataB = strings.Trim(`
OOO
OXO
AOO
`, "\n")

var testDataC = strings.Trim(`
VVVVVCRRCCCCCCCYYC
CVCCVCCCCCCCCCCYCC
CCCCCCCCCCCCCCCCCC
CCCQQCCCCCCCCCCCCC
QQQQCCCCCCCCCCCCCC
QQQQQQCCCCCCCCCCCC
QQQQQQQCCKCKKCCCYY
QQQQQQQQKKKKKKKCYY
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
	expect := 1930
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
	expect := 1206
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartB(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testDataB, -1)
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
	expect := 78
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartC(t *testing.T) {
	var lines = regexp.MustCompile("\r?\n").Split(testDataC, -1)
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
	expect := 4614
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
