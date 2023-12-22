package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	timeStart := time.Now()
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	pt1 := solvePartOne(lines)
	timeElapsed := time.Since(timeStart)

	timeStart2 := time.Now()
	pt2 := solvePartTwo(lines)

	fmt.Printf("--- Day 01: Trebuchet?! ---\n")
	fmt.Printf("Part 1: %d | %.2fms\n", pt1, float64(timeElapsed.Microseconds())/1000)
	fmt.Printf("Part 2: %d | %.2fms\n", pt2, float64(time.Since(timeStart2).Microseconds())/1000)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

// 56108
func solvePartOne(lines []string) int {
	result := 0
	for _, line := range lines {
		var numbers []int
		for _, char := range line {
			if n := int(char - '0'); 0 <= n && n <= 9 {
				numbers = append(numbers, n)
			}
		}
		result += numbers[0]*10 + numbers[len(numbers)-1]
	}
	return result
}

// 55652
func solvePartTwo(lines []string) int {
	result := 0
	textNumber := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range lines {
		var numbers []int
		for i, char := range line {
			if n := int(char - '0'); 0 <= n && n <= 9 {
				numbers = append(numbers, n)
				continue
			}

			for text, number := range textNumber {
				if i+len(text)-1 < len(line) && line[i:i+len(text)] == text {
					numbers = append(numbers, number)
					break
				}
			}
		}
		result += numbers[0]*10 + numbers[len(numbers)-1]
	}

	return result
}
