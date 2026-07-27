// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"testing"

	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

func TestSolve2015IsNiceV1(t *testing.T) {
	tests := []struct {
		value string
		want  bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
		{"haegwjzuvuyypabu", false},
		{"haegwjzuvuyypcdu", false},
		{"haegwjzuvuyyppqu", false},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got := y2015.IsNiceV1(tt.value)
			if got != tt.want {
				t.Errorf("IsNiceV1(%q) = %t, want %t", tt.value, got, tt.want)
			}
		})
	}
}

func TestSolve2015IsNiceV2(t *testing.T) {
	tests := []struct {
		value string
		want  bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got := y2015.IsNiceV2(tt.value)
			if got != tt.want {
				t.Errorf("IsNiceV2(%q) = %t, want %t", tt.value, got, tt.want)
			}
		})
	}
}

func TestSolve2015HasPairOfLettersWithMoreThanOneOccurrence(t *testing.T) {
	tests := []struct {
		value string
		want  bool
	}{
		{"xyxy", true},
		{"aabcdefgaa", true},
		{"abaaab", true},
		{"a", false},
		{"aa", false},
		{"aaa", false},
		{"aaaa", true},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got := y2015.HasPairOfLettersWithMoreThanOneOccurrence(tt.value)
			if got != tt.want {
				t.Errorf("HasPairOfLettersWithMoreThanOneOccurrence(%q) = %t, want %t", tt.value, got, tt.want)
			}
		})
	}
}

func TestSolve2015HasLetterThatIsTheBreadOfALetterSandwich(t *testing.T) {
	tests := []struct {
		value string
		want  bool
	}{
		{"xyx", true},
		{"abcdefeghi", true},
		{"a", false},
		{"aa", false},
		{"aaa", true},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got := y2015.HasLetterThatIsTheBreadOfALetterSandwich(tt.value)
			if got != tt.want {
				t.Errorf("HasLetterThatIsTheBreadOfALetterSandwich(%q) = %t, want %t", tt.value, got, tt.want)
			}
		})
	}
}
