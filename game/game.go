package game

import (
	"fmt"
	"ticTacToe/field"
	"ticTacToe/player"
)

type Game struct {
	players []player.Player
	field   *field.Field
}

func NewGame(players []player.Player) *Game {
	return &Game{
		players: players,
		field:   field.NewField(),
	}
}

func (game *Game) Play() (result string) {
	fmt.Println("Let's play game")

	f := game.field
	for {

		for _, p := range game.players {

			x, y, sign := p.Move(f)

			if err := f.Put(sign, x, y); err != nil {
				return err.Error() //todo repeat moves ???
			}

			fmt.Println(f)

			state := checkState(f, sign)
			if state == Win {
				return sign + " won!"
			} else if state == Draw {
				return "Draw!"
			}
		}

	}
}
