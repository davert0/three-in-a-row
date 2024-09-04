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

	// Constructor
	// предусловие - X, Y положительные
	// постусловие - создано поле размерами X на Y
	Constructor[Cell string] func(sizeX, sizeY int) Field
)

type (
	Coord struct {
		X, Y int
	}
)
