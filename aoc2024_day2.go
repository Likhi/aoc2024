package aoc2024

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func (x *aoc2024) Day2Part1() (int, error) {
	ret := 0
	sep := " "
	// safeChange := map[int]struct{}{1: {}, 2: {}, 3: {}}

	f, err := ReadInput("input/day2.txt")
	if err != nil {
		return ret, fmt.Errorf("failed to generate day 2 part 1 solution: %w", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	// read report in the file and analyze it
	var done bool
	for !done {
		rawLine, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return ret, fmt.Errorf("failed to read day 2 txt line: %w", err)
			}
		}

		report := strings.Split(string(rawLine), sep)

		fmt.Println(report)

		thisReportSafe, err := isSafe(report)
		if err != nil {
			return ret, err
		}

		if thisReportSafe {
			println("safe")
			ret += 1
		}
	}

	return ret, err
}

func (x *aoc2024) Day2Part2() (int, error) {
	ret := 0
	sep := " "
	// safeChange := map[int]struct{}{1: {}, 2: {}, 3: {}}

	f, err := ReadInput("input/day2.txt")
	if err != nil {
		return ret, fmt.Errorf("failed to generate day 2 part 1 solution: %w", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	// read report in the file and analyze it
	var done bool
	for !done {
		rawLine, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return ret, fmt.Errorf("failed to read day 2 txt line: %w", err)
			}
		}

		report := strings.Split(string(rawLine), sep)
		fmt.Println("original:", report)

		safe := false
		safe, _ = isSafe(report)
		fmt.Println("original should nto change:", report)
		if !safe { // check if safe with removal
			for i := 0; !safe && i < len(report); i++ {
				fmt.Println("i =", i, "remove:", report[i])
				left := report[:i]
				right := report[i+1:]
				new := append([]string{}, left...)
				new = append(new, right...)

				fmt.Println("try this:", new)
				safe, _ = isSafe(new)
			}
		}

		if safe {
			println("safe")
			ret += 1
		} else {
			println("not safe")
		}

		// fmt.Scanln()
	}

	return ret, err
}

// isSafe returns true if the report is safe
func isSafe(report []string) (bool, error) {
	safeChange := map[int]struct{}{1: {}, 2: {}, 3: {}}

	oldLevel := -9999
	wasDirection := 0
	thisReportSafe := true

	for i, levelS := range report {
		newLevel, err := strconv.Atoi(levelS)
		if err != nil {
			return false, err
		}

		if i == 0 {
			oldLevel = newLevel
			continue
		}

		diff := newLevel - oldLevel
		nowDirection := 0
		switch {
		case diff > 0:
			nowDirection = 1
		case diff == 0:
			nowDirection = 0
		default: // diff < 0, decreasing
			nowDirection = -1
		}

		// fmt.Println(wasDirection, nowDirection)

		// compare wasDirection to nowDirection; break if -1 -> 0 or 1, 1 -> 0 or -1, 0 -> 0
		if wasDirection == 0 && nowDirection == 0 {
			fmt.Println("no change", oldLevel, newLevel)
			thisReportSafe = false
		}

		if (wasDirection == -1 && nowDirection == 0) || (wasDirection == -1 && nowDirection == 1) {
			fmt.Println("was decrementing but not anymore", oldLevel, newLevel)
			thisReportSafe = false
		}

		if (wasDirection == 1 && nowDirection == 0) || (wasDirection == 1 && nowDirection == -1) {
			fmt.Println("was incrementing but not anymore", oldLevel, newLevel)
			thisReportSafe = false
		}

		// if the absolute change in value is too large, then break
		absDiff := diff
		if diff < 0 {
			absDiff = absDiff * -1
		}
		if _, ok := safeChange[absDiff]; !ok {
			fmt.Println("change too big or 0:", oldLevel, newLevel)
			thisReportSafe = false
		}

		if !thisReportSafe {
			// skip this value
			break // early break
		} else {
			oldLevel = newLevel         // update for next iteration
			wasDirection = nowDirection // update for next iteration
		}
	}

	return thisReportSafe, nil
}
