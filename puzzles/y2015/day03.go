// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015

import (
	"fmt"

	"github.com/martincostello/advent-of-go/puzzles"
)

var origin = puzzles.Point{X: 0, Y: 0}
var left = puzzles.Point{X: -1, Y: 0}
var right = puzzles.Point{X: 1, Y: 0}
var up = puzzles.Point{X: 0, Y: 1}
var down = puzzles.Point{X: 0, Y: -1}

// Day03 solves the puzzle for day 3 of Advent of Code 2015.
func Day03(input *puzzles.PuzzleData) puzzles.PuzzleSolution {
	var santa = origin
	var directions = input.AsString()
	var visited = make(map[puzzles.Point]bool)

	for _, direction := range directions {
		santa = move(santa, direction)
		visited[santa] = true
	}

	housesWithPresentsFromSanta := len(visited)

	// Reset state
	clear(visited)
	santa = origin

	var robot = origin
	var current = santa
	var previous = robot

	visited[santa] = true

	for _, direction := range directions {
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

func move(point puzzles.Point, direction rune) puzzles.Point {
	switch direction {
	case '<':
		return puzzles.Add(point, left)
	case '>':
		return puzzles.Add(point, right)
	case '^':
		return puzzles.Add(point, up)
	case 'v':
		return puzzles.Add(point, down)
	default:
		panic(fmt.Sprintf("invalid direction: %q", direction))
	}
}
