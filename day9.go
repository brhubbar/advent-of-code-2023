package main

import (
	"brh/aoc2023/internal/helpers"
	"strconv"
	"strings"
)

func Day9Part1(data string) int {
	total := 0
	var err error
	for _, line := range strings.Split(data, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		split := strings.Split(line, " ")
		history := make([]int, len(split))
		for idx, numString := range split {
			history[idx], err = strconv.Atoi(numString)
			helpers.Check(err)
		}

		total += extrapolateOnePoint(history)
	}
	return total
}

func Day9Part2(data string) int {
	total := 0
	var err error
	for _, line := range strings.Split(data, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		split := strings.Split(line, " ")
		history := make([]int, len(split))
		for idx, numString := range split {
			// Create in reverse to extrapolate the other end.
			history[len(history)-idx-1], err = strconv.Atoi(numString)
			helpers.Check(err)
		}
		total += extrapolateOnePoint(history)
	}
	return total
}

func extrapolateOnePoint(history []int) int {
	descendingSeries := [][]int{history}
outer:
	for {
		history = calcDiff(history)
		descendingSeries = append(descendingSeries, history)
		for _, val := range history {
			if val != 0 {
				// Take another derivative.
				continue outer
			}
		}
		// All of the values were 0, we're done.
		break
	}

	// I went through and optimized this code using pointers, pre-allocation, etc, but
	// found that Go was probably already doing this optimization under the hood (no
	// performance improvements).
	var diff, extrapolating *[]int
	var extrapolated int
	for i := len(descendingSeries) - 1; i > 0; i-- {
		diff = &descendingSeries[i]
		extrapolating = &descendingSeries[i-1]
		// The [int] optionally explicitly asserts that this will be an int input.
		extrapolated = getLast[int](*extrapolating) + getLast[int](*diff)
		*extrapolating = append(*extrapolating, extrapolated)
	}
	return getLast(*extrapolating)
}

func calcDiff(series []int) []int {
	diff := make([]int, len(series)-1)
	for i := 0; i < len(diff); i++ {
		diff[i] = series[i+1] - series[i]
	}
	return diff
}

// Accepts a slice of any type T, and returns the last value in the slice.
func getLast[T any](series []T) T {
	return series[len(series)-1]
}
