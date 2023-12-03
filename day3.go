package main

import (
	"brh/aoc2023/internal/helpers"
	"regexp"
	"strconv"
	"strings"
)

var PART_SYMBOL *regexp.Regexp = regexp.MustCompile(`[^\w\d\s.\n]`)
var PART_NUMBER *regexp.Regexp = regexp.MustCompile(`\d+`)

func Day3Part1(data string) int {
	COLUMN_WIDTH := strings.Index(data, "\n") + 1
	TOTAL_SIZE := len(data)
	total := 0

	// Locate the span of all part numbers.
	part_numbers := PART_NUMBER.FindAllStringIndex(data, -1)
	// Part_id, start, end, start, end, start, end]. Cannot use a map because of
	// repeats!
	var part_number_spans [][7]int

	// Essentially a for each, or a for .. in enumerate() from python.
	for _, span := range part_numbers {
		part_id, err := strconv.Atoi(data[span[0]:span[1]])
		helpers.Check(err)

		part_number_spans = append(part_number_spans, [7]int{
			part_id,
			span[0] - 1 - COLUMN_WIDTH, span[1] + 1 - COLUMN_WIDTH,
			// Clip so we don't fall off the page. The others are check for falling off
			// the stage at search time.
			max(0, span[0]-1), min(TOTAL_SIZE, span[1]+1),
			span[0] - 1 + COLUMN_WIDTH, span[1] + 1 + COLUMN_WIDTH,
		})
	}

	symbols := PART_SYMBOL.FindAllStringIndex(data, -1)

	var part_id, sym_loc int
	for _, sym_span := range symbols {
		sym_loc = sym_span[0]
		// symbol := data[sym_loc]
		// fmt.Printf("\nsymbol location: %v\n", sym_loc)
		for _, num_spans := range part_number_spans {
			part_id = num_spans[0]
			// fmt.Printf("%v search locations: %v\n", part_id, num_spans)
			if sym_loc >= num_spans[1] && sym_loc < num_spans[2] ||
				sym_loc >= num_spans[3] && sym_loc < num_spans[4] ||
				sym_loc >= num_spans[5] && sym_loc < num_spans[6] {

				// fmt.Printf("Found symbol %c (%v) in %v range (%v)\n", symbol, sym_loc, part_id, num_spans)
				total += part_id
			}
		}
	}
	return total
}

func Day3Part2(data string) int {
	COLUMN_WIDTH := strings.Index(data, "\n") + 1
	TOTAL_SIZE := len(data)
	total := 0

	// Locate the span of all part numbers.
	part_numbers := PART_NUMBER.FindAllStringIndex(data, -1)
	// Part_id, start, end, start, end, start, end]. Cannot use a map because of
	// repeats!
	var part_number_spans [][7]int

	// Essentially a for each, or a for .. in enumerate() from python.
	for _, span := range part_numbers {
		part_id, err := strconv.Atoi(data[span[0]:span[1]])
		helpers.Check(err)

		part_number_spans = append(part_number_spans, [7]int{
			part_id,
			span[0] - 1 - COLUMN_WIDTH, span[1] + 1 - COLUMN_WIDTH,
			// Clip so we don't fall off the page. The others are check for falling off
			// the stage at search time.
			max(0, span[0]-1), min(TOTAL_SIZE, span[1]+1),
			span[0] - 1 + COLUMN_WIDTH, span[1] + 1 + COLUMN_WIDTH,
		})
	}

	symbols := PART_SYMBOL.FindAllStringIndex(data, -1)

	var part_id, sym_loc, n_touching, gear_ratio int
	for _, sym_span := range symbols {
		sym_loc = sym_span[0]
		// Turns out only * gears ever touch more than one part number, so no need to
		// filter by it (except for efficiency's sake).
		// symbol := data[sym_loc]
		gear_ratio = 1
		n_touching = 0

		// fmt.Printf("\nsymbol location: %v\n", sym_loc)
		for _, num_spans := range part_number_spans {
			part_id = num_spans[0]
			// fmt.Printf("%v search locations: %v\n", part_id, num_spans)
			if sym_loc >= num_spans[1] && sym_loc < num_spans[2] ||
				sym_loc >= num_spans[3] && sym_loc < num_spans[4] ||
				sym_loc >= num_spans[5] && sym_loc < num_spans[6] {

				// fmt.Printf("Found symbol %c (%v) in %v range (%v)\n", symbol, sym_loc, part_id, num_spans)
				gear_ratio *= part_id
				n_touching += 1
			}
		}
		if n_touching > 1 {
			total += gear_ratio
		}
	}
	return total
}
