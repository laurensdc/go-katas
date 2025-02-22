package print

import "fmt"

type GamePrinter interface {
	ClearScreen()
	PrintGameState(state [][]bool)
}

func GetPrinterToConsole() GamePrinter {
	return &PrinterToConsole{}
}

type PrinterToConsole struct{}

func (p *PrinterToConsole) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (p *PrinterToConsole) PrintGameState(state [][]bool) {
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
