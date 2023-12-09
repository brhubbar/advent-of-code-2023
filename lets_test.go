package main

import (
	"brh/aoc2023/internal/helpers"
	"testing"
)

// Iterate through the days and test cases to see if the shit works.
func TestParts(t *testing.T) {
	// This is me wishing for pytest.mark.parametrize. Thanks to this gist:
	// https://gist.github.com/zaky/8304b5abbdc03600c29e
	parameters := []struct {
		file     string
		fn       func(string) int
		expected int
	}{
		// Put the most recent at the top to help speed things up.
		{"test/inputs/d9", Day9Part2, 2},
		{"test/inputs/d9", Day9Part1, 114},
		// {"test/inputs/d8_1", Day8Part1, 2},
		// {"test/inputs/d8_2", Day8Part1, 6},
		// {"test/inputs/d8_3", Day8Part2, 6},
		// {"test/inputs/d6", Day6Part2, 71503},
		// {"test/inputs/d6", Day6Part1, 288},
		// {"test/inputs/d4", Day4Part2, 30},
		// {"test/inputs/d4", Day4Part1, 13},
		// {"test/inputs/d3", Day3Part2, 467835},
		// {"test/inputs/d3", Day3Part1, 4361},
		// {"test/inputs/d2", Day2Part2, 2286},
		// {"test/inputs/d2", Day2Part1, 8},
		// {"test/inputs/d1p2", Day1Part2, 281 + 83 + 79},
		// {"test/inputs/d1p1", Day1Part1, 142},
	}

	for i := range parameters {
		param := parameters[i]
		data := helpers.Read(param.file)
		actual := param.fn(data)
		if actual != param.expected {
			t.Logf("expected %d, got %d", param.expected, actual)
			t.Fail()
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	data := helpers.Read("test/inputs/d9")
	for i := 0; i < b.N; i++ {
		Day9Part1(data)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := helpers.Read("test/inputs/d9")
	for i := 0; i < b.N; i++ {
		Day9Part2(data)
	}
}
