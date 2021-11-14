package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWorld(t *testing.T) {
	world := CreateWorld()
	assert.Equal(t, 0, len(world.Cells))
}

func TestAddCell(t *testing.T) {
	world := CreateWorld()
	world.AddCell(Cell{X: 2, Y: 3})
	assert.Equal(t, 1, len(world.Cells))
}

func TestFindNumNeighbours(t *testing.T) {
	world := CreateWorld()

	cell := Cell{X: 0, Y: 0}
	world.AddCell(cell)
	world.AddCell(Cell{X: 0, Y: 1})
	world.AddCell(Cell{X: 1, Y: 0})
	world.AddCell(Cell{X: -1, Y: 0})
	world.AddCell(Cell{X: 0, Y: -1})
	world.AddCell(Cell{X: 2, Y: 3})

	n := world.FindNumNeighbours(cell)
	assert.Equal(t, 4, n)
}
