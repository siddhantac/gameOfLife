package main

type World struct {
	Cells []Cell
}

func CreateWorld() *World {
	return &World{}
}

type Cell struct {
	X, Y int
}

func (w *World) AddCell(cell Cell) {
	w.Cells = append(w.Cells, cell)
}

func (w *World) FindNumNeighbours(cell Cell) int {
	count := 0
	top := Cell{X: 0, Y: cell.Y + 1}
	right := Cell{X: cell.X + 1, Y: 0}
	bottom := Cell{X: 0, Y: cell.Y - 1}
	left := Cell{X: cell.X - 1, Y: 0}
	for _, c := range w.Cells {
		if c == top || c == right || c == left || c == bottom {
			count++
		}
	}

	return count
}
