package main

import (
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputFile, _ := os.ReadFile("example.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 161

	got := PartOne(input)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func TestPartOneSplit(t *testing.T) {
	inputFile, _ := os.ReadFile("example.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 161

	got := PartOneSplit(input)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	inputFile, _ := os.ReadFile("example2.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 48

	got := PartTwo(input)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func TestPartTwoSplit(t *testing.T) {
	inputFile, _ := os.ReadFile("example2.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 48

	got := PartTwoSplit(input)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func BenchmarkPartOne(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartOne(input)
	}
}

func BenchmarkPartOneParser(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartOneParser(input)
	}
}

func BenchmarkPartOneSplit(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartOneSplit(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartTwo(input)
	}
}

func BenchmarkPartTwoParser(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartTwoParser(input)
	}
}

func BenchmarkPartTwoSplit(b *testing.B) {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	for i := 0; i < b.N; i++ {
		PartTwoSplit(input)
	}
}

