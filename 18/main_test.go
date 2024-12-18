package main

import (
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputFile, _ := os.ReadFile("example.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 22

	got := PartOne(input, 7, 7, 12)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	inputFile, _ := os.ReadFile("example.txt")
	input := strings.Split(string(inputFile), "\n")
	want := "6,1"

	got := PartTwo(input, 7, 7)

	if got != want {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartOne(input, 71, 71, 1024)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartTwo(input, 71, 71)
	}
}
