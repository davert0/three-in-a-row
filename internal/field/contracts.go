package field

type (
	Field[Cell comparable] interface {
		// запросы

		// предусловие - поле не пустое
		Cells() [][]Cell

		// команды

		// предусловие - поле не пустое
		// предусловие - клетки являются соседями
		// постусловие - клетки поменяны местами
		SwapCells(first, second Coord) error

		// предусловие - поле не пустое
		// постусловие - обновлено поле
		UpdateCells([][]Cell) error
	}

	// Constructor
	// предусловие - X, Y положительные
	// постусловие - создано поле размерами X на Y
	Constructor[Cell comparable] func(sizeX, sizeY int) Field[Cell]
)

type (
	Coord struct {
		X, Y int
	}
)
