package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Ball struct {
	Count int
	Color string
}

func resolve(lines []string) (int, int) {
	pt1 := 0
	pt2 := 0
	for i, line := range lines {
		line = strings.Split(line, ":")[1]
		isOk, res := compute(line)

		if isOk {
			pt1 += i + 1
		}
		pt2 += res
	}
	return pt1, pt2
}

func compute(line string) (bool, int) {
	isOk := true
	r, g, b := 1, 1, 1
	parts := strings.Split(line, ";")
	for _, part := range parts {
		for _, ball := range strings.Split(part, ",") {
			var ba Ball
			_, err := fmt.Sscanf(ball, "%d %s", &ba.Count, &ba.Color)
			if err != nil {
				log.Fatal(err)
			}

			if (ba.Color == "red" && ba.Count > 12) || (ba.Color == "green" && ba.Count > 13) || (ba.Color == "blue" && ba.Count > 14) {
				isOk = false
			}
			if ba.Color == "red" && ba.Count > r {
				r = ba.Count
			}
			if ba.Color == "green" && ba.Count > g {
				g = ba.Count
			}
			if ba.Color == "blue" && ba.Count > b {
				b = ba.Count
			}
		}
	}
	return isOk, r * g * b
}

func main() {
	timeStart := time.Now()
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")

	pt1, pt2 := resolve(lines)
	timeElapsed := time.Since(timeStart)

	fmt.Printf("--- Day 02: Cube Conundrum ---\n")
	fmt.Printf("Part 1: %d | %.2fms\n", pt1, float64(timeElapsed.Microseconds())/1000)
	fmt.Printf("Part 2: %d | %.2fms\n", pt2, float64(timeElapsed.Microseconds())/1000)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
