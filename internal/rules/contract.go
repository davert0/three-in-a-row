package rules

import "three_in_row/internal/field"

type (
	GameRule interface {
		// предусловие - поле инициализировано
		// постусловие - правило применено к полю
		Apply(field field.Field) (bool, error)
	}
)
