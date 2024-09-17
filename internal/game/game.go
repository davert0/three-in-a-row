package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"three_in_row/internal/elemGenerator"
	"three_in_row/internal/field"
	"three_in_row/internal/renderer"
	"three_in_row/internal/rules"
	"three_in_row/internal/utility"
)

const emptyCell = "✅"

type (
	GameImpl struct {
		field         field.Field
		renderer      renderer.Renderer
		elemGenerator elemGenerator.ElemGenerator
		statistics    utility.Statistics
		crossBonuses  int
	}
)

func NewGame(field field.Field, renderer renderer.Renderer, statistics utility.Statistics) Game {
	generator := elemGenerator.DefaultElemGenerator{}
	g := &GameImpl{
		field:         field,
		renderer:      renderer,
		statistics:    statistics,
		elemGenerator: generator,
	}
	return g
}

func (g *GameImpl) Run() error {
	g.fillField()
	g.renderer.Render(g.field)

	for {
		if g.isEndGame() {
			fmt.Println("Игра оконччена")
			return nil
		}
		err := g.run()
		if err != nil {
			return err
		}
	}

}

func (g *GameImpl) run() error {
	combinations := rules.FindAllCombinations(g.field)

	if len(combinations) == 0 {

		c1, c2, err := g.Input()
		if err != nil {
			fmt.Println(err)
			return nil
		}

		err = g.field.SwapCells(c1, c2)
		if err != nil {
			fmt.Println(err)
		}

		combinations = rules.FindAllCombinations(g.field)
		if len(combinations) == 0 {
			_ = g.field.SwapCells(c1, c2)
			fmt.Println("Ход невозможен")
		}

		g.renderer.Render(g.field)
		return nil
	}

	crosses := rules.FindCrossCombinations(g.field)
	g.addCrossBonuses(len(crosses))
	g.statistics.CountScore(combinations)
	err := rules.RemoveCombinations(g.field, combinations)
	if err != nil {
		return err
	}

	g.renderer.Render(g.field)
	err = g.floatUpEmptyCells()
	if err != nil {
		return err
	}
	g.fillField()
	g.renderer.Render(g.field)
	return nil
}

func (g *GameImpl) isEndGame() bool {
	rows := len(g.field.Cells())
	cols := len(g.field.Cells()[0])

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if x+1 < cols {
				first := field.Coord{X: x, Y: y}
				second := field.Coord{X: x + 1, Y: y}
				err := g.field.SwapCells(first, second)
				if err != nil {
					continue
				}
				comb := rules.FindAllCombinations(g.field)
				_ = g.field.SwapCells(first, second)

				if len(comb) > 0 {
					return false
				}

			}

			if y+1 < rows {
				first := field.Coord{X: x, Y: y}
				second := field.Coord{X: x, Y: y + 1}
				err := g.field.SwapCells(first, second)
				if err != nil {
					continue
				}
				comb := rules.FindAllCombinations(g.field)
				_ = g.field.SwapCells(first, second)
				if len(comb) > 0 {
					return false
				}
			}
		}
	}
	return true
}

func (g *GameImpl) addCrossBonuses(n int) {
	if n != 0 {
		fmt.Printf("Получено бонусов за крестовину: %d\n", n)
	}
	g.crossBonuses += n
}

func (g *GameImpl) fillField() {
	cells := g.field.Cells()
	if len(cells) == 0 || len(cells[0]) == 0 {
		return // Поле пустое, нечего заполнять
	}

	newCells := make([][]string, len(cells))
	for y := range cells {
		newCells[y] = make([]string, len(cells[y]))
		for x := range cells[y] {
			if cells[y][x] != "" && cells[y][x] != emptyCell {
				newCells[y][x] = cells[y][x]
				continue
			}
			newCells[y][x] = g.elemGenerator.Generate()
		}
	}

	err := g.field.UpdateCells(newCells)
	if err != nil {
		fmt.Println("Ошибка при обновлении поля:", err.Error())
	}
}

func (g *GameImpl) floatUpEmptyCells() error {
	cells := g.field.Cells()
	if len(cells) == 0 || len(cells[0]) == 0 {
		return errors.New("Поле пустое")
	}

	rows := len(cells)
	cols := len(cells[0])

	for col := 0; col < cols; col++ {
		fillPosition := rows - 1
		for row := rows - 1; row >= 0; row-- {
			if cells[row][col] != emptyCell {
				cells[fillPosition][col] = cells[row][col]
				fillPosition--
			}
		}
		for fillPosition >= 0 {
			cells[fillPosition][col] = emptyCell
			fillPosition--
		}
	}

	err := g.field.UpdateCells(cells)
	if err != nil {
		return fmt.Errorf("error updating cells: %v", err)
	}

	return nil
}

func (c *GameImpl) Input() (field.Coord, field.Coord, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите координаты соседних клеток в формате: \"0 0, 0 1\"\nДля получения журнала событий введите \"журнал\"")

	input, err := reader.ReadString('\n')
	if err != nil {
		return field.Coord{}, field.Coord{}, fmt.Errorf("ошибка чтения ввода: %v", err)
	}

	input = strings.TrimSpace(input)

	if input == "журнал" {

		return field.Coord{}, field.Coord{}, fmt.Errorf("Запись журнала событий: \n %v", c.statistics.MoveLogs())
	}

	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return field.Coord{}, field.Coord{}, fmt.Errorf("ввод некорректен, требуется две координаты, разделенные запятой")
	}

	firstCoord, err := parseCoord(parts[0])
	if err != nil {
		return field.Coord{}, field.Coord{}, fmt.Errorf("ошибка при разборе первой координаты: %v", err)
	}

	secondCoord, err := parseCoord(parts[1])
	if err != nil {
		return field.Coord{}, field.Coord{}, fmt.Errorf("ошибка при разборе второй координаты: %v", err)
	}
	c.statistics.Log(fmt.Sprintf("Был совершен ход: %s, %s", parts[0], parts[1]))
	c.statistics.CurrentScore()
	return firstCoord, secondCoord, nil
}

func parseCoord(input string) (field.Coord, error) {
	trimmed := strings.TrimSpace(input)
	xy := strings.Split(trimmed, " ")
	if len(xy) != 2 {
		return field.Coord{}, fmt.Errorf("координата должна состоять из двух чисел, разделённых пробелом")
	}

	x, err := strconv.Atoi(xy[0])
	if err != nil {
		return field.Coord{}, fmt.Errorf("ошибка при разборе X: %v", err)
	}

	y, err := strconv.Atoi(xy[1])
	if err != nil {
		return field.Coord{}, fmt.Errorf("ошибка при разборе Y: %v", err)
	}

	return field.Coord{X: x, Y: y}, nil
}
