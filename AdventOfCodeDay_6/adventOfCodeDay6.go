package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkForDups(input string) bool {
	for i, a := range input {
		for y, b := range input {
			if i == y {
				break
			} else if string(a) == string(b) {
				return true
			}
		}
	}
	return false
}

func findMarker(line string, markerSize int) int {
	foundNoDoup := false
	startSlice := 0
	endSlice := markerSize

	for {
		checkline := line[startSlice:endSlice]
		if checkForDups(checkline) == false {
			foundNoDoup = true
		} else {
			startSlice++
			endSlice++
		}
		if foundNoDoup == true {
			break
		}
	}
	return endSlice
}

func problem() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("❌", err)
	}

	// Create Scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(findMarker(line, 4))
		fmt.Println(findMarker(line, 14))

	}

}

func main() {
	problem()
}
