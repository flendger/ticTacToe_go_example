package main

import (
	"fmt"
	"ticTacToe/game"
	"ticTacToe/player"
)

func main() {

	fmt.Println("Welcome to tic-tac-toe game!!!")

	players := []player.Player{
		player.NewAi("O"),
		player.NewAnimal("X"),
	}

	g := game.NewGame(players)
	result := g.Play()

	fmt.Println(result)
}
