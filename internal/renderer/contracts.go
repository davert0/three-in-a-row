package renderer

import (
	"three_in_row/internal/field"
)

type Renderer interface {
	Render(field field.Field)
}
