package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rover struct {
	x         int
	y         int
	direction string
}

func (r *Rover) Move(dx, dy int) {
	r.x += dx
	r.y += dy
}

func (r *Rover) Rotate(newDirection string) {
	r.direction = newDirection
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
			rover.Rotate("Right")
		}
	}

	fmt.Println("Rover is at: ", rover.x, rover.y)
	fmt.Println("Rover is facing: ", rover.direction)
}
