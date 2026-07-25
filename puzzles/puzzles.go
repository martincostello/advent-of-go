package puzzles

// PuzzleInput represents the input data for a single Advent of Code puzzle.
type PuzzleInput struct {
	Year  int
	Day   int
	Input []byte
}

// PuzzleSolution represents the solution to a single Advent of Code puzzle.
type PuzzleSolution struct {
	Part1         string
	Part2         string
	Visualization string
}
