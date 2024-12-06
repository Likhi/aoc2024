package aoc2024

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const inputUrlFormat = "https://adventofcode.com/2024/day/%s/input"

// OnPage reads the entirety of a Day's input and returns it as a string
func OnPage(dayN int) string {
	res, err := http.Get(fmt.Sprintf(inputUrlFormat, strconv.Itoa(dayN)))
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// ReadInput opens the file at path
func ReadInput(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read file at %s", path)
		return nil, err
	}

	return f, nil
}
