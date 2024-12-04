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
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
`, "\n")

var testDataB = strings.Trim(`
xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
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
	expect := 161
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}

func TestPartTwoA(t *testing.T) {
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
	expect := 48
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
