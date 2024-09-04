package renderer

import (
	"fmt"
	"strings"
	"three_in_row/internal/field"
)

type ConsoleRenderer struct{}

func NewConsoleRenderer() *ConsoleRenderer {
	return &ConsoleRenderer{}
}

func (r *ConsoleRenderer) Render(f field.Field) {
	cells := f.Cells()
	if len(cells) == 0 {
		fmt.Println("Поле пустое")
		return
	}

	// Определяем максимальную ширину ячейки для выравнивания
	maxWidth := 1
	for _, row := range cells {
		for _, cell := range row {
			if len(cell) > maxWidth {
				maxWidth = len(cell)
			}
		}
	}

	// Выводим верхнюю границу
	fmt.Println("+" + strings.Repeat("-", len(cells[0])*(maxWidth+1)+1) + "+")

	// Выводим содержимое поля
	for _, row := range cells {
		fmt.Print("|")
		for _, cell := range row {
			if cell == "" {
				fmt.Printf("%-*s|", maxWidth+1, " ")
			} else {
				fmt.Printf("%-*s|", maxWidth+1, cell)
			}
		}
		fmt.Println()
	}

	// Выводим нижнюю границу
	fmt.Println("+" + strings.Repeat("-", len(cells[0])*(maxWidth+1)+1) + "+")
}
