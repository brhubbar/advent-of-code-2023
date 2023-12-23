package main

import (
	"fmt"
	"strings"
)

const N uint8 = 1 << 4 // 0b1000
const S uint8 = 1 << 3 // 0b0100
const E uint8 = 1 << 2 // 0b0010
const W uint8 = 1 << 1 // 0b0001

// 1111 --> nsew
var PIPE_TYPES = map[rune]uint8{
	'J': N | W, // 0b1001, wn
	'L': N | E, // 0b1010, ne

	'F': S | E, // 0b0110, es
	'7': S | W, // 0b0101, sw

	'-': E | W, // 0b0011, ew
	'|': N | S, // 0b1100, ns

	'S': N | S | E | W,
}

// Joint is valid if:
// Looking left to right, left has an east component (0b0010) and right has a west
// component (0b0001).
//
// Looking up to down, up has a south component (0b0100) and down has a north component
// (0b1000).
//
// Put more rationally, when scanning right, I need to see 0bxx1x, then 0bxxx1 for a
// valid contiguous pipe. Scanning down, I need to see 0bx1xx, then 0b1xxx.

func Day10Part1(data string) int {
	width := strings.Index(data, "\n")
	height := strings.Count(data, "\n")
	data = strings.ReplaceAll(data, "\n", "")
	mapMask := make([]uint8, width*height)

	for idx, char := range data {
		mapMask[idx] = PIPE_TYPES[char]
	}

	// Evaluate each point for valid horizontal connections.
	for idx := 0; idx < width*height; idx++ {
		if (idx+1)%width == 0 {
			// At the right edge of the map, so nothing can match. Unset the East bit.
			mapMask[idx] &= ^E
			continue
		}
		if (idx)%width == 0 {
			// At the left edge of the map, so unset the West bit.
			mapMask[idx] &= ^W
		}
		if !checkHorz(mapMask[idx], mapMask[idx+1]) {
			// Not a valid connection to the right. Unset those bits for both parties.
			mapMask[idx] &= ^E
			mapMask[idx+1] &= ^W
		}
	}

	// Evaluate each point for valid vertical connections.
	for idx := 0; idx < width*height; idx++ {
		if idx >= width*(height-1) {
			// At the bottom edge of the map, so nothing can match. Unset the South bit.
			mapMask[idx] &= ^S
			continue
		}
		if idx < width {
			// At the top edge of the map, so unset the North bit.
			mapMask[idx] &= ^N
		}
		if !checkVert(mapMask[idx], mapMask[idx+width]) {
			// Not a valid connection below. Unset those bits for both parties.
			mapMask[idx] &= ^S
			mapMask[idx+1] &= ^N
		}
	}

	printMask(data, mapMask, width)

	// Wipe anything that should be dead.
	for idx := 0; idx < width*height; idx++ {
		nBitsSet := popcount(mapMask[idx])
		fmt.Printf("%v has %v bits set.\n", mapMask[idx], nBitsSet)
		if nBitsSet == 2 || nBitsSet == 0 {
			continue
		} else if nBitsSet == 1 {
			wipe(mapMask, idx, width)
		} else {
			panic("3 or more bits set?!")
		}
	}

	printMask(data, mapMask, width)

	return 0
}

func Day10Part2(data string) int {
	return 0
}

// Recursively clear out connected pipes.
func wipe(mapMask []uint8, brokenEndIdx, width int) {
	if popcount(mapMask[brokenEndIdx]) > 1 {
		panic("Wiping something that doesn't warrant wiping.")
	}
	if (mapMask[brokenEndIdx] & N) > 0 {
		mapMask[brokenEndIdx] &= ^N
		attachedIdx := brokenEndIdx - width
		mapMask[attachedIdx] &= ^S
		wipe(mapMask, attachedIdx, width)
		return
	}
	if (mapMask[brokenEndIdx] & S) > 0 {
		mapMask[brokenEndIdx] &= ^S
		attachedIdx := brokenEndIdx + width
		mapMask[attachedIdx] &= ^N
		wipe(mapMask, attachedIdx, width)
		return
	}
	if (mapMask[brokenEndIdx] & E) > 0 {
		mapMask[brokenEndIdx] &= ^E
		attachedIdx := brokenEndIdx + 1
		mapMask[attachedIdx] &= ^W
		wipe(mapMask, attachedIdx, width)
		return
	}
	if (mapMask[brokenEndIdx] & W) > 0 {
		mapMask[brokenEndIdx] &= ^W
		attachedIdx := brokenEndIdx - 1
		mapMask[attachedIdx] &= ^E
		wipe(mapMask, attachedIdx, width)
		return
	}
}

func printMask(data string, mapMask []uint8, width int) {
	for idx, mask := range mapMask {
		if mask > 1 {
			fmt.Print(string(data[idx]))
		} else {
			fmt.Print(".")
		}
		if (idx+1)%width == 0 {
			fmt.Print("\n")
		}
	}
}

func checkHorz(left, right uint8) bool {
	// https://stackoverflow.com/a/23192263
	leftHasEast := (left & E) > 0
	rightHasWest := (right & W) > 0
	return leftHasEast && rightHasWest
}

func checkVert(above, below uint8) bool {
	// https://stackoverflow.com/a/23192263
	aboveHasSouth := (above & S) > 0
	belowHasNorth := (below & N) > 0
	return aboveHasSouth && belowHasNorth
}

// This is better when most bits in x are 0
// This algorithm works the same for all data sizes.
// This algorithm uses 3 arithmetic operations and 1 comparison/branch per "1" bit in x.
// https://en.wikipedia.org/wiki/Hamming_weight
func popcount(x uint8) uint8 {
	var count uint8
	for count = 0; x > 0; count++ {
		x &= x - 1
	}
	return count
}
