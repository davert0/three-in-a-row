package game

import (
	"three_in_row/internal/field"
	"three_in_row/internal/renderer"
	"three_in_row/internal/utility"
	"time"
)

type (
	GameImpl struct {
		engine     Engine
		field      field.Field
		renderer   renderer.Renderer
		statistics utility.Statistics
	}
)

func NewGame(engine Engine, field field.Field, renderer renderer.Renderer) Game {
	g := &GameImpl{engine: engine, field: field, renderer: renderer}
	return g
}

func (g *GameImpl) Run() error {
	g.engine.FillField(g.field)
	g.renderer.Render(g.field)

	// пока не останется правил, применяем их
	for {
		status := g.engine.ApplyRules(g.field)
		g.renderer.Render(g.field)
		time.Sleep(3 * time.Second)
		switch status {
		case NothingChanged:
			break
		case Error:
			panic(status)
		default:
			continue
		}
	}

	err := g.engine.Input("s")
	if err != nil {
		return err
	}

	return nil
}
