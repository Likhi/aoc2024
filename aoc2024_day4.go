package aoc2024

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func (x *aoc2024) D4P1() int {
	f, err := ReadInput("input/day4.txt")
	if err != nil {
		log.Fatalf("failed to open file")
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	input := [][]string{}

	// read input the whole puzzle input into 2D slice
	for {
		rawLine, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("failed to read line")
			}
		}

		input = append(input, strings.Split(string(rawLine), ""))
	}

	count := 0

	// first check left-right and right-left matches in each row ↔️
	for _, row := range input {
		count += countXMASInSlice(row)
	}

	// next check top-bottom and bottom-top matches in each column ↕️
	for col := 0; col < len(input[0]); col++ {
		column := []string{}
		for row := 0; row < len(input); row++ {
			column = append(column, input[row][col])
		}

		count += countXMASInSlice(column)
	}

	// check diagonals top right to bottom left
	trbl := [][]string{}
	rows := len(input)
	cols := len(input[0])
	for d := 0; d < rows+cols-1; d++ {
		var diagonal []string

		for i := 0; i < rows; i++ {
			j := d - i
			if j >= 0 && j < cols {
				diagonal = append(diagonal, input[i][j])
			}
		}
		if len(diagonal) > 0 {
			trbl = append(trbl, diagonal)
		}
	}

	for _, s := range trbl {
		count += countXMASInSlice(s)
	}

	// check diagonals top-left to bottom-right
	tlbr := [][]string{}
	for d := 0; d < rows+cols-1; d++ {
		var diagonal []string
		for i := 0; i < rows; i++ {
			j := d - i
			if j >= 0 && j < cols {
				diagonal = append(diagonal, input[i][cols-1-j])
			}
		}
		if len(diagonal) > 0 {
			tlbr = append(tlbr, diagonal)
		}
	}

	for _, s := range tlbr {
		count += countXMASInSlice(s)
	}

	return count
}

func (x *aoc2024) D4P2() int {
	f, err := ReadInput("input/day4.txt")
	if err != nil {
		log.Fatalf("failed to open file")
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	input := [][]string{}

	// read input the whole puzzle input into 2D slice
	for {
		rawLine, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("failed to read line")
			}
		}

		input = append(input, strings.Split(string(rawLine), ""))
	}

	count := 0

	height := len(input)
	width := len(input[0])

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			center := input[i][j]
			topleft := input[i-1][j-1]
			topright := input[i-1][j+1]
			bottomleft := input[i+1][j-1]
			bottomright := input[i+1][j+1]

			x := topleft + topright + center + bottomleft + bottomright
			/*
				M S     M M     S S     S M
				 A       A       A       A
				M S     S S     M M     S M
				MSAMS   MMASS   SSAMM   SMASM
			*/
			if x == "MSAMS" || x == "MMASS" || x == "SSAMM" || x == "SMASM" {
				count += 1
			}
		}
	}

	return count
}

// countXMASInSlice counts occurrences of XMAS or SAMX in the slice
func countXMASInSlice(in []string) int {
	return strings.Count(strings.Join(in, ""), "XMAS") + strings.Count(strings.Join(in, ""), "SAMX")
}

func reverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// identical returns true if the elemts of the two slices are identical.
// If the slices are not the same length, the result is false.
func identical(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	ret := true
	for i := range a {
		ret = ret && (a[i] == b[i])
	}

	return ret
}
