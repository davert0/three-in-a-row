package rules

import (
	"errors"
	"three_in_row/internal/field"
)

type MatchThreeRule struct{}

func NewRules() []GameRule {
	return []GameRule{&MatchThreeRule{}}
}

func (r *MatchThreeRule) Apply(f field.Field) (bool, error) {
	if f == nil {
		return false, errors.New("поле не инициализировано")
	}

	cells := f.Cells()
	if len(cells) == 0 || len(cells[0]) == 0 {
		return false, errors.New("поле пустое")
	}

	changed := false
	newCells := make([][]string, len(cells))
	for i := range newCells {
		newCells[i] = make([]string, len(cells[i]))
		copy(newCells[i], cells[i])
	}

	// Проверяем горизонтальные ряды
	for y := 0; y < len(cells); y++ {
		start := 0
		for x := 1; x <= len(cells[y]); x++ {
			if x == len(cells[y]) || cells[y][x] != cells[y][start] {
				if x-start >= 3 {
					for i := start; i < x; i++ {
						newCells[y][i] = ""
					}
					changed = true
				}
				start = x
			}
		}
	}

	// Проверяем вертикальные столбцы
	for x := 0; x < len(cells[0]); x++ {
		start := 0
		for y := 1; y <= len(cells); y++ {
			if y == len(cells) || cells[y][x] != cells[start][x] {
				if y-start >= 3 {
					for i := start; i < y; i++ {
						newCells[i][x] = ""
					}
					changed = true
				}
				start = y
			}
		}
	}

	if changed {
		return true, f.UpdateCells(newCells)
	}

	return false, nil
}
