package player

import "ticTacToe/field"

type Player interface {
	Move(field *field.Field) (x, y int, sign string)
}
