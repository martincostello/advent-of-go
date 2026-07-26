// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

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
	lines := strings.Split(string(*input), "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

// Returns the input as a single string.
func (input *PuzzleData) AsString() string {
	return strings.TrimSuffix(string(*input), "\n")
}

// Represents a point in 2D space.
type Point struct {
	X int
	Y int
}

// Adds two points together and returns the result as a new point.
func Add(left Point, right Point) Point {
	return Point{
		X: left.X + right.X,
		Y: left.Y + right.Y,
	}
}
