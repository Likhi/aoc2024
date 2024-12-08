package aoc2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const showGrid = true

type direction string

const (
	up    direction = "up"
	down  direction = "down"
	left  direction = "left"
	right direction = "right"
)

type point struct {
	row, col  int
	spaceType string // "." walkable or "#" obstruction or "O" test obstruction
}

type guard struct {
	location *point               // the guard is here
	visited  map[*point]direction // the guard as been here already and was pointing some direction
	facing   direction            // left <, right >, up ^, down v // the guard will move this direction on their next move
}

func (x *aoc2024) D6P1() int {
	grid, guard, _ := readInputToGrid()

	if showGrid {
		printGridWithGuard(grid, guard)
	}

	for {
		done, err := guard.step(grid)
		if err != nil {
			log.Fatal("deadlock detected")
		}

		if showGrid {
			printGridWithGuard(grid, guard)
		}

		if done {
			break
		}
	}

	return len(guard.visited)
}

func (x *aoc2024) D6P2() int {
	grid, guard, guardStart := readInputToGrid()
	lastSpaceType := ""

	// put an obstruction "O" on all points next to guard and south of guard and check if deadlock
	deadlocks := 0
	for _, rows := range grid[guard.location.row:] {
		for _, pt := range rows {
			lastSpaceType = pt.spaceType
			if pt.spaceType != "^" && pt.spaceType != "#" {
				pt.spaceType = "O" // place an "O"...
			} else {
				continue
			}

			// if showGrid {
			// 	printGridWithGuard(grid, guard)
			// }

			// then attempt the simulation
			for {
				done, err := guard.step(grid)
				if err != nil && err.Error() == "Deadlock" {
					// fmt.Println("err:", err) // assuming this is only deadlock error :)
					fmt.Printf("deadlock with 'O' at [%d,%d]\n", pt.row, pt.col)
					deadlocks++
					printGridWithGuard(grid, guard)
					time.Sleep(2 * time.Second)
					break
				}

				// if showGrid {
				// 	printGridWithGuard(grid, guard)
				// }

				if done {
					break
				}
			}
			pt.spaceType = lastSpaceType // revert the "O"

			// reset the guard
			guard.facing = up
			guard.location = guardStart
			guard.visited = map[*point]direction{guardStart: up}

			// fmt.Println("deadlocks", deadlocks)
		}

	}

	return deadlocks
}

// readInputToGrid returns a grid and a guard with known starting location (visited starting spot, facing up)
// All points carefully created a pointers so the guard's point is the same as the grid points
func readInputToGrid() ([][]*point, *guard, *point) {
	retGuard := &guard{}
	retStartPoint := &point{}

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
			newPoint := &point{row: row, col: i, spaceType: space}
			if space == "^" {
				// found guard
				newPoint.spaceType = "." // replace ^ with . for correct pathing later

				retGuard.facing = up
				retGuard.location = newPoint
				retGuard.visited = map[*point]direction{newPoint: up}
				retStartPoint = newPoint
			}

			newRow = append(newRow, newPoint)
		}
		row++
		grid = append(grid, newRow)
	}

	return grid, retGuard, retStartPoint
}

// step moves the guard through the grid. If the guard is leaving the grid, return true.
func (g *guard) step(grid [][]*point) (bool, error) {
	switch g.facing {
	case up:
		nextcol := g.location.col
		nextrow := g.location.row - 1
		if nextrow < 0 {
			return true, nil
		}

		// detect deadlock (if guard was at grid[nextrow][nextcol] with same direction already, it's deadlock)
		if _, beenThere := g.visited[grid[nextrow][nextcol]]; beenThere {
			wasFacing := g.visited[grid[nextrow][nextcol]]
			if wasFacing == g.facing {
				return false, fmt.Errorf("Deadlock")
			}
		}

		// update new location arrival
		if grid[nextrow][nextcol].spaceType == "." {
			g.updateLocation(grid[nextrow][nextcol], up)
		}

		if grid[nextrow][nextcol].spaceType == "#" || grid[nextrow][nextcol].spaceType == "O" {
			g.facing = right
		}

	case down:
		nextcol := g.location.col
		nextrow := g.location.row + 1
		if nextrow >= len(grid) {
			return true, nil
		}

		// detect deadlock (if guard was at grid[nextrow][nextcol] with same direction already, it's deadlock)
		if _, beenThere := g.visited[grid[nextrow][nextcol]]; beenThere {
			wasFacing := g.visited[grid[nextrow][nextcol]]
			if wasFacing == down {
				return false, fmt.Errorf("Deadlock")
			}
		}

		if grid[nextrow][nextcol].spaceType == "." {
			g.updateLocation(grid[nextrow][nextcol], down)
		}

		if grid[nextrow][nextcol].spaceType == "#" || grid[nextrow][nextcol].spaceType == "O" {
			g.facing = left
		}

	case right:
		nextcol := g.location.col + 1
		nextrow := g.location.row
		if nextcol >= len(grid[0]) {
			return true, nil
		}

		// detect deadlock (if guard was at grid[nextrow][nextcol] with same direction already, it's deadlock)
		if _, beenThere := g.visited[grid[nextrow][nextcol]]; beenThere {
			wasFacing := g.visited[grid[nextrow][nextcol]]
			if wasFacing == right {
				return false, fmt.Errorf("Deadlock")
			}
		}

		if grid[nextrow][nextcol].spaceType == "." {
			g.updateLocation(grid[nextrow][nextcol], right)
		}

		if grid[nextrow][nextcol].spaceType == "#" || grid[nextrow][nextcol].spaceType == "O" {
			g.facing = down
		}

	case left:
		nextcol := g.location.col - 1
		nextrow := g.location.row
		if nextcol < 0 {
			return true, nil
		}

		// detect deadlock (if guard was at grid[nextrow][nextcol] with same direction already, it's deadlock)
		if _, beenThere := g.visited[grid[nextrow][nextcol]]; beenThere {
			wasFacing := g.visited[grid[nextrow][nextcol]]
			if wasFacing == left {
				return false, fmt.Errorf("Deadlock")
			}
		}

		if grid[nextrow][nextcol].spaceType == "." {
			g.updateLocation(grid[nextrow][nextcol], left)
		}

		if grid[nextrow][nextcol].spaceType == "#" || grid[nextrow][nextcol].spaceType == "O" {
			g.facing = up
		}
	default:
		log.Fatalf("unknown facing direction %v", g.facing)
	}

	return false, nil
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
func (g *guard) updateLocation(p *point, d direction) {
	g.location = p
	g.addVisited(p, d)
}

// addVisited adds the point p to the guard's visited points tracker
func (g *guard) addVisited(p *point, d direction) {
	g.visited[p] = d
}

// hasVisited returns true if the guard as been the point p
func (g *guard) hasVisited(p *point) bool {
	_, ok := g.visited[p]

	return ok
}

// printGrid prints the grid.
func printGrid(grid [][]*point, g *guard) {
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
			} else {
				fmt.Print(pt.spaceType) // a . or #
			}

		}
		fmt.Println()
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~", "visited:", len(g.visited))
	fmt.Print("\033[H")
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
				fmt.Print(pt.spaceType) // a . or #
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
