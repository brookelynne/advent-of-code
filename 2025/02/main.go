package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFilename   = "input.txt"
	exampleFilename = "example.txt"
	comma           = ","
	hyphen          = "-"
)

// Day 2: Invalid IDs
func main() {
	example := getLocalFile(exampleFilename)
	fmt.Println("Part 1 (example):", Part1(example))
	fmt.Println("Part 2 (example):", Part2(example))

	input := getLocalFile(inputFilename)
	fmt.Println("Part 1 (input):", Part1(input))
	fmt.Println("Part 2 (input):", Part2(input))
}

func getLocalFile(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

// Part1 sums the IDs of invalid IDs in the provided ranges.
// An invalid ID has a number sequence repeated once, such as 2121 or 43214321.
func Part1(input string) (aggregate int) {
	// pull each range from the input
	for _, idRange := range strings.Split(input, comma) {
		low, high := getLowHigh(idRange)

		// iterate through this range
		for i := low; i <= high; i++ {
			strNum := strconv.Itoa(i)
			if len(strNum)%2 != 0 {
				continue // skip odd-length numbers
			}
			halfLength := len(strNum) / 2
			if strNum[:halfLength] == strNum[halfLength:] {
				aggregate += i
			}
		}
	}
	return aggregate
}

// Part2 sums the IDs of invalid IDs in the provided ranges.
// An invalid ID for part 2 has a number sequence repeated any number of times, such as 21212121 or 432143214321.
func Part2(input string) (aggregate int) {
	for _, idRange := range strings.Split(input, comma) {
		low, high := getLowHigh(idRange)

		// iterate through this range
		for id := low; id <= high; id++ {
			if IsInvalidIDPart2(strconv.Itoa(id)) {
				aggregate += id
			}
		}
	}
	return aggregate
}

// IsInvalidIDPart2 returns whether the provided ID is invalid according to Part 2 rules.
// An invalid ID contains repetition, e.g. 11111 or 121212 or 123123 or 123412341234.
// It checks first whether the first character repeats all the way to the end of the string.
// If not, it tries the first two characters, then three and so on until it gets to half the length of the input string.
// At that point, if the first half doesn't match the second half, it's a valid ID.
func IsInvalidIDPart2(number string) bool {
	length := len(number)
	// i is the length of the substring we're checking at the moment.
	for i := 1; i <= length/2; i++ {
		// If the length of the input string is not divisible by i, we can skip checking this substring because we
		// already know it won't repeat perfectly.
		if length%i != 0 {
			continue
		}

		// j iterates through the input number, incrementing by i, to check whether the pattern repeats perfectly throughout.
		repeatingPattern := true
		for j := 0; j+i <= length; j += i {
			if number[j:j+i] != number[:i] {
				repeatingPattern = false
				break
			}
		}
		if repeatingPattern {
			return true
		}
	}

	return false
}

// getLowHigh returns the low and high IDs in the provided idRange, which should be formatted as "low-high".
func getLowHigh(idRange string) (low int, high int) {
	lowStr, highStr, found := strings.Cut(idRange, hyphen)
	if !found {
		log.Fatal("Unable to find hyphen in ID range:", idRange)
	}

	low, err := strconv.Atoi(lowStr)
	if err != nil {
		log.Fatal("Unable to parse low ID:", err)
	}
	high, err = strconv.Atoi(highStr)
	if err != nil {
		log.Fatal("Unable to parse high ID:", err)
	}
	return low, high
}
