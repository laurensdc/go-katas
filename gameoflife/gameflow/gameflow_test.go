package gameflow

import (
	"fmt"
	"testing"
	"time"
)

type MockGamePrinter struct {
	ClearScreenCalled    bool
	PrintGameStateCalled bool
}

func (p *MockGamePrinter) ClearScreen() {
	p.ClearScreenCalled = true
}

func (p *MockGamePrinter) PrintGameState(state [][]bool) {
	p.PrintGameStateCalled = true
}

func Test_game_clears_and_prints_gamestate_on_every_tick(t *testing.T) {
	ticker := time.NewTicker(1 * time.Microsecond)
	mockPrinter := &MockGamePrinter{
		ClearScreenCalled:    false,
		PrintGameStateCalled: false,
	}
	startingState := [][]bool{
		{false, true, false},
		{false, true, false},
		{false, false, true},
	}

	StartGame(startingState, ticker, mockPrinter)

	fmt.Printf("%v", mockPrinter.ClearScreenCalled)
	fmt.Printf("%v", mockPrinter.PrintGameStateCalled)

	if mockPrinter.ClearScreenCalled == false {
		t.Errorf("Should have cleared screen")
	}

	if mockPrinter.PrintGameStateCalled == false {
		t.Errorf("Should have printed game state")
	}
}
