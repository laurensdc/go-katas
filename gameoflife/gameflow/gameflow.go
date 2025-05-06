package gameflow

import (
	"reflect"
	"time"

	"github.com/laurensdc/go-katas/gameoflife/life"
	"github.com/laurensdc/go-katas/gameoflife/print"
)

func StartGame(state [][]bool, t *time.Ticker, p print.GamePrinter) {
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
