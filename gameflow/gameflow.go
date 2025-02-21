package gameflow

import (
	"reflect"
	"time"

	"github.com/laurensdc/go-game-of-life/life"
	"github.com/laurensdc/go-game-of-life/print"
)

func StartGame(t *time.Ticker, p print.GamePrinter) {
	defer t.Stop()

	for range t.C {
		newState := life.Tick(state)

		if reflect.DeepEqual(state, newState) {
			break // game over because game is not evolving anymome
		}

		p.ClearScreen()

		state = newState

		p.PrintGameState(state)
	}
}
