package main

import (
	"fmt"
	"time"
)

const (
	width  = 20
	height = 10
)

func main() {
	x, y := 1, 1   // Ball position
	dx, dy := 1, 1 // Ball direction

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Draw the board
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if i == y && j == x {
					fmt.Print("O") // Ball
				} else {
					fmt.Print(".") // Empty space
				}
			}
			fmt.Println()
		}

		// Update ball position
		x += dx
		y += dy

		// Bounce from left/right walls
		if x == 0 || x == width-1 {
			dx = -dx
		}

		// Bounce from top/bottom walls
		if y == 0 || y == height-1 {
			dy = -dy
		}

		// Delay so animation is visible
		time.Sleep(100 * time.Millisecond)
	}
}
