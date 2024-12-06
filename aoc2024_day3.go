package aoc2024

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func (x *aoc2024) D3P1() int {
	ret := 0
	f, err := ReadInput("input/day3.txt")
	if err != nil {
		log.Fatalf("failed to generate day 3 part 1 solution: %s", err.Error())
		return ret
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	input := ""
	for {
		inputRaw, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal("failed to read line")
			}
		}
		input += string(inputRaw)
	}

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		match[0] = strings.Replace(match[0], "(", "\\(", -1)
		match[0] = strings.Replace(match[0], ")", "\\)", -1)
		re2 := regexp.MustCompile(match[0])
		if re2OK := re2.MatchString(input); !re2OK {
			log.Fatalf("match is fake?")
		}

		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		ret += (a * b)
	}

	return ret
}

func (x *aoc2024) D3P2() int {
	ret := 0
	f, err := ReadInput("input/day3.txt")
	if err != nil {
		log.Fatalf("failed to generate day 3 part 1 solution: %s", err.Error())
		return ret
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	mulstate, dostate, dontstate := 0, 0, 0
	multiplicationEn := true
	a, b := "", ""
	reNumber := regexp.MustCompile(`[0-9]`)
	for {
		inputRaw, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal("failed to read line")
			}
		}

		// define mul states:
		// 0 - searching
		// 1 - m
		// 2 - mu
		// 3 - mul
		// 4 - mul(
		// 5 - mul(x
		// 6 - mul(xx *optional
		// 7 - mul(xxx *optional
		// 8 - mul(xxx,
		// 9 - mul(xxx,y
		// 10 - mul(xxx,yy *optional
		// 11 - mul(xxx,yyy *optional
		// 12 - mul(xxx,yyy)
		// at 12, multiply xxx and yyy and add the product to summation variable 'ret'

		// define do states:
		// 0 - searching
		// 1 - d
		// 2 - do
		// 3 - do(
		// 4 - do()
		// at 4, enable multiplication

		// define don't states:
		// 0 - searching
		// 1 - d
		// 2 - do
		// 3 - don
		// 4 - don'
		// 5 - don't
		// 6 - don't(
		// 7 - don't()
		// at 7, disable multiplication

		in := string(inputRaw)

		// update mulstate
		switch mulstate {
		case 0: // searching
			if in == "m" {
				mulstate = 1
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 1: // m
			if in == "u" {
				mulstate = 2
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 2: // mu
			if in == "l" {
				mulstate = 3
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 3: // mul
			if in == "(" {
				mulstate = 4
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 4: // mul(
			if reNumber.MatchString(in) { // -> mul(x
				mulstate = 5
				a += in
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 5: // mul(x
			if reNumber.MatchString(in) { // -> mul(xx
				mulstate = 6
				a += in
			} else if in == "," { // -> mul(x,
				mulstate = 8
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 6: // mul(xx
			if reNumber.MatchString(in) { // -> mul(xxx
				mulstate = 7
				a += in
			} else if in == "," { // -> mul(xx,
				mulstate = 8
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 7: // mul(xxx
			if in == "," { // -> mul(xxx,
				mulstate = 8
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 8: // mul(xxx,
			if reNumber.MatchString(in) { // -> mul(xxx,y
				mulstate = 9
				b += in
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 9: // mul(xxx,y
			if reNumber.MatchString(in) { // -> mul(xxx,yy
				mulstate = 10
				b += in
			} else if in == ")" { // -> mul(xxx,y)
				mulstate = 12
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 10: // mul(xxx,yy
			if reNumber.MatchString(in) { // -> mul(xxx,yyy
				mulstate = 11
				b += in
			} else if in == ")" { // -> mul(xxx,yy)
				mulstate = 12
			} else {
				mulstate = 0
				a, b = "", ""
			}
		case 11: // mul(xxx,yyy
			if in == ")" { // -> mul(xxx,yyy)
				mulstate = 12
			} else {
				mulstate = 0
				a, b = "", ""
			}
		default:
			log.Fatalf("unexpected mulstate %d", mulstate)
		}
		if mulstate == 12 {
			if multiplicationEn {
				ai, _ := strconv.Atoi(a)
				bi, _ := strconv.Atoi(b)
				ret += ai * bi
			}
			mulstate = 0
			a, b = "", ""
		}

		// update dostate
		switch dostate {
		case 0: // searching
			if in == "d" { // -> d
				dostate = 1
			} else {
				dostate = 0
			}
		case 1: // d
			if in == "o" { // -> do
				dostate = 2
			} else {
				dostate = 0
			}
		case 2:
			if in == "(" { // -> do(
				dostate = 3
			} else {
				dostate = 0
			}
		case 3:
			if in == ")" { // -> do()
				dostate = 4
			} else {
				dostate = 0
			}
		default:
			log.Fatalf("unexpected dostate %d", dostate)
		}
		if dostate == 4 {
			multiplicationEn = true
			dostate = 0
		}

		// update dontstate
		switch dontstate {
		case 0: // searching
			if in == "d" { // -> d
				dontstate = 1
			} else {
				dontstate = 0
			}
		case 1: // d
			if in == "o" { // -> do
				dontstate = 2
			} else {
				dontstate = 0
			}
		case 2:
			if in == "n" { // -> don
				dontstate = 3
			} else {
				dontstate = 0
			}
		case 3:
			if in == "'" { // -> don'
				dontstate = 4
			} else {
				dontstate = 0
			}
		case 4:
			if in == "t" { // -> don't
				dontstate = 5
			} else {
				dontstate = 0
			}
		case 5:
			if in == "(" { // -> don't(
				dontstate = 6
			} else {
				dontstate = 0
			}
		case 6:
			if in == ")" { // -> don't()
				dontstate = 7
			} else {
				dontstate = 0
			}
		default:
			log.Fatalf("unexpected dontstate %d", dontstate)
		}
		if dontstate == 7 {
			multiplicationEn = false
			dontstate = 0
		}
	}

	return ret
}
