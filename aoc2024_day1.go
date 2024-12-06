package aoc2024

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func (x *aoc2024) Day1() (int, error) {
	errValue := -99999
	errFString := "failed to get solution for Day 1: %w"

	f, err := ReadInput("input/day1.txt")
	if err != nil {
		return errValue, fmt.Errorf(errFString, err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	var left []int
	var right []int

	var done bool
	for !done {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return errValue, fmt.Errorf(errFString, err)
			}
		}

		leftRight := strings.Split(string(line), "   ")
		// fmt.Println(leftRight)
		if len(leftRight) < 2 {
			return errValue, fmt.Errorf("failed to split input string on triple-space")
		}

		leftInt, err := strconv.Atoi(leftRight[0])
		if err != nil {
			return errValue, fmt.Errorf("failed to convert input %s to int: %w", leftRight[0], err)
		}

		rightInt, err := strconv.Atoi(leftRight[1])
		if err != nil {
			return errValue, fmt.Errorf("failed to convert input %s to int: %w", leftRight[1], err)
		}

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	if len(left) != len(right) {
		return errValue, fmt.Errorf("left and right input values lengths mismatch: %d != %d", len(left), len(right))
	}
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := range left {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = diff * -1
		}
		sum = sum + diff
	}

	return sum, nil
}

func (x *aoc2024) Day1Part2() (int, error) {
	errValue := -99999
	errFString := "failed to get solution for Day 1: %w"

	f, err := ReadInput("input/day1.txt")
	if err != nil {
		return errValue, fmt.Errorf(errFString, err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	left := make(map[int]int)
	right := make(map[int]int)

	var done bool
	for !done {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return errValue, fmt.Errorf(errFString, err)
			}
		}

		leftRight := strings.Split(string(line), "   ")
		// fmt.Println(leftRight)
		if len(leftRight) < 2 {
			return errValue, fmt.Errorf("failed to split input string on triple-space")
		}

		leftInt, err := strconv.Atoi(leftRight[0])
		if err != nil {
			return errValue, fmt.Errorf("failed to convert input %s to int: %w", leftRight[0], err)
		}

		rightInt, err := strconv.Atoi(leftRight[1])
		if err != nil {
			return errValue, fmt.Errorf("failed to convert input %s to int: %w", leftRight[1], err)
		}

		// for left, just track existence of values
		if _, ok := left[leftInt]; !ok {
			left[leftInt] = 0
		}

		// for right, count occurrence of that value
		if _, ok := right[rightInt]; !ok {
			right[rightInt] = 1
		} else {
			right[rightInt] += 1
		}

	}

	sum := 0
	for kLeft, _ := range left {
		if _, ok := right[kLeft]; ok {
			sum += kLeft * right[kLeft]
			// fmt.Println(sum)
		}
	}

	return sum, nil
}
