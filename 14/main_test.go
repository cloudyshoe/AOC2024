package main

import (
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputFile, _ := os.ReadFile("example.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 12

	got := PartOne(input, 7, 11)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

/*
func TestPartTwo(t *testing.T) {
	inputFile, _ := os.ReadFile("example.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 0

	got := PartTwo(input, 7, 11)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartOne(input, 101, 103)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartTwo(input, 101, 103)
	}
}
*/
