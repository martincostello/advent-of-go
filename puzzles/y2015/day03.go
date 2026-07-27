// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015

import (
	"fmt"

	"github.com/martincostello/advent-of-go/puzzles"
)

var (
	origin = puzzles.Point{X: 0, Y: 0}
	left   = puzzles.Point{X: -1, Y: 0}
	right  = puzzles.Point{X: 1, Y: 0}
	up     = puzzles.Point{X: 0, Y: 1}
	down   = puzzles.Point{X: 0, Y: -1}
)

// Day03 solves the puzzle for day 3 of Advent of Code 2015.
func Day03(input string) (puzzles.PuzzleSolution, error) {
	var santa = origin
	var visited = make(map[puzzles.Point]bool)

	for _, direction := range input {
		var err error
		santa, err = move(santa, direction)
		if err != nil {
			return puzzles.PuzzleSolution{}, fmt.Errorf("failed to move santa: %w", err)
		}
		visited[santa] = true
	}

	housesWithPresentsFromSanta := len(visited)

	// Reset state
	clear(visited)
	santa = origin

	var (
		robot    = origin
		current  = santa
		previous = robot
	)

	visited[santa] = true

	for _, direction := range input {
		current, _ = move(current, direction)
		visited[current] = true
		current, previous = previous, current
	}

	housesWithPresentsFromSantaAndRoboSanta := len(visited)

	return puzzles.PuzzleSolution{
		Part1: fmt.Sprint(housesWithPresentsFromSanta),
		Part2: fmt.Sprint(housesWithPresentsFromSantaAndRoboSanta),
	}, nil
}

func move(p puzzles.Point, direction rune) (puzzles.Point, error) {
	switch direction {
	case '<':
		return puzzles.Add(p, left), nil
	case '>':
		return puzzles.Add(p, right), nil
	case '^':
		return puzzles.Add(p, up), nil
	case 'v':
		return puzzles.Add(p, down), nil
	default:
		return puzzles.Point{}, fmt.Errorf("invalid direction: %q", direction)
	}
}
