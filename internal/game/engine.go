package game

import (
	"math/rand"
	"three_in_row/internal/field"
	"three_in_row/internal/rules"
	"time"
)

type (
	GameEngine struct {
		rules []rules.GameRule
	}
)

func NewGameEngine(rules []rules.GameRule) *GameEngine {
	return &GameEngine{rules: rules}
}

func (g *GameEngine) ApplyRules(f field.Field) Status {
	for _, rule := range g.rules {
		changed, err := rule.Apply(f)
		if err != nil {
			return Error
		}
		if changed {
			return RulesApplied
		}
	}
	return NothingChanged
}

func (g *GameEngine) FillField(f field.Field) {
	cells := f.Cells()
	if len(cells) == 0 || len(cells[0]) == 0 {
		return // Поле пустое, нечего заполнять
	}

	elements := []string{"😀", "🥵", "🥶", "🤢", "😈"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	newCells := make([][]string, len(cells))
	for y := range cells {
		newCells[y] = make([]string, len(cells[y]))
		for x := range cells[y] {
			if cells[y][x] == "" {
				newCells[y][x] = elements[r.Intn(len(elements))]
			} else {
				newCells[y][x] = cells[y][x]
			}
		}
	}

	err := f.UpdateCells(newCells)
	if err != nil {
		// Обработка ошибки, если не удалось обновить поле
		// В данном случае мы просто логируем ошибку, но вы можете обработать её по-другому
		println("Ошибка при обновлении поля:", err.Error())
	}
}

func (g *GameEngine) Input(data InputData) error {
	//TODO implement me
	panic("implement me")
}
