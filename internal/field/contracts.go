package field

type (
	Field interface {
		// запросы

		// предусловие - поле не пустое
		Cells() [][]string

		// команды

		// предусловие - поле не пустое
		// предусловие - клетки являются соседями
		// постусловие - клетки поменяны местами
		SwapCells(first, second Coord) error

		// предусловие - поле не пустое
		// постусловие - обновлено поле
		UpdateCells([][]string) error
	}
)

type (
	Coord struct {
		X, Y int
	}
)
