package main

import (
	"reflect"
	"time"

	"github.com/laurensdc/go-game-of-life/life"
	"github.com/laurensdc/go-game-of-life/print"
)

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		newState := life.Tick(state)

		if reflect.DeepEqual(state, newState) {
			break // game over because game is not evolving anymome
		}

		printer := print.GetPrinter()
		printer.ClearScreen()

		state = newState

		printer.PrintGameState(state)

	}
}
