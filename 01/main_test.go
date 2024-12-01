package main

import (
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputFile, _ := os.ReadFile("example.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 11

	got := PartOne(input)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	inputFile, _ := os.ReadFile("example.txt")
	input := strings.Split(string(inputFile), "\n")
	want := 31

	got := PartTwo(input)

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

/*
func BenchmarkPartOne(b *testing.B) {
    inputFile, _ := os.ReadFile("input.txt")
    inputLen := len(inputFile)
    input := strings.Split(string(inputFile[:inputLen-1]), "\n")

    for i := 0; i < b.N; i++ {
        PartOne(&input)
    }
}

func BenchmarkPartTwo(b *testing.B) {
    inputFile, _ := os.ReadFile("input.txt")
    inputLen := len(inputFile)
    input := strings.Split(string(inputFile[:inputLen-1]), "\n")

    for i := 0; i < b.N; i++ {
        PartTwo(&input)
    }
}
*/
