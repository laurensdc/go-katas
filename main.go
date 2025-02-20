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

	state := [][]bool{
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, false, true, false, false, true, false, true, false, false, true, true, false, false, true, false, true, false},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, false, true, false, false, true, false, true, false, false, true, true, false, false, true, false, true, false},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false},
		{false, false, true, false, false, true, false, true, false, false, true, true, false, false, true, false, true, false},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false},
	}

	for range ticker.C {
		newState := life.Tick(state)

		if reflect.DeepEqual(state, newState) {
			break // game over because game is not evolving anymome
		}

		print.PrintClearScreen()

		state = newState

		print.PrintGameState(state)

	}
}
