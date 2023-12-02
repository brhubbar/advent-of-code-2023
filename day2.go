package main

import (
	"brh/aoc2023/internal/helpers"
	"regexp"
	"strconv"
	"strings"
)

var R = regexp.MustCompile(`(\d+) red`)
var G = regexp.MustCompile(`(\d+) green`)
var B = regexp.MustCompile(`(\d+) blue`)
var GAME_ID = regexp.MustCompile(`^Game (\d+):`)

const MAX_R = 12
const MAX_G = 13
const MAX_B = 14

func Day2Part1(data string) int {
	result := 0

	lines := strings.Split(data, "\n")

	var max_r, max_g, max_b int
	var line string
	for i := range lines {
		line = lines[i]
		if strings.TrimSpace(line) == "" {
			// Skip blank lines.
			continue
		}
		// Ugh - string conversion is verBOSE.
		game_id, err := strconv.Atoi(
			string(
				GAME_ID.FindSubmatch([]byte(line))[1],
			),
		)
		helpers.Check(err)

		// This was a nice little function to slow things down.
		max_r = get_max(R.FindAllSubmatch([]byte(line), -1))
		max_g = get_max(G.FindAllSubmatch([]byte(line), -1))
		max_b = get_max(B.FindAllSubmatch([]byte(line), -1))

		if max_r <= MAX_R &&
			max_g <= MAX_G &&
			max_b <= MAX_B {
			result += game_id
		}

	}

	return result
}

func Day2Part2(data string) int {
	result := 0

	lines := strings.Split(data, "\n")

	var max_r, max_g, max_b int
	var line string
	for i := range lines {
		line = lines[i]
		if strings.TrimSpace(line) == "" {
			// Skip blank lines.
			continue
		}
		max_r = get_max(R.FindAllSubmatch([]byte(line), -1))
		max_g = get_max(G.FindAllSubmatch([]byte(line), -1))
		max_b = get_max(B.FindAllSubmatch([]byte(line), -1))
		result += max_r * max_g * max_b
	}

	return result
}

// Expects a list of two-item lists (match, then the only capturing group) of byte
// arrays. Returns the highest value number found in each of the capturing groups, or 0
// if none captured.
func get_max(list [][][]byte) int {
	val := 0
	for i := range list {
		num, err := strconv.Atoi(string(list[i][1]))
		helpers.Check(err)
		val = max(val, num)
	}
	return val
}
