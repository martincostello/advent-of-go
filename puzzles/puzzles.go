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
func (p *PuzzleData) Lines() []string {
	ll := strings.Split(string(*p), "\n")
	if len(ll) > 0 && ll[len(ll)-1] == "" {
		ll = ll[:len(ll)-1]
	}
	return ll
}

// Returns the input as a single string.
func (p *PuzzleData) String() string {
	return strings.TrimSuffix(string(*p), "\n")
}

// Represents a point in 2D space.
type Point struct {
	X int
	Y int
}

// Adds two points together and returns the result as a new point.
func Add(left, right Point) Point {
	return Point{
		X: left.X + right.X,
		Y: left.Y + right.Y,
	}
}
