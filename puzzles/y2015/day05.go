// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015

import (
	"fmt"
	"strings"

	"github.com/martincostello/advent-of-go/puzzles"
)

var notNice = []string{"ab", "cd", "pq", "xy"}

// Day05 solves the puzzle for day 5 of Advent of Code 2015.
func Day05(input []string) puzzles.PuzzleSolution {
	niceStringCountV1 := count(input, IsNiceV1)
	niceStringCountV2 := count(input, IsNiceV2)

	return puzzles.PuzzleSolution{
		Part1: fmt.Sprint(niceStringCountV1),
		Part2: fmt.Sprint(niceStringCountV2),
	}
}

func IsNiceV1(s string) bool {
	if containsAny(s, notNice) {
		return false
	}

	vowels := 0
	hasAnyConsecutiveLetters := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		if isVowel(rune(c)) {
			vowels++
		}

		if i > 0 && !hasAnyConsecutiveLetters {
			hasAnyConsecutiveLetters = c == s[i-1]
		}

		if hasAnyConsecutiveLetters && vowels > 2 {
			return true
		}
	}

	return false
}

func IsNiceV2(s string) bool {
	return HasPairOfLettersWithMoreThanOneOccurrence(s) && HasLetterThatIsTheBreadOfALetterSandwich(s)
}

func HasLetterThatIsTheBreadOfALetterSandwich(s string) bool {
	if len(s) < 3 {
		return false
	}

	for i := 1; i < len(s)-1; i++ {
		if s[i-1] == s[i+1] {
			return true
		}
	}

	return false
}

func HasPairOfLettersWithMoreThanOneOccurrence(s string) bool {
	if len(s) < 4 {
		return false
	}

	pairs := make(map[string]int, len(s)-1)

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if j, ok := pairs[pair]; ok {
			if i-j >= 2 {
				return true
			}
			continue
		}

		pairs[pair] = i
	}

	return false
}

func count(ss []string, predicate func(string) bool) int {
	c := 0
	for _, s := range ss {
		if predicate(s) {
			c++
		}
	}
	return c
}

func containsAny(s string, ss []string) bool {
	for _, v := range ss {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
