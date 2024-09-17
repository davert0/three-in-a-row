package rules

import (
	"three_in_row/internal/field"
)

const MinMatch = 3

type Combination struct {
	StartRow, StartCol int
	Length             int
	Direction          string // "horizontal" or "vertical"
}

func RemoveCombinations(f field.Field, combinations []Combination) error {
	cells := f.Cells()

	for _, combination := range combinations {
		remove(&cells, combination)
	}

	err := f.UpdateCells(cells)
	if err != nil {
		return err
	}

	return nil
}

func remove(cells *[][]string, combination Combination) {
	row := combination.StartRow
	col := combination.StartCol

	i := 0

	for i < combination.Length {
		(*cells)[row][col] = "✅"
		switch combination.Direction {
		case "horizontal":
			col++
		case "vertical":
			row++
		}
		i++
	}
}

func FindAllCombinations(f field.Field) []Combination {
	var combinations []Combination

	for row := 0; row < len(f.Cells()); row++ {
		combinations = append(combinations, findCombinationsInLine(f, row, 0, 0, 1)...)
	}

	for col := 0; col < len(f.Cells()); col++ {
		combinations = append(combinations, findCombinationsInLine(f, 0, col, 1, 0)...)
	}

	return combinations
}

func findCombinationsInLine(f field.Field, startRow, startCol, rowDelta, colDelta int) []Combination {
	var combinations []Combination

	cells := f.Cells()

	for i := 0; i < len(f.Cells()); {
		row, col := startRow+i*rowDelta, startCol+i*colDelta
		currentElement := cells[row][col]
		count := 1

		for j := 1; i+j < len(f.Cells()); j++ {
			nextRow, nextCol := row+j*rowDelta, col+j*colDelta
			if cells[nextRow][nextCol] == currentElement {
				count++
			} else {
				break
			}
		}

		if count >= MinMatch {
			direction := "horizontal"
			if rowDelta == 1 {
				direction = "vertical"
			}
			combinations = append(combinations, Combination{row, col, count, direction})
			i += count // Skip the entire combination
		} else {
			i++ // Move to the next element
		}
	}

	return combinations
}

func FindCrossCombinations(f field.Field) []Combination {
	var crosses []Combination
	cells := f.Cells()
	fieldSize := len(cells)

	for row := 1; row < len(f.Cells())-1; row++ {
		for col := 1; col < fieldSize-1; col++ {
			currentElement := cells[row][col]
			if currentElement == "" {
				continue
			}

			horizontalCount := 1
			verticalCount := 1

			for i := 1; col-i >= 0 && cells[row][col-i] == currentElement; i++ {
				horizontalCount++
			}
			for i := 1; col+i < fieldSize && cells[row][col+i] == currentElement; i++ {
				horizontalCount++
			}

			for i := 1; row-i >= 0 && cells[row-i][col] == currentElement; i++ {
				verticalCount++
			}
			for i := 1; row+i < fieldSize && cells[row+i][col] == currentElement; i++ {
				verticalCount++
			}

			if horizontalCount >= MinMatch && verticalCount >= MinMatch {
				crosses = append(crosses, Combination{row, col, 0, "cross"})
			}
		}
	}

	return crosses
}
