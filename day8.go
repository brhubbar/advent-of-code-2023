package main

import (
	"brh/aoc2023/internal/helpers"
	"fmt"
	"regexp"
	"strings"
)

func Day8Part1(data string) int {
	directions, network := convertInputToDirectionsAndNetwork(data)

	start := "AAA"
	finish := "ZZZ"

	return countStepsFromStartToFinish(start, finish, directions, network)
}

func Day8Part2(data string) int {
	directions, network := convertInputToDirectionsAndNetwork(data)

	// Be careful not to create a pre-filled sucker.
	steps_to_finish := make([]int, 0, 10)
	finish := "Z"

	for start, _ := range network {
		if !strings.HasSuffix(start, "A") {
			continue
		}
		steps_to_finish = append(
			steps_to_finish,
			countStepsFromStartToFinish(start, finish, directions, network),
		)
	}
	// Boldly assume periodicity of each parallel path. This means that they all arrive
	// at the least common multiple of the distance to each finishing point.
	return helpers.LCM(steps_to_finish...)
}

// Returns the number of steps taken when the current node ends with `finish`.
func countStepsFromStartToFinish(
	start,
	finish string,
	directions []int,
	network map[string][2]string,
) (n_steps int) {
	here := start
	// This is a while True.
	for {
		for _, step := range directions {
			here = network[here][step]
			n_steps += 1
			if strings.HasSuffix(here, finish) {
				return n_steps
			}
		}
	}
}

// Parse the input into a slice of integers representing left/right steps (used to index
// into the mapping), and a mapping pointing each node to a 2-sized array of nodes to
// arrive at, depending on if left or right step is taken.
func convertInputToDirectionsAndNetwork(data string) ([]int, map[string][2]string) {
	split := strings.Split(data, "\n")
	directions := make([]int, 0, len(split[0]))

	for _, step := range split[0] {
		// L = 0, R = 1.
		if step == 'L' {
			directions = append(directions, 0)
		} else if step == 'R' {
			directions = append(directions, 1)
		} else {
			fmt.Printf("[WARNING] Dropping a %q from directions.\n", step)
		}
	}

	mapping := make(map[string][2]string, len(directions))

	var nodes []string
	NODE_ID := regexp.MustCompile(`\w\w\w`)
	for _, node := range split[1:] {
		if strings.TrimSpace(node) == "" {
			continue
		}

		nodes = NODE_ID.FindAllString(node, 3)
		// here = [left, right]
		mapping[nodes[0]] = [2]string{nodes[1], nodes[2]}
	}
	return directions, mapping
}
