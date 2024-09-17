package elemGenerator

import (
	"math/rand"
	"time"
)

type DefaultElemGenerator struct{}

func (d DefaultElemGenerator) Generate() string {
	elements := []string{"😀", "🥵", "🥶", "🤢", "😈"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return elements[r.Intn(len(elements))]
}
