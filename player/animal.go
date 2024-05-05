package player

import (
	"fmt"
	"ticTacToe/field"
)

type Animal struct {
	sign string
}

func NewAnimal(sign string) *Animal {
	return &Animal{sign: sign}
}

func (a *Animal) Move(field *field.Field) (x, y int, sign string) {

	for {
		if _, err := fmt.Scanln(&x, &y); err != nil {
			panic(err) //todo handle errors ???
		}

		if field.Empty(x, y) {
			break
		}
	}

	return x, y, a.sign
}
