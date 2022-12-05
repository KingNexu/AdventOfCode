package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err, "❌")
	}

	scanner := bufio.NewScanner(file)

	rockPoints := 1
	paperPoints := 2
	scissorsPoints := 3

	drawPoints := 3
	wonPoints := 6

	scoreP1 := 0
	//Problem 1
	for scanner.Scan() {
		//Splits Opponents Move from your Move
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		opponentMove, yourMove := splitLine[0], splitLine[1]

		// Problem 1
		scoreP1 := 0
		switch {
		case opponentMove == "A":
			if yourMove == "X" {
				scoreP1 = scoreP1 + rockPoints + drawPoints
			} else if yourMove == "Y" {
				scoreP1 = scoreP1 + paperPoints + wonPoints
			} else {
				scoreP1 = scoreP1 + scissorsPoints
			}
		case opponentMove == "B":
			if yourMove == "X" {
				scoreP1 = scoreP1 + rockPoints
			} else if yourMove == "Y" {
				scoreP1 = scoreP1 + paperPoints + drawPoints
			} else {
				scoreP1 = scoreP1 + scissorsPoints + wonPoints
			}
		case opponentMove == "C":
			if yourMove == "X" {
				scoreP1 = scoreP1 + rockPoints + wonPoints
			} else if yourMove == "Y" {
				scoreP1 = scoreP1 + paperPoints
			} else {
				scoreP1 = scoreP1 + scissorsPoints + drawPoints
			}
		}
		fmt.Println("🧝: ", scoreP1)
	}

	// Problem 2
	scoreP2 := 0
	for scanner.Scan() {
		//Splits Opponents Move from your Move
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		opponentMove, yourMove := splitLine[0], splitLine[1]

		switch {
		case opponentMove == "A":
			if yourMove == "Y" {
				scoreP2 = scoreP2 + rockPoints + drawPoints
			} else if yourMove == "Z" {
				scoreP2 = scoreP2 + paperPoints + wonPoints
			} else {
				scoreP2 = scoreP2 + scissorsPoints
			}
		case opponentMove == "B":
			if yourMove == "X" {
				scoreP2 = scoreP2 + rockPoints
			} else if yourMove == "Y" {
				scoreP2 = scoreP2 + paperPoints + drawPoints
			} else {
				scoreP2 = scoreP2 + scissorsPoints + wonPoints
			}
		case opponentMove == "C":
			if yourMove == "Z" {
				scoreP2 = scoreP2 + rockPoints + wonPoints
			} else if yourMove == "X" {
				scoreP2 = scoreP2 + paperPoints
			} else {
				scoreP2 = scoreP1 + scissorsPoints + drawPoints
			}
		}
		fmt.Println("🧝: ", scoreP1)
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
	}

}
