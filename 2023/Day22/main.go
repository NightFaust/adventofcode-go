package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

// Thx mnml (https://www.reddit.com/user/mnml/)
// => https://github.com/mnml/aoc/blob/main/2023/22/1.go
type Brick struct {
	Start Cord
	End   Cord
}

type Cord struct {
	X, Y, Z int
}

func Parse(lines [][]byte) []Brick {
	bricks := make([]Brick, 0, 7)

	for _, line := range lines {
		var brick Brick
		_, err := fmt.Sscanf(
			string(line),
			"%d,%d,%d~%d,%d,%d",
			&brick.Start.X,
			&brick.Start.Y,
			&brick.Start.Z,
			&brick.End.X,
			&brick.End.Y,
			&brick.End.Z,
		)
		if err != nil {
			return nil
		}
		bricks = append(bricks, brick)
	}

	slices.SortFunc(bricks, func(i, j Brick) int { return cmp.Compare(i.Start.Z, j.Start.Z) })

	return bricks
}

// Decrease Z by 1 if possible, and return number of bricks that were modified
func drop(bricks []Brick) ([]Brick, int) {
	bs := slices.Clone(bricks)
	n := 0

	for i, a := range bs {
		if a.Start.Z > 1 && a.End.Z > 1 {
			a.Start.Z, a.End.Z = a.Start.Z-1, a.End.Z-1

			if !isOverlapping(bs[:i], a) {
				bs[i] = a
				n++
			}
		}
	}
	return bs, n
}

func isOverlapping(bricks []Brick, a Brick) bool {
	for _, b := range bricks {
		if a.Start.X <= b.End.X && a.End.X >= b.Start.X &&
			a.Start.Y <= b.End.Y && a.End.Y >= b.Start.Y &&
			a.Start.Z <= b.End.Z && a.End.Z >= b.End.Z {
			return true
		}
	}
	return false
}

func main() {
	timeStart := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(input, []byte("\n"))

	bricks := Parse(lines)

	var n int
	for {
		bricks, n = drop(bricks)
		if n == 0 {
			break
		}
	}

	pt1, pt2 := 0, 0
	for i := range bricks {
		_, n := drop(slices.Delete(slices.Clone(bricks), i, i+1))
		if n == 0 {
			pt1++
		}
		pt2 += n
	}

	fmt.Printf("--- Day 22: Sand Slabs ---\n")
	fmt.Printf("Part 1: %d | %.2fms\n", pt1, float64(time.Since(timeStart).Microseconds())/1000)
	fmt.Printf("Part 2: %d | %.2fms\n", pt2, float64(time.Since(timeStart).Microseconds())/1000)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
