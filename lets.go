package main

import (
	"brh/aoc2023/internal/helpers"
	"fmt"
)

// Main is called upon `go run .`
func main() {
	data := helpers.Read("./inputs/day9")

	fmt.Printf("Part 1: %v\n", Day9Part1(data))
	fmt.Printf("Part 2: %v\n", Day9Part2(data))
}
