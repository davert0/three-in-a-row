package field

import (
	"errors"
)

type (
	fieldImpl struct {
		cells [][]string
	}
)

func NewField(sizeX, sizeY int) Field {
	if sizeX <= 0 || sizeY <= 0 {
		panic("размеры поля должны быть положительными")
	}

	cells := make([][]string, sizeY)
	for i := range cells {
		cells[i] = make([]string, sizeX)
	}

	return &fieldImpl{cells: cells}
}

func (f *fieldImpl) Cells() [][]string {
	if len(f.cells) == 0 {
		return nil
	}
	return f.cells
}

func (f *fieldImpl) SwapCells(first, second Coord) error {
	if len(f.cells) == 0 {
		return errors.New("поле пустое")
	}

	if !f.areNeighbors(first, second) {
		return errors.New("клетки не являются соседями")
	}

	if !f.isValidCoord(first) || !f.isValidCoord(second) {
		return errors.New("недопустимые координаты")
	}

	f.cells[first.Y][first.X], f.cells[second.Y][second.X] = f.cells[second.Y][second.X], f.cells[first.Y][first.X]
	return nil
}

func (f *fieldImpl) UpdateCells(newCells [][]string) error {
	if len(f.cells) == 0 {
		return errors.New("поле пустое")
	}

	if len(newCells) != len(f.cells) || len(newCells[0]) != len(f.cells[0]) {
		return errors.New("размеры нового поля не совпадают с текущим")
	}

	f.cells = newCells
	return nil
}

func (f *fieldImpl) areNeighbors(c1, c2 Coord) bool {
	dx := abs(c1.X - c2.X)
	dy := abs(c1.Y - c2.Y)
	return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
}

func (f *fieldImpl) isValidCoord(c Coord) bool {
	return c.X >= 0 && c.X < len(f.cells[0]) && c.Y >= 0 && c.Y < len(f.cells)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
