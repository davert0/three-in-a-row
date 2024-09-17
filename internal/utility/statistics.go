package utility

import (
	"fmt"
	"three_in_row/internal/rules"
	"time"
)

type (
	Statistics interface {
		CurrentScore() int
		MoveLogs() []string
		CountScore([]rules.Combination)
		Log(string)
	}
)

type GameStatistics struct {
	score    int
	moveLogs []string
}

func NewGameStatistics() *GameStatistics {
	return &GameStatistics{
		score:    0,
		moveLogs: make([]string, 0),
	}
}

func (s *GameStatistics) CurrentScore() int {
	return s.score
}

func (s *GameStatistics) Log(st string) {
	s.moveLogs = append(s.moveLogs, fmt.Sprintf("[%s] %s \n", time.Now().Format("2006-01-02 15:04:05"), st))
}

func (s *GameStatistics) CountScore(combinations []rules.Combination) {
	score := 0
	for _, combination := range combinations {
		score += 30

		for i := 4; i <= combination.Length; i++ {
			score += (i - 2) * 10
		}
	}
	s.score += score
	fmt.Printf("[%s] Добавлено %d очков. Текущий счет: %d\n", time.Now().Format("2006-01-02 15:04:05"), score, s.score)

	log := fmt.Sprintf("[%s] Добавлено %d очков. Текущий счет: %d\n", time.Now().Format("2006-01-02 15:04:05"), score, s.score)
	s.moveLogs = append(s.moveLogs, log)
}

func (s *GameStatistics) MoveLogs() []string {
	return append([]string{}, s.moveLogs...)
}
