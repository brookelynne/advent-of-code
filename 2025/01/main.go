package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	dialSize = 100
	newline  = "\n"
)

func main() {
	part1()
	part2Division()
	part2Brute()
}

// Count the number of times the dial lands on zero.
func part1() {
	var zeroCount uint32
	currentPos := 50 // Given in instructions

	for _, twist := range strings.Split(input, newline) {
		if len(twist) == 0 {
			continue
		}

		magnitude, err := strconv.Atoi(twist[1:])
		if err != nil {
			log.Fatal("Unable to parse magnitude from twist instruction:", err)
		}

		switch twist[0] {
		case 'R':
			currentPos += magnitude
		case 'L':
			currentPos -= magnitude
		default:
			log.Fatal("Invalid twist direction")
		}

		if currentPos%dialSize == 0 {
			zeroCount++
		}
	}

	fmt.Println("Number of Zeroes:", zeroCount)
}

// Count the number of times the dial touches zero, whether it crosses it or lands on it.
// Solution achieved with the help of GitHub Copilot (Gemini 2.5 Pro).
func part2Division() {
	var zeroCount int
	currentPos := 50 // Starting position

	for _, twist := range strings.Split(input, newline) {
		if len(twist) == 0 {
			continue
		}

		magnitude, err := strconv.Atoi(twist[1:])
		if err != nil {
			log.Fatal("Unable to parse magnitude from twist instruction:", err)
		}

		var distanceToZero int
		switch twist[0] {
		case 'R':
			distanceToZero = dialSize - currentPos
			currentPos += magnitude
		case 'L':
			distanceToZero = currentPos
			currentPos -= magnitude
		default:
			log.Fatal("Invalid twist direction")
		}

		if distanceToZero == 0 { // If we start at 0, the first cross is a full rotation away
			distanceToZero = dialSize
		}
		if magnitude >= distanceToZero {
			zeroCount++
			zeroCount += (magnitude - distanceToZero) / dialSize
		}

		// Keep the current position within the dial size.
		currentPos %= dialSize
		if currentPos < 0 {
			currentPos += dialSize
		}
	}

	fmt.Println("Number of Zeroes:", zeroCount)
}

// Count the number of times the dial touches zero, whether it crosses it or lands on it.
func part2Brute() {
	var zeroCount uint32
	currentPos := 50 // Given in instructions

	for _, twist := range strings.Split(input, newline) {
		if len(twist) == 0 {
			continue
		}

		magnitude, err := strconv.Atoi(twist[1:])
		if err != nil {
			log.Fatal("Unable to parse magnitude from twist instruction:", err)
		}

		switch twist[0] {
		case 'R':
			for range magnitude {
				currentPos++
				if currentPos >= dialSize {
					currentPos -= dialSize
				}
				if currentPos == 0 {
					zeroCount++
				}
			}
		case 'L':
			for range magnitude {
				currentPos--
				if currentPos < 0 {
					currentPos += dialSize
				}
				if currentPos == 0 {
					zeroCount++
				}
			}
		default:
			log.Fatal("Invalid twist direction")
		}
	}

	fmt.Println("Number of Zeroes:", zeroCount)

}
