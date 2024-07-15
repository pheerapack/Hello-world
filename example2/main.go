// robot.go
package main

import (
	"fmt"
)

// Define the size of the grid
const gridSize = 9

// Initial position of the robot
var x, y int

// Function to move the robot up
func moveUp() {
	if y > 0 {
		y--
	}
}

// Function to move the robot down
func moveDown() {
	if y < gridSize-1 {
		y++
	}
}

// Function to move the robot left
func moveLeft() {
	if x > 0 {
		x--
	}
}

// Function to move the robot right
func moveRight() {
	if x < gridSize-1 {
		x++
	}
}

func main() {
	process()
}

func process() (int, int) {
	// Set initial position
	x, y = 4, 4
	fmt.Println("Starting position:", x, y)

	// Example movements
	moveUp()
	moveDown()
	moveLeft()
	moveRight()

	fmt.Println("Final position:", x, y)
	return x, y
}
