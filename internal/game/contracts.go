package game

import (
	"three_in_row/internal/field"
	"three_in_row/internal/rules"
)

const (
	RulesApplied Status = iota
	NothingChanged
	Error
)

type (
	Game interface {

		// Run
		// предусловие - созданы ресурсы для запуска игры
		// постусловие - запущена игра, ожидается ввод пользователя
		Run() error
	}

	// логика игры
	// сгенерировано поле
	// применены правила
	// удалены тривряд, начислены очки, посчитаны бонусы
	// сгенерированы новые элменты игрового поля
	// пока не будет достигнуто стабильное состояние на доске, правила продолжат применяться
	// ожидание ввода пользователя
	Engine interface {

		// предусловие - игровое поле не пустое
		// постусловия - изменилось состояние поле
		// постусловия - изменилось состояние очков игрока
		ApplyRules(f field.Field) Status

		// предусловие - на поле есть пустые клетки
		// постусловие - пустые клетки заполнены
		FillField(f field.Field)

		// предусловие - данные для ввода валидны
		// постусловие - измененилось состояние поля
		// постусловие - изменился журнал ходов
		Input(data InputData) error
	}

	EngineConstructor func(field field.Field, rules []rules.GameRule) (Engine, error)

	InputData interface{}

	Status int
)
