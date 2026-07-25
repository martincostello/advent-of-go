package puzzles

import "strings"

// Represents the input data for a single Advent of Code puzzle.
type PuzzleData []byte

// PuzzleInput represents the input data for a single Advent of Code puzzle.
type PuzzleInput struct {
	Year  int
	Day   int
	Input *PuzzleData
}

// PuzzleSolution represents the solution to a single Advent of Code puzzle.
type PuzzleSolution struct {
	Part1         string
	Part2         string
	Visualization string
}

// Returns the input as a slice of lines.
func (input *PuzzleData) AsLines() []string {
	return strings.Split(string(*input), "\n")
}

// Returns the input as a single string.
func (input *PuzzleData) AsString() string {
	return string(*input)
}
