package player

import (
	"math/rand"
	"ticTacToe/field"
	"time"
)

type Ai struct {
	sign string
	rnd  *rand.Rand
}

func NewAi() *Ai {
	return &Ai{
		sign: "O",
		rnd:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (a Ai) Move(f *field.Field) (x, y int, sign string) {

	sign = a.sign

	for {
		x = a.rnd.Intn(3) + 1
		y = a.rnd.Intn(3) + 1

		if f.Empty(x, y) {
			return
		}
	}
}
