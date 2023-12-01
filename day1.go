// Sub-packages go into sub-direcotories. Putting everything into the main package just
// makes the code available across the files, which is interesting and not super
// intuitive in my opinion. Going to need to learn more about the idiomatic way to
// structure code. A bunch of 1-file directories doesn't make a ton of sense either.
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Build two-digit numbers using the first and last number found in each line.
func Day1Part1(data string) int {
	// Walrus := is shorthand for var result string = ...
	var lines []string = strings.Split(data, "\n")

	find_digits := regexp.MustCompile(`(\d)`)

	var total int

	for i := range lines {
		line := lines[i]
		if line == "" {
			continue
		}
		result := find_digits.FindAll([]byte(line), -1)
		if result == nil {
			panic(fmt.Sprintf("Empty result on %q", line))
		}

		var digits string = string(result[0]) + string(result[len(result)-1])
		number, err := strconv.Atoi(digits)
		if err != nil {
			panic(err)
		}
		total += number
	}

	return total
}

// Build two-digit numbers using the first and last number found in each line. Support
// spelled-out digits like one, two... This includes eighthree --> 8 and 3
func Day1Part2(data string) int {
	// Walrus := is shorthand for var result string = ...
	var lines []string = strings.Split(data, "\n")

	find_digits := regexp.MustCompile(`(\d)`)

	var total int

	for i := range lines {
		line := lines[i]
		if line == "" {
			continue
		}
		line = replaceWordsWithNumbers(line)
		result := find_digits.FindAll([]byte(line), -1)
		if result == nil {
			panic(fmt.Sprintf("Empty result on %q", line))
		}

		var digits string = string(result[0]) + string(result[len(result)-1])
		number, err := strconv.Atoi(digits)
		if err != nil {
			panic(err)
		}
		total += number
	}

	return total
}

// Replace spelled out numbers with their numerical counterpart. Do this in such a way
// that two words mushed together results in both numbers being present (eighthree --> 8
// and 3).
func replaceWordsWithNumbers(line string) string {
	// Leave the end letters for overlapping words. Thanks, u/EyedMoon.
	// https://www.reddit.com/r/adventofcode/comments/1884fpl/comment/kbiywz6/
	conversions := []struct{ word, digit string }{
		{"one", "o1e"},
		{"two", "t2o"},
		{"three", "t3e"},
		{"four", "f4r"},
		{"five", "f5e"},
		{"six", "s6x"},
		{"seven", "s7n"},
		{"eight", "e8t"},
		{"nine", "n9e"},
	}

	for i := range conversions {
		conv := conversions[i]
		line = strings.ReplaceAll(line, conv.word, conv.digit)
	}

	return line
}
