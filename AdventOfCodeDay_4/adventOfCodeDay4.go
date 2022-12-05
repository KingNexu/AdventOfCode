package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	startPos int64
	endPos   int64
}

func containsFullRange(elf1 Elf, elf2 Elf) bool {
	if elf1.startPos >= elf2.startPos && elf1.endPos <= elf2.endPos {
		return true
	} else if elf2.startPos >= elf1.startPos && elf2.endPos <= elf1.endPos {
		return true
	} else {
		return false
	}
}

func containsAnyRange(elf1 Elf, elf2 Elf) bool {
	if elf1.startPos >= elf2.startPos && elf1.startPos <= elf2.endPos {
		return true
	} else if elf2.startPos >= elf1.startPos && elf2.startPos <= elf1.endPos {
		return true
	} else {
		return false
	}
}

func problem() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("❌", err)
	}

	scanner := bufio.NewScanner(file)

	elfRangCounter := 0
	elfAnyCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		// Split Each Elfes Start Range and End Range
		// Then add it to a Slice of Elfes
		var elfSlice []Elf
		for _, v := range splitLine {

			// Split Start and End Range
			posSlice := strings.Split(v, "-")
			startPosSlice, err := strconv.ParseInt(posSlice[0], 10, 16)
			endPosSlice, err := strconv.ParseInt(posSlice[1], 10, 16)
			if err != nil {
				log.Fatal("❌: ", err)
			}

			//Add to Slice
			elfSlice = append(elfSlice, Elf{startPosSlice, endPosSlice})
		}

		// PROBLEM 1
		if containsFullRange(elfSlice[0], elfSlice[1]) == true {
			elfRangCounter++
		}

		// PROBLEM 2
		if containsAnyRange(elfSlice[0], elfSlice[1]) == true {
			elfAnyCounter++
		}
	}
	fmt.Printf("Problem 1: %v\nProblem2: %v\n", elfRangCounter, elfAnyCounter)

}

func main() {
	problem()

}
