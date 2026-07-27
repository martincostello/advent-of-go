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
func Day03(input string) puzzles.PuzzleSolution {
	var santa = origin
	var visited = make(map[puzzles.Point]bool)

	for _, direction := range input {
		santa = move(santa, direction)
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
		current = move(current, direction)
		visited[current] = true
		current, previous = previous, current
	}

	housesWithPresentsFromSantaAndRoboSanta := len(visited)

	return puzzles.PuzzleSolution{
		Part1: fmt.Sprint(housesWithPresentsFromSanta),
		Part2: fmt.Sprint(housesWithPresentsFromSantaAndRoboSanta),
	}
}

func move(p puzzles.Point, direction rune) puzzles.Point {
	switch direction {
	case '<':
		return puzzles.Add(p, left)
	case '>':
		return puzzles.Add(p, right)
	case '^':
		return puzzles.Add(p, up)
	case 'v':
		return puzzles.Add(p, down)
	default:
		panic(fmt.Sprintf("invalid direction: %q", direction))
	}
}
