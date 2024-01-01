package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "input.txt", "")
	flag.Parse()

	timeStart := time.Now()

	input, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range strings.Split(string(input), "\n") {
		fmt.Println(c)
	}

	pt1 := 0
	pt1Elapsed := time.Since(timeStart)

	timeStart2 := time.Now()
	pt2 := 0

	fmt.Printf("--- Day 17: Clumsy Crucible ---\n")
	fmt.Printf("Part 1: %d | %.2fms\n", pt1, float64(pt1Elapsed.Microseconds())/1000)
	fmt.Printf("Part 2: %d | %.2fms\n", pt2, float64(time.Since(timeStart2).Microseconds())/1000)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
