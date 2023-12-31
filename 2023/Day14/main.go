package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// Thx dvk0 (https://www.reddit.com/user/dvk0/)
// => https://github.com/dannyvankooten/advent-of-code/blob/main/2023/14-parabolic-reflector-dish/main.go
const ObjRock = byte('O')
const ObjEmpty = byte('.')

var directions = [][]int{
	{0, -1}, // N
	{-1, 0}, // W
	{0, 1},  // S
	{1, 0},  // E
}

func tilt(grid [][]byte, dir []int) {
	width := len(grid[0])
	height := len(grid)

	startCol := 0
	colStep := 1
	if dir[0] == 1 {
		startCol = width - 1
		colStep = -1
	}
	startRow := 0
	rowStep := 1
	if dir[1] == 1 {
		startRow = height - 1
		rowStep = -1
	}

	for row := startRow; row < height && row >= 0; row += rowStep {
		for col := startCol; col < width && col >= 0; col += colStep {
			if grid[row][col] != ObjRock {
				continue
			}

			// walk into direction until obstacle
			c := col
			r := row
			for c+dir[0] >= 0 && c+dir[0] < width && r+dir[1] >= 0 && r+dir[1] < height && grid[r+dir[1]][c+dir[0]] == ObjEmpty {
				c += dir[0]
				r += dir[1]
			}

			// place rock at new position
			if r != row || c != col {
				grid[r][c] = ObjRock
				grid[row][col] = ObjEmpty
			}
		}
	}
}

func weight(grid [][]byte) int {
	width := len(grid[0])
	height := len(grid)
	weight := 0

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if grid[row][col] != ObjRock {
				continue
			}

			weight += height - row
		}
	}

	return weight
}

func isEqual(a [][]byte, b [][]byte) bool {
	for row := 0; row < len(a); row++ {
		if !bytes.Equal(a[row], b[row]) {
			return false
		}
	}

	return true
}

func copyGrid(a [][]byte) [][]byte {
	rows := make([][]byte, len(a))

	for i := range a {
		rows[i] = make([]byte, len(a[i]))
		copy(rows[i], a[i])
	}
	return rows
}

// Floyd's Tortoise and Hare algorithm
// https://en.wikipedia.org/wiki/Cycle_detection
func floyd(f func([][]byte), x0 [][]byte) (int, [][]byte) {
	hare := copyGrid(x0)
	tortoise := copyGrid(x0)

	f(tortoise)
	f(hare)
	f(hare)

	for !isEqual(tortoise, hare) {
		f(tortoise)
		f(hare)
		f(hare)
	}

	// move hare forward until start of next cycle
	// tortoise remains in place
	cycleLength := 1
	f(hare)
	for !isEqual(tortoise, hare) {
		f(hare)
		cycleLength += 1
	}

	// return length of cycle and current state of grid
	return cycleLength, hare
}

func main() {
	var file string
	flag.StringVar(&file, "f", "input.txt", "")
	flag.Parse()

	timeStart := time.Now()

	input, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	grid := bytes.Split(bytes.TrimSpace(input), []byte("\n"))

	// pt1
	grid1 := copyGrid(grid)
	tilt(grid1, directions[0])
	pt1 := weight(grid1)
	timeElapsed := time.Since(timeStart)
	timeStart2 := time.Now()
	// part 2
	cycle := func(g [][]byte) {
		tilt(g, directions[0])
		tilt(g, directions[1])
		tilt(g, directions[2])
		tilt(g, directions[3])
	}
	cycleLength, state := floyd(cycle, grid)
	cyclesRemaining := 1e9 % cycleLength
	for i := 0; i < cyclesRemaining; i++ {
		cycle(state)
	}

	pt2 := weight(state)
	fmt.Printf("--- Day 14: Parabolic Reflector Dish ---\n")
	fmt.Printf("Part 1: %d | %.2fms\n", pt1, float64(timeElapsed.Microseconds())/1000)
	fmt.Printf("Part 2: %d | %.2fms\n", pt2, float64(time.Since(timeStart2).Microseconds())/1000)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
