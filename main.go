package main

import (
	"fmt"
	"github.com/laurensdc/go-game-of-life/life"
	"time"
)

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	startingState := [][]bool{
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, false, true, false, false, true, false, true, false, false, true, true, false, false, true, false, true, false},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, false, true, false, false, true, false, true, false, false, true, true, false, false, true, false, true, false},
		{false, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false},
		{false, true, false, true, false, true, false, true, false, true, true, false, false, false, false, false, false, true},
		{false, false, true, false, false, true, false, true, false, false, true, true, false, false, true, false, true, false},
		{false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false},
	}

	for range ticker.C {
		fmt.Print("\033[H\033[2J") // clear screen
		printGameState(state)

		startingState = life.Tick(startingState)

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
