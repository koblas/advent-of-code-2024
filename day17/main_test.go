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
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`, "\n")

var testDataB = strings.Trim(`
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
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
	expect := "4,6,3,5,6,3,5,2,1,0"
	if value != expect {
		t.Errorf("Expected %s got %s", expect, value)
	}
}

func TestPartTwoB(t *testing.T) {
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
	expect := 117440
	if value != expect {
		t.Errorf("Expected %d got %d", expect, value)
	}
}
