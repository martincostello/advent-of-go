package main

import (
	"fmt"

	"github.com/martincostello/advent-of-go/cmd"
	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/solver"
)

func main() {
	year, day, data := cmd.Parse()

	fmt.Printf("Solving Advent of Code for day %02d of %04d\n", day, year)
	fmt.Println()

	var input = &puzzles.PuzzleInput{
		Year:  year,
		Day:   day,
		Input: data,
	}

	solution := solver.Solve(input)

	fmt.Printf("Part 1: %s\n", solution.Part1)
	fmt.Printf("Part 2: %s\n", solution.Part2)

	if len(solution.Visualization) > 0 {
		fmt.Printf("Visualization:\n%s\n", solution.Visualization)
	}
}
