package main

import (
	"crypto/rand"
	"encoding/binary"
)

type Food struct {
	Position Point
}

func NewFood() *Food {
	// Generate a random 32-bit unsigned integer using crypto/rand
	var x, y uint32
	binary.Read(rand.Reader, binary.LittleEndian, &x)
	binary.Read(rand.Reader, binary.LittleEndian, &y)

	// Calculate the position of the food object
	foodX := int(x % uint32(SCREEN_WIDTH/TILE_SIZE))
	foodY := int(y % uint32(SCREEN_HEIGTH/TILE_SIZE))

	return &Food{
		Position: Point{
			X: foodX,
			Y: foodY,
		},
	}
}
