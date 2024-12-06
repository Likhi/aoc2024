package aoc2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	up    = "up"
	down  = "down"
	left  = "left"
	right = "right"
)

type point struct {
	row, col int
	space    string // "." walkable or "#" obstruction
}

type guard struct {
	location *point              // the guard is here
	visited  map[*point]struct{} // the guard as been here already
	facing   string              // left <, right >, up ^, down v // the guard will move this direction on their next move
}

func (x *aoc2024) D6P1() int {
	grid, guard := readInputToGrid()

	printGridWithGuard(grid, guard)
	for !guard.step(grid) {
		printGridWithGuard(grid, guard)
	}

	return len(guard.visited)
}

// readInputToGrid returns a grid and a guard with known starting location (visited starting spot, facing up)
// All points carefully created a pointers so the guard's point is the same as the grid points
func readInputToGrid() ([][]*point, *guard) {
	retGuard := &guard{}

	// read in the grid
	f, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal("failed to open input/day6.txt")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := [][]*point{}
	row := 0
	for scanner.Scan() {
		inputColumns := strings.Split(scanner.Text(), "")
		newRow := []*point{}
		for i, space := range inputColumns {
			newPoint := &point{row: row, col: i, space: space}
			if space == "^" {
				// found guard
				newPoint.space = "." // replace ^ with . for correct pathing later

				retGuard.facing = up
				retGuard.location = newPoint
				retGuard.visited = map[*point]struct{}{newPoint: struct{}{}}
			}

			newRow = append(newRow, newPoint)
		}
		row++
		grid = append(grid, newRow)
	}

	return grid, retGuard
}

// step moves the guard through the grid. If the guard is leaving the grid, return true.
func (g *guard) step(grid [][]*point) bool {
	switch g.facing {
	case up:
		newcol := g.location.col
		newrow := g.location.row - 1
		if newrow < 0 {
			return true
		}

		if grid[newrow][newcol].space == "." {
			g.updateLocation(grid[newrow][newcol])
		}

		if grid[newrow][newcol].space == "#" {
			g.facing = right
		}

	case down:
		newcol := g.location.col
		newrow := g.location.row + 1
		if newrow >= len(grid) {
			return true
		}

		if grid[newrow][newcol].space == "." {
			g.updateLocation(grid[newrow][newcol])
		}

		if grid[newrow][newcol].space == "#" {
			g.facing = left
		}

	case right:
		newcol := g.location.col + 1
		newrow := g.location.row
		if newcol >= len(grid[0]) {
			return true
		}

		if grid[newrow][newcol].space == "." {
			g.updateLocation(grid[newrow][newcol])
		}

		if grid[newrow][newcol].space == "#" {
			g.facing = down
		}

	case left:
		newcol := g.location.col - 1
		newrow := g.location.row
		if newcol < 0 {
			return true
		}

		if grid[newrow][newcol].space == "." {
			g.updateLocation(grid[newrow][newcol])
		}

		if grid[newrow][newcol].space == "#" {
			g.facing = up
		}
	default:
		log.Fatalf("unknown facing direction %v", g.facing)
	}

	return false
}

// uniqueGridPointPointers returns the number of unique point pointers in grid
func uniqueGridPointPointers(grid [][]*point) int {
	d := make(map[*point]struct{})

	for _, row := range grid {
		for _, pt := range row {
			d[pt] = struct{}{}
		}
	}
	return len(d)
}

// gridPoints returns the area (or count of points) in grid
func gridPoints(grid [][]*point) int {
	return len(grid) * len(grid[0])
}

// updateLocation changes the guard's location to point and updates the guard's visited points field
func (g *guard) updateLocation(p *point) {
	g.location = p
	g.addVisited(p)
}

// addVisited adds the point p to the guard's visited points tracker
func (g *guard) addVisited(p *point) {
	g.visited[p] = struct{}{}
}

// hasVisited returns true if the guard as been the point p
func (g *guard) hasVisited(p *point) bool {
	_, ok := g.visited[p]

	return ok
}

// printGridWithGuard prints a grid and historical movements of the guard.
// Visited spaces are marked with 'X' and guard's current location and direction marked by ^, >, v, or <.
func printGridWithGuard(grid [][]*point, g *guard) {
	clearScreen()
	for _, row := range grid {
		for _, pt := range row {
			if pt == g.location {
				switch g.facing {
				case up:
					fmt.Print("^")
				case right:
					fmt.Print(">")
				case down:
					fmt.Print("v")
				case left:
					fmt.Print("<")
				default:
					log.Fatalf("unexpected direction %v", g.facing)
				}
			} else if _, ok := g.visited[pt]; ok {
				fmt.Print("X")
			} else {
				fmt.Print(pt.space) // a . or #
			}

		}
		fmt.Println()
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~", "visited:", len(g.visited))
	fmt.Print("\033[H")
}

// clearScreen clears the terminal window
func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// printGuard pretty-prints info about the guard, optionally the list of visited locations
func printGuard(g *guard, suppressVisited bool) {
	fmt.Printf("guard at: [%d, %d]\n", g.location.row, g.location.col)
	fmt.Println("guard facing:", g.facing)
	if !suppressVisited {
		fmt.Print("guard as visited points:")
		for p := range g.visited {
			fmt.Printf("[%d,%d], ", p.row, p.col)
		}
		fmt.Println("")
	}

}

// printPoint pretty-prints a point as [row,col]
func printPoint(p *point) {
	fmt.Printf("[%d,%d]\n", p.row, p.col)
}
