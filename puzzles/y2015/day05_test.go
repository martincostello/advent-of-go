// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

func TestY2015IsNiceV1(t *testing.T) {
	t.Parallel()
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
		c := tt
		t.Run(c.value, func(t *testing.T) {
			t.Parallel()
			got := y2015.IsNiceV1(c.value)
			require.Equal(t, c.want, got, "IsNiceV1(%q) = %t, want %t", c.value, got, c.want)
		})
	}
}

func TestY2015IsNiceV2(t *testing.T) {
	t.Parallel()
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
		c := tt
		t.Run(c.value, func(t *testing.T) {
			t.Parallel()
			got := y2015.IsNiceV2(c.value)
			require.Equal(t, c.want, got, "IsNiceV2(%q) = %t, want %t", c.value, got, c.want)
		})
	}
}

func TestY2015HasPairOfLettersWithMoreThanOneOccurrence(t *testing.T) {
	t.Parallel()
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
		c := tt
		t.Run(c.value, func(t *testing.T) {
			t.Parallel()
			got := y2015.HasPairOfLettersWithMoreThanOneOccurrence(c.value)
			require.Equal(t, c.want, got, "HasPairOfLettersWithMoreThanOneOccurrence(%q) = %t, want %t", c.value, got, c.want)
		})
	}
}

func TestY2015HasLetterThatIsTheBreadOfALetterSandwich(t *testing.T) {
	t.Parallel()
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
		c := tt
		t.Run(c.value, func(t *testing.T) {
			t.Parallel()
			got := y2015.HasLetterThatIsTheBreadOfALetterSandwich(c.value)
			require.Equal(t, c.want, got, "HasLetterThatIsTheBreadOfALetterSandwich(%q) = %t, want %t", c.value, got, c.want)
		})
	}
}
