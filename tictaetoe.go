package main

import (
	"fmt"
	"ticTacToe/field"
	"ticTacToe/player"
)

func main() {

	fmt.Println("Hello into tic-tac-toe game!!!")

	f := field.NewField()

	players := []player.Player{
		player.NewAi("O"),
		player.NewAi("X"),
	}

	for _, p := range players {

		x, y, sign := p.Move(f)
		err := f.Put(sign, x, y)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(f)
	}
}
