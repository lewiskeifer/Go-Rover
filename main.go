package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Lateral int

const (
	LEFT  Lateral = iota // 0
	RIGHT                // 1
)

type Direction int

const (
	NORTH Direction = iota // 0
	EAST                   // 1
	SOUTH                  // 2
	WEST                   // 3
)

type Rover struct {
	x         int
	y         int
	direction Direction
	path      [][2]int // track all positions visited
}

func (r *Rover) Move(distance int) {
	for step := 0; step < distance; step++ {
		switch r.direction {
		case NORTH:
			r.y++
		case EAST:
			r.x++
		case SOUTH:
			if r.y > 0 {
				r.y--
			}
		case WEST:
			if r.x > 0 {
				r.x--
			}
		}
		r.path = append(r.path, [2]int{r.x, r.y})
	}
}

func (r *Rover) Rotate(newDirection Lateral) {
	if newDirection == RIGHT {
		r.direction = ((r.direction + 1%4) + 4) % 4

	} else {
		r.direction = ((r.direction - 1%4) + 4) % 4
	}
}

func (r *Rover) Print() {
	if len(r.path) == 0 {
		fmt.Println("No path yet")
		return
	}

	minX, maxX := r.path[0][0], r.path[0][0]
	minY, maxY := r.path[0][1], r.path[0][1]
	for _, p := range r.path {
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
		if p[1] < minY {
			minY = p[1]
		}
		if p[1] > maxY {
			maxY = p[1]
		}
	}

	// Build grid
	width := maxX - minX + 1
	height := maxY - minY + 1
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.' // empty space
		}
	}

	// Mark path
	for _, p := range r.path {
		grid[maxY-p[1]][p[0]-minX] = '*' // invert Y for printing
	}

	// Mark rover
	grid[maxY-r.y][r.x-minX] = 'R'

	// Print grid
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	rover := Rover{path: [][2]int{{0, 0}}}
	name := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your command (or 'q' to quit): ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "q" {
			break
		}

		switch name {
		case "m":
			rover.Move(1)
		case "r":
			rover.Rotate(RIGHT)
		case "p":
			rover.Print()
		}
	}

	fmt.Println("Rover is at: ", rover.x, rover.y)
	fmt.Println("Rover is facing: ", Direction(rover.direction))
}
