package aoc2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFilePath = "input/day7.txt"
	sep           = " "
	colon         = ":"
)

func (x *aoc2024) D7P1() int {
	ret := 0

	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to read %s", inputFilePath)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		n := inputToNumbers(scanner.Text())
		n.print()
	}

	return ret
}

type puzzleRow struct {
	goal   int
	values []int
}

func (n *puzzleRow) String() string {
	return fmt.Sprintf("goal: %d, values: %v", n.goal, n.values)
}

func (n *puzzleRow) print() {
	fmt.Println(n.String())
}

func inputToNumbers(rawInput string) *puzzleRow {
	ret := &puzzleRow{}
	// yolo no length checks
	tokens := strings.Split(rawInput, sep)

	// first token is the goal; strip the colon and put it in our return struct
	goalString := tokens[0] // "296:"
	goalString = goalString[:len(goalString)-1]
	ret.goal = stringToInt(goalString)

	// the remaining tokens are the numbers
	ret.values = stringsToInts(tokens[1:])

	return ret
}

func stringsToInts(stringSlice []string) []int {
	intSlice := make([]int, len(stringSlice))
	for i, str := range stringSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("failed to convert %v to []string", stringSlice)
		}
		intSlice[i] = num
	}
	return intSlice
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to convert %d to string", i)
	}
	return i
}
