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
	instructionBreakChar = ""
)

// a node points to its child nodes
// the child nodes are organized by their pointer in a map so you can access them like node.children[&aChildNode]
type node struct {
	name     string
	children map[*node]struct{}
}

func (x *aoc2024) D5P1() int {

	f, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal("failed to open input/day5.txt")
	}
	defer f.Close()

	// root := node{}

	reader := bufio.NewScanner(f)
	// Because we're generating nodes randomly, we'll have a bunch of separated nodes. Use this map to organize all the node nodes by their root name
	nodes := make(map[string]*node)
	mode := 0 // 0: make nodes, 1: analyze pages
	sum := 0
	for reader.Scan() {
		raw := reader.Bytes()
		if string(raw) == instructionBreakChar {
			mode = 1
			continue
		}

		if mode == 0 { // make nodes
			line := strings.Split(string(raw), "|")
			rootName := line[0]
			childName := line[1]

			// track the root or child if they aren't tracked yet
			if _, ok := nodes[rootName]; !ok {
				nodes[rootName] = newNode(rootName)
			}

			if _, ok := nodes[childName]; !ok {
				nodes[childName] = newNode(childName)
			}

			// point the root to the child
			appendChild(nodes[rootName], nodes[childName])
		} else { //mode ==1, analyze pages
			// printNodesByStringMap(nodes)
			good := true
			pages := strings.Split(string(raw), ",")
			for i, startPage := range pages {
				// fmt.Println("startPage:", startPage)
				for _, nextPage := range pages[i+1:] {
					// fmt.Println("nextPage:", nextPage)
					nextPageIsChildOfStartPage := false
					for childNode := range nodes[startPage].children {
						// fmt.Println(nextPage, "is child of", startPage, "?", "nextPage == childNode.name:", nextPage, "==", childNode.name, "->", nextPage == childNode.name)
						nextPageIsChildOfStartPage = (nextPage == childNode.name)
						if nextPageIsChildOfStartPage {
							break // stop searching child nodes; found a match; all good
						}
					}
					good = good && nextPageIsChildOfStartPage
					if !good {
						break // stop searching the remaining pages following the start page
					}
				}
				if !good {
					// fmt.Println("not good because", startPage)
					break // stop searching this row of pages completely
				}
			}

			if good {
				middleValue, err := strconv.Atoi(pages[len(pages)/2])
				if err != nil {
					log.Fatal("failed to get middle element of row")
				}
				sum += middleValue
			}
		}
	}

	return sum
}

func (x *aoc2024) D5P2() int {
	f, err := os.Open("input/day5.txt")
	if err != nil {
		log.Fatal("failed to open input/day5.txt")
	}
	defer f.Close()

	// root := node{}

	reader := bufio.NewScanner(f)
	// Because we're generating nodes randomly, we'll have a bunch of separated nodes. Use this map to organize all the nodes by their root name
	nodes := make(map[string]*node)
	mode := 0 // 0: make nodes, 1: analyze pages
	sum := 0
	badRows := [][]string{}
	for reader.Scan() {
		raw := reader.Bytes()
		if string(raw) == instructionBreakChar {
			mode = 1
			continue // this is a blank line; get the next input
		}

		if mode == 0 { // make nodes
			line := strings.Split(string(raw), "|")
			rootName := line[0]
			childName := line[1]

			// track the root or child if they aren't tracked yet
			if _, ok := nodes[rootName]; !ok {
				nodes[rootName] = newNode(rootName)
			}

			if _, ok := nodes[childName]; !ok {
				nodes[childName] = newNode(childName)
			}

			// point the root to the child
			appendChild(nodes[rootName], nodes[childName])
		} else { //mode ==1, analyze pages
			// printNodesByStringMap(nodes)
			good := true
			pages := strings.Split(string(raw), ",")
			for i, startPage := range pages {
				// fmt.Println("startPage:", startPage)
				for _, nextPage := range pages[i+1:] {
					// fmt.Println("nextPage:", nextPage)
					nextPageIsChildOfStartPage := false
					for childNode := range nodes[startPage].children {
						// fmt.Println(nextPage, "is child of", startPage, "?", "nextPage == childNode.name:", nextPage, "==", childNode.name, "->", nextPage == childNode.name)
						nextPageIsChildOfStartPage = (nextPage == childNode.name)
						if nextPageIsChildOfStartPage {
							break // stop searching child nodes; found a match; all good
						}
					}
					good = good && nextPageIsChildOfStartPage
					if !good {
						break // stop searching the remaining pages following the start page
					}
				}
				if !good {
					// fmt.Println("not good because", startPage)
					break // stop searching this row of pages completely
				}
			}

			if good {
				middleValue, err := strconv.Atoi(pages[len(pages)/2])
				if err != nil {
					log.Fatal("failed to get middle element of row")
				}
				sum += middleValue
			} else {
				badRows = append(badRows, pages)
			}
		}
	}

	// fmt.Println("old sum:", sum)

	// printNodesByStringMap(nodes)

	// analyze all bad pages
	fixedBadSum := 0
	for _, badRow := range badRows {
		newRow := []string{}
		// fmt.Println("badRow:", badRow)
		for _, page := range badRow {
			if len(newRow) == 0 {
				newRow = append(newRow, page)
			} else {
				for i, newPage := range newRow {
					// if page is not a child of newPage, insert page before newPage
					// else, put page after its last parent and break
					if !nodes[page].isChildOf(nodes[newPage]) {
						temp := append([]string{page}, newRow[i:]...)
						newRow = append(newRow[:i], temp...)
						// newRow = append([]string{page}, newRow...)
						// fmt.Println("page", page, "is not a child of", newPage)
						// fmt.Println("check newrow...", newRow)
						break
					} else {
						if i == len(newRow)-1 { // if we've gotten to the last page, then put page at the end of the newRow
							newRow = append(newRow, page)
							// fmt.Println("attached page", page, "to the end of newRow")
						}
					}

				}
			}

			// fmt.Println("newRow:", newRow)
		}
		// fmt.Println("newRow:", newRow)
		middleValue, err := strconv.Atoi(newRow[len(newRow)/2])
		// fmt.Println("middlevalue:", middleValue)
		if err != nil {
			log.Fatal("failed to get middle element of row")
		}
		fixedBadSum += middleValue

	}

	return fixedBadSum
}

// newNode creates a new node without children
func newNode(rootName string) *node {
	childMap := make(map[*node]struct{})

	root := &node{name: rootName, children: childMap}

	return root
}

func appendChild(root, child *node) {
	if root == nil || child == nil {
		log.Fatal("root or child node is nil, cannot append")
	}

	root.children[child] = struct{}{}
}

func (n *node) isChildOf(parent *node) bool {
	for child := range parent.children {
		if n.name == child.name {
			return true
		}
	}

	return false
}

// prints the names of the nodes stored in the tracking map m
func printNodesByStringMap(m map[string]*node) {
	for name, node := range m {
		fmt.Print(name, ": ")
		for childNode, _ := range node.children {
			fmt.Print(childNode.name, ", ")
		}
		fmt.Println()
	}
}
