package main

import "testing"

func TestMoveRight(t *testing.T) {
	x, y = 4, 4
	process()

	expectedX, expectedY := 1, 1
	if x != expectedX || y != expectedY {
		t.Errorf("moveRight failed, expected position (%d, %d), got (%d, %d)", expectedX, expectedY, x, y)
	}
}
