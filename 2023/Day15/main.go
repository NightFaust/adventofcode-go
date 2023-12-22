package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Hash(input string) int32 {
	var current int32 = 0
	for _, c := range input {
		current = (current + c) * 17 % 256
	}

	return current
}

func Compute(input [][]byte) int32 {
	var res int32 = 0
	for _, c := range input {
		res += Hash(string(c))
	}

	return res
}

func GetLensIndex(box [][]string, lensName string) int {
	for l, lens := range box {
		if lens[0] == lensName {
			return l
		}
	}
	return -1
}

func ComputePt2(input [][]byte) int {
	res := 0
	boxes := MakeBoxes(input)
	for i, box := range boxes {
		for y, lens := range box {
			focalLength, _ := strconv.Atoi(lens[1])
			res += focalLength * (i + 1) * (y + 1)
		}
	}

	return res
}

func MakeBoxes(input [][]byte) [][][]string {
	var boxes = make([][][]string, 256)

	for _, c := range input {
		label, opAndFocal := GetData(string(c))
		boxNumber := Hash(label)
		lensIndex := GetLensIndex(boxes[boxNumber], label)

		if opAndFocal[0] == byte('=') {
			if lensIndex == -1 {
				boxes[boxNumber] = append(boxes[boxNumber], []string{label, opAndFocal[1:]})
			} else {
				boxes[boxNumber][lensIndex][1] = opAndFocal[1:]
			}
		} else {
			if lensIndex > -1 {
				boxes[boxNumber] = append(boxes[boxNumber][:lensIndex], boxes[boxNumber][lensIndex+1:]...)
			}
		}
	}

	return boxes
}

func GetData(s string) (string, string) {
	for i, r := range s {
		if r == '=' || r == '-' {
			return s[:i], s[i:]
		}
	}
	return "", ""
}

func main() {
	timeStart := time.Now()

	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sanitized := bytes.Replace(in, []byte("\n"), []byte{}, -1)
	sequence := bytes.Split(sanitized, []byte(","))
	pt1 := Compute(sequence)
	pt1Elapsed := time.Since(timeStart)

	timeStart2 := time.Now()
	pt2 := ComputePt2(sequence)

	fmt.Printf("--- Day 15: Lens Library ---\n")
	fmt.Printf("Part 1: %d | %.2fms\n", pt1, float64(pt1Elapsed.Microseconds())/1000)
	fmt.Printf("Part 2: %d | %.2fms\n", pt2, float64(time.Since(timeStart2).Microseconds())/1000)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
