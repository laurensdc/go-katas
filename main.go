package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/laurensdc/go-game-of-life/life"
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

		printClearScreen()

		state = newState

		printGameState(state)

	}
}

func printClearScreen() {
	fmt.Print("\033[H\033[2J") // clear screen

}

func printGameState(state [][]bool) {
	for i := range state {
		for _, v := range state[i] {
			if v == true {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
