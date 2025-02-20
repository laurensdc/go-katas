package print

import "fmt"

func PrintClearScreen() {
	fmt.Print("\033[H\033[2J") // clear screen

}

func PrintGameState(state [][]bool) {
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
