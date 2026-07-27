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

func IsNiceV1(value string) bool {
	if containsAny(value, &notNice) {
		return false
	}

	vowels := 0
	hasAnyConsecutiveLetters := false

	for i, ch := range value {
		if isVowel(ch) {
			vowels++
		}

		if i > 0 && !hasAnyConsecutiveLetters {
			hasAnyConsecutiveLetters = byte(ch) == value[i-1]
		}

		if hasAnyConsecutiveLetters && vowels > 2 {
			return true
		}
	}

	return false
}

func IsNiceV2(value string) bool {
	return HasPairOfLettersWithMoreThanOneOccurrence(value) && HasLetterThatIsTheBreadOfALetterSandwich(value)
}

func HasLetterThatIsTheBreadOfALetterSandwich(value string) bool {
	if len(value) < 3 {
		return false
	}

	for i := 1; i < len(value)-1; i++ {
		if value[i-1] == value[i+1] {
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

func count(values []string, predicate func(string) bool) int {
	count := 0
	for _, v := range values {
		if predicate(v) {
			count++
		}
	}
	return count
}

func containsAny(value string, values *[]string) bool {
	for _, v := range *values {
		if strings.Contains(value, v) {
			return true
		}
	}
	return false
}

func isVowel(value rune) bool {
	switch value {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
