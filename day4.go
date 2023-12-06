package main

import (
	"math"
	"strings"
)

func Day4Part1(data string) int {
	total := 0
	var you_win int
	for _, card := range strings.Split(data, "\n") {
		if strings.TrimSpace(card) == "" {
			continue
		}
		you_win = countWinningNumbers(card)
		if you_win > 0 {
			// fmt.Printf("%v won %v times\n", card, you_win)
			total += int(math.Pow(2, float64(you_win-1)))
		}
	}
	return total
}

func Day4Part2(data string) int {
	total := 0
	cards := strings.Split(data, "\n")
	var you_win int
	// -1 to account for the "" line at the end.
	n_wins_per_card := make([]int, len(cards)-1)
	for idx, card := range cards {
		if strings.TrimSpace(card) == "" {
			continue
		}
		you_win = countWinningNumbers(card)
		n_wins_per_card[idx] = you_win
	}

	n_cards := make([]int, len(cards)-1)
	for idx, n_wins := range n_wins_per_card {
		// Count the original card!
		n_cards[idx] += 1
		for i := 0; i < n_cards[idx]; i++ {
			for j := idx + 1; j <= idx+n_wins; j++ {
				n_cards[j] += 1
			}
		}
		total += n_cards[idx]
	}
	return total
}

func countWinningNumbers(card string) int {
	_, numbers, isFound := strings.Cut(card, ":")
	checkFound(isFound)
	winning, yours, isFound := strings.Cut(numbers, "|")
	checkFound(isFound)
	s_winning := SetFromList(strings.Split(winning, " "))
	s_yours := SetFromList(strings.Split(yours, " "))
	s_won := s_winning.Intersection(s_yours)
	// fmt.Printf("Winners: %v\n", s_won)
	// "": true seems to always be present, so just exclude it from the count.
	return len(s_won) - 1
}

func checkFound(isFound bool) {
	if !isFound {
		panic("Failed to split.")
	}
}

type Set map[string]bool

func SetFromList(list_ []string) Set {
	s := map[string]bool{}
	for _, str := range list_ {
		s[str] = true
	}
	return s
}

func (s1 Set) Intersection(s2 Set) Set {
	// https://stackoverflow.com/a/34020023
	s_intersection := map[string]bool{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k, _ := range s1 {
		if s2[k] {
			s_intersection[k] = true
		}
	}
	return s_intersection
}
