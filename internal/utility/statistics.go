package utility

import (
	"fmt"
	"time"
)

type (
	Statistics interface {
		CurrentScore() int64
		MoveLogs() []string
		AddScore(int64)
	}
)

type GameStatistics struct {
	score    int64
	moveLogs []string
}

func NewGameStatistics() *GameStatistics {
	return &GameStatistics{
		score:    0,
		moveLogs: make([]string, 0),
	}
}

func (s *GameStatistics) CurrentScore() int64 {
	return s.score
}

func (s *GameStatistics) MoveLogs() []string {
	return append([]string{}, s.moveLogs...)
}

func (s *GameStatistics) AddScore(points int64) {
	s.score += points
	log := fmt.Sprintf("[%s] Добавлено %d очков. Текущий счет: %d", time.Now().Format("2006-01-02 15:04:05"), points, s.score)
	s.moveLogs = append(s.moveLogs, log)
}
