package main

import (
	"fmt"
	"ticTacToe/field"
	"ticTacToe/player"
)

func main() {

	fmt.Println("Hello into tic-tac-toe game!!!")

	f := field.NewField()
	fmt.Println(f)

	if err := f.Put("X", 1, 1); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f)

	players := []player.Player{player.NewAi()}
	for _, p := range players {

		x, y, sign := p.Move(f)
		err := f.Put(sign, x, y)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(f)
	}
}
