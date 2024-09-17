package renderer

import (
	"fmt"
	"strings"
	"three_in_row/internal/field"
	"time"
)

type ConsoleRenderer struct{}

func NewConsoleRenderer() *ConsoleRenderer {
	return &ConsoleRenderer{}
}

func (r *ConsoleRenderer) Render(field field.Field) {
	fmt.Println("*******************************************")

	cells := field.Cells()
	fmt.Println(r.renderHorizontalBorder(len(cells[0]), true))
	for i, row := range cells {
		fmt.Print(i)
		for _, cell := range row {
			fmt.Printf(" %s  |", cell)
		}
		fmt.Println()
		fmt.Println(r.renderHorizontalBorder(len(cells[0]), false))
	}

	time.Sleep(1 * time.Second)

}

// renderHorizontalBorder renders a horizontal border of the given length.
func (r *ConsoleRenderer) renderHorizontalBorder(length int, withCoords bool) string {
	// Изменена ширина ячейки для 8x8 поля
	b := strings.Builder{}
	b.WriteString("+")
	for i := 0; i < length; i++ {
		if withCoords {
			b.WriteString(fmt.Sprintf("--%d--+", i))
		} else {
			b.WriteString("-----+")
		}

	}
	return b.String()
}
