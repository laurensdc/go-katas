package print

import "fmt"

type GamePrinter interface {
	ClearScreen()
	PrintGameState(state [][]bool)
}

func GetPrinter() GamePrinter {
	return GamePrinterToConsole{}
}

type GamePrinterToConsole struct{}

func (c GamePrinterToConsole) ClearScreen() {
	fmt.Print("\033[H\033[2J") // clear screen
}

func (c GamePrinterToConsole) PrintGameState(state [][]bool) {
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
