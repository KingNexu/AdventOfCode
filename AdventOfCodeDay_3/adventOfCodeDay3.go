package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type backpack struct {
	index        int
	compartment1 string
	compartment2 string
}

type elfGroup struct {
	back1 string
	back2 string
	back3 string
}

var POINTS = make(map[string]int)
var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func fillPoint() {
	for n, l := range chars {
		l_str := string(l)
		POINTS[l_str] = n + 1
	}
}

func intValueOf(letter string) int {
	var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alphabet, letter) + 1
}

func main() {
	//fillpoints
	fillPoint()

	//Open File
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("❌", err)
	}

	//Create Scanner
	// scanner := bufio.NewScanner(file)
	scanner := bufio.NewScanner(file)

	var tempGroupSlice []string
	totalItemValueGroups := 0

	// PROBLEM 1

	//Compare Compartments
	lineNumber := 0
	score := 0
	for scanner.Scan() {
		line := scanner.Text()

		elfpack := new(backpack)

		elfpack.index = lineNumber
		elfpack.compartment1 = line[:(len(line) / 2)]
		elfpack.compartment2 = line[(len(line) / 2):]

		addedScore := 0

		for i := 0; i < len(elfpack.compartment1); i++ {
			for j := 0; j < len(elfpack.compartment2); j++ {
				if string(elfpack.compartment1[i]) == string(elfpack.compartment2[j]) {
					if addedScore == 0 {
						score = score + POINTS[string(elfpack.compartment2[j])]
						addedScore++
						// fmt.Printf("Line: %v // PointsA: %v SCORE: %v Chars: %v %v Positions// A: %v B: %v\n", lineNumber, POINTS[string(elfpack.compartment2[j])], score, string(elfpack.compartment1[i]), string(elfpack.compartment2[j]), i, j)
					}
				}
			}
		}

		lineNumber++

		//PROBLEM 2

		tempGroupSlice = append(tempGroupSlice, line)

		// once our tempGroupSlice slice reaches the length of 3, iterate over the chars of the first string and check if any
		// of the characters are present in both the second and the third string
		// if it does, add it's int value to the total group value and nil the slice
		if len(tempGroupSlice) == 3 {
			for _, char := range tempGroupSlice[0] {
				if strings.Contains(tempGroupSlice[1], string(char)) && strings.Contains(tempGroupSlice[2], string(char)) {
					totalItemValueGroups = totalItemValueGroups + intValueOf(string(char))
					tempGroupSlice = nil
					break
				}
			}
		}

	}
	fmt.Println("Score:", score)
	fmt.Println("Group Score", totalItemValueGroups)

}
