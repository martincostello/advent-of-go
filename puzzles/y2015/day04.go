// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"sync"

	"github.com/martincostello/advent-of-go/puzzles"
)

// Day04 solves the puzzle for day 4 of Advent of Code 2015.
func Day04(input string) (puzzles.PuzzleSolution, error) {
	var (
		err             error
		lowestZeroHash5 = -1
		lowestZeroHash6 = -1
		wg              sync.WaitGroup
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Go(func() {
		lowestZeroHash5, err = GetLowestPositiveNumberHash(input, 5, ctx)
		if err != nil {
			cancel()
			return
		}
	})
	wg.Go(func() {
		lowestZeroHash6, err = GetLowestPositiveNumberHash(input, 6, ctx)
		if err != nil {
			cancel()
			return
		}
	})
	wg.Wait()

	return puzzles.PuzzleSolution{
		Part1: fmt.Sprint(lowestZeroHash5),
		Part2: fmt.Sprint(lowestZeroHash6),
	}, err
}

func GetLowestPositiveNumberHash(secretKey string, zeroes int, ctx context.Context) (int, error) {
	var (
		parallelism = 20
		rangeSize   = 500
	)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	maxInt := int(^uint(0) >> 1)

	for i := 0; i < maxInt; i += parallelism * rangeSize {
		solutions := make(chan int, parallelism)

		var wg sync.WaitGroup
		for j := range parallelism {
			wg.Go(func() {
				solution := searchForSolution(secretKey, zeroes, i+(j*rangeSize), rangeSize, ctx)
				if solution != -1 {
					cancel()
				}
				solutions <- solution
			})
		}

		wg.Wait()
		close(solutions)

		best := maxInt

		for solution := range solutions {
			if solution < best && solution != -1 {
				best = solution
			}
		}

		if best != maxInt {
			return best, nil
		}
	}

	return -1, errors.New("no solution was found for the specified secret key")
}

func searchForSolution(secretKey string, zeroes int, start int, length int, ctx context.Context) int {
	limit := start + length
	for i := start; i < limit; i++ {
		select {
		case <-ctx.Done():
			return -1
		default:
			if isSolution(i, secretKey, zeroes) {
				return i
			}
		}
	}

	return -1
}

func isSolution(value int, secretKey string, zeroes int) bool {
	target := fmt.Sprintf("%s%d", secretKey, value)

	// codeql[go/weak-sensitive-data-hashing] not used for real passwords
	hash := md5.Sum([]byte(target))

	wholeBytes := zeroes / 2
	remainder := zeroes % 2
	hasHalfByte := remainder == 1

	for _, b := range hash[:wholeBytes] {
		if b != 0 {
			return false
		}
	}

	return !hasHalfByte || (hash[wholeBytes] < 0x10)
}
