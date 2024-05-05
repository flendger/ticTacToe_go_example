package game

import "ticTacToe/field"

type State int

const (
	InProgress State = iota
	Win
	Draw
)

func checkState(field *field.Field, sign string) State {

	if checkWin(field, sign) {
		return Win
	}

	if checkDraw(field) {
		return Draw
	}

	return InProgress
}

func checkWin(f *field.Field, sign string) bool {

	//rows
	win := false
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {

			s, err := f.Sign(x, y)
			if err != nil {
				panic(err)
			}

			win = s == sign
			if !win {
				break
			}
		}

		if win {
			return true
		}
	}

	//cols
	for y := 1; y <= 3; y++ {
		for x := 1; x <= 3; x++ {

			s, err := f.Sign(x, y)
			if err != nil {
				panic(err)
			}

			win = s == sign
			if !win {
				break
			}
		}

		if win {
			return true
		}
	}

	//cross left->right
	for x := 1; x <= 3; x++ {

		s, err := f.Sign(x, x)
		if err != nil {
			panic(err)
		}

		win = s == sign
		if !win {
			break
		}
	}

	if win {
		return true
	}

	//cross right->left
	for x := 1; x <= 3; x++ {

		s, err := f.Sign(x, 4-x)
		if err != nil {
			panic(err)
		}

		win = s == sign
		if !win {
			break
		}
	}

	return win
}

func checkDraw(f *field.Field) bool {

	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			if f.Empty(x, y) {
				return false
			}
		}
	}

	return true
}
