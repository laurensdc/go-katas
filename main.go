package main

import (
	"time"

	"github.com/laurensdc/go-game-of-life/gameflow"
	"github.com/laurensdc/go-game-of-life/print"
)

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	printer := print.GetPrinter()
	gameflow.StartGame(startingState, ticker, printer)
}
