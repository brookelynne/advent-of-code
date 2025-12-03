package main

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	const exampleExpected = 1227775554
	if r := Part1(getExampleFile()); r != exampleExpected {
		t.Errorf("Expected %d, got %d", exampleExpected, r)
	}
	const inputExpected = 16793817782
	if r := Part1(getInputFile()); r != inputExpected {
		t.Errorf("Expected %d, got %d", inputExpected, r)
	}
}

func Test_IsInvalidIDPart2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"1234567890", false},
		{"1122334455", false},
		{"1237894561", false},
		{"11111", true},
		{"121212121212", true},
		{"12341234", true},
		{"123123123", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsInvalidIDPart2(tt.input); got != tt.want {
				t.Errorf("IsInvalidIDPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{getExampleFile(), 4174379265},
		{"11-22", 11 + 22},
		{"95-115", 99 + 111},
		{"998-1012", 999 + 1010},
		{"1188511880-1188511890", 1188511885},
		{"222220-222224", 222222},
		{"1698522-1698528", 0},
		{"446443-446449", 446446},
		{"38593856-38593862", 38593859},
		{"565653-565659", 565656},
		{"824824821-824824827", 824824824},
		{"2121212118-2121212124", 2121212121},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Part2(tt.input); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}

}
