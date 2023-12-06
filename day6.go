package main

import (
	"brh/aoc2023/internal/helpers"
	"strconv"
	"strings"
)

func Day6Part1(data string) int {
	lines := strings.Split(data, "\n")

	time_strings := strings.Split(lines[0], " ")
	// Exclude the label.
	race_times := convertToSliceOfNumbers(time_strings[1:])
	distance_strings := strings.Split(lines[1], " ")
	record_distances := convertToSliceOfNumbers(distance_strings[1:])

	// 1 because we multiply
	total := 1
	var min_hold_time, max_hold_time int
	for i := 0; i < len(race_times); i++ {
		min_hold_time, max_hold_time = findWinningRange(race_times[i], record_distances[i])
		total *= (max_hold_time - min_hold_time + 1)
	}

	return total
}

func Day6Part2(data string) int {
	lines := strings.Split(data, "\n")

	time_string := strings.TrimSpace(strings.Split(lines[0], ":")[1])
	distance_string := strings.TrimSpace(strings.Split(lines[1], ":")[1])

	// Remove the label.
	var race_time, record_distance int
	var err error
	race_time, err = strconv.Atoi(strings.ReplaceAll(time_string, " ", ""))
	helpers.Check(err)
	record_distance, err = strconv.Atoi(strings.ReplaceAll(distance_string, " ", ""))
	helpers.Check(err)

	var min_hold_time, max_hold_time int

	min_hold_time, max_hold_time = findWinningRange(race_time, record_distance)

	return max_hold_time - min_hold_time + 1
}

func findWinningRange(race_time, record_distance int) (min_hold_time, max_hold_time int) {
	var hold_time int
	for hold_time = 0; hold_time < race_time; hold_time++ {
		if checkIfWins(hold_time, race_time, record_distance) {
			min_hold_time = hold_time
			break
		}
	}

	for hold_time = race_time - 1; hold_time >= 0; hold_time-- {
		if checkIfWins(hold_time, race_time, record_distance) {
			max_hold_time = hold_time
			break
		}
	}
	// Returns the declared return values.
	return min_hold_time, max_hold_time
}

func checkIfWins(hold_time, race_time, record_distance int) bool {
	// Hold for 6 seconds, travel for 1, total time is 7.
	travel_time := race_time - hold_time
	// speed_mm_ms = hold_time
	travel_distance := hold_time * travel_time
	return travel_distance > record_distance
}

func convertToSliceOfNumbers(slice []string) []int {
	new_slice := make([]int, 0, len(slice))
	var race_time int
	var err error
	for _, string_ := range slice {
		if strings.TrimSpace(string_) == "" {
			continue
		}
		race_time, err = strconv.Atoi(string_)
		helpers.Check(err)
		new_slice = append(new_slice, race_time)
	}
	return new_slice
}
