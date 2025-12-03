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
	RIGHT                // 1 	// 3
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
	actions   []int
}

func (r *Rover) Move(dx, dy int) {
	r.x += dx
	r.y += dy
}

func (r *Rover) Rotate(newDirection Lateral) {
	if newDirection == RIGHT {
		r.direction = ((r.direction + 1%4) + 4) % 4

	} else {
		r.direction = ((r.direction - 1%4) + 4) % 4
	}
}

func (r *Rover) Print() {
	for i := 0; i < r.x; i++ {
		for j := 0; j < r.y; j++ {
			fmt.Print("-")
		}
		fmt.Println("x")
	}
}

func main() {

	rover := Rover{}
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
			rover.Move(1, 1)
		case "r":
			rover.Rotate(RIGHT)
		case "p":
			rover.Print()
		}
	}

	fmt.Println("Rover is at: ", rover.x, rover.y)
	fmt.Println("Rover is facing: ", rover.direction)
}
