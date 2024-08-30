package rules

import "three_in_row/internal/field"

type (
	GameRule[Cell comparable] interface {
		// предусловие - поле инициализировано
		// постусловие - правило применено к полю
		Apply(field *field.Field[Cell]) error
	}
)
