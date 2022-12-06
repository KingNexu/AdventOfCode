package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	index int
	data  []string
}

type move struct {
	amount      int
	source      int
	destination int
}

func createStartStack() []stack {
	var stacks []stack

	// Create Start Stacks
	stacks = append(
		stacks,
		stack{8, []string{"H", "B", "V", "W", "N", "M", "L", "P"}},
		stack{3, []string{"M", "Q", "H"}},
		stack{8, []string{"N", "D", "B", "G", "F", "Q", "M", "L"}},
		stack{7, []string{"Z", "T", "F", "Q", "M", "W", "G"}},
		stack{4, []string{"M", "T", "H", "P"}},
		stack{8, []string{"C", "B", "M", "J", "D", "H", "G", "T"}},
		stack{6, []string{"M", "N", "B", "F", "V", "R"}},
		stack{7, []string{"P", "L", "H", "M", "R", "G", "S"}},
		stack{5, []string{"P", "D", "B", "C", "N"}},
	)
	return stacks
}

func createTestStack() []stack {
	var stacks []stack
	stacks = append(stacks,
		stack{2, []string{"Z", "N"}},
		stack{3, []string{"M", "C", "D"}},
		stack{1, []string{"P"}},
	)

	return stacks
}

func getMove(input string) move {
	// Get Values in String
	split1 := strings.Split(input, " from ")
	moveUnsplit, toUnsplit := split1[0], split1[1]

	//Amount
	moveSplit := strings.Split(moveUnsplit, " ")
	stringAmount := moveSplit[1]

	toSplit := strings.Split(toUnsplit, " to ")
	stringSource := toSplit[0]
	stringDestination := toSplit[1]

	// Convert to Int
	amount, err := strconv.ParseInt(stringAmount, 10, 16)
	source, err := strconv.ParseInt(stringSource, 10, 16)
	destination, err := strconv.ParseInt(stringDestination, 10, 16)
	if err != nil {
		log.Fatal("❌", err)
	}

	return move{int(amount), int(source) - 1, int(destination) - 1}
}

func moveCrate(craneMove move, source *stack, dest *stack) {

	var copyItems []string
	for i := 0; i < craneMove.amount; i++ {
		source.index--
		copyItems = append(copyItems, source.data[source.index])
	}

	// Reverse copy Items for Crane 9001
	for i, j := 0, len(copyItems)-1; i < j; i, j = i+1, j-1 {
		copyItems[i], copyItems[j] = copyItems[j], copyItems[i]
	}

	for _, v := range copyItems {
		if len(dest.data) <= dest.index {
			dest.data = append(dest.data, v)
			dest.index++
		} else {
			dest.data[dest.index] = v
			dest.index++
		}
	}

}

func problem() {

	// Create Stack
	portStack := createStartStack()

	// Read File
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("❌", err)
	}

	// Create Scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		craneMove := getMove(line)
		moveCrate(craneMove, &portStack[craneMove.source], &portStack[craneMove.destination])
	}

	for _, v := range portStack {
		fmt.Printf("%v", v.data[v.index-1])
	}
	fmt.Println()

}

func main() {
	problem()
}
