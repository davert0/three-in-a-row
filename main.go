package main

import (
	"three_in_row/internal/field"
	"three_in_row/internal/game"
	"three_in_row/internal/renderer"
	"three_in_row/internal/utility"
)

func main() {
	stats := utility.NewGameStatistics()

	renderer_ := renderer.NewConsoleRenderer()
	field_ := field.NewField(8, 8)
	newGame := game.NewGame(field_, renderer_, stats)

	err := newGame.Run()
	if err != nil {
		panic(err)
	}

}
