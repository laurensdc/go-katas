package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/laurensdc/go-katas/counterstring"
	"github.com/laurensdc/go-katas/fizzbuzz"
)

func main() {
	runCounterstring()
}

func runFizzBuzz() {
	numbers := fizzbuzz.GenerateNumbers(100)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}

func runCounterstring() {
	i := extractNumberFromArgs()
	fmt.Println(counterstring.Counterstring(uint(i)))
}

func extractNumberFromArgs() uint64 {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a number as argument")
		os.Exit(1)
	}

	input := os.Args[1]
	i, err := strconv.ParseUint(input, 10, 64)

	if err != nil {
		fmt.Printf("Could not parse number %v", input)
		os.Exit(1)
	}

	return i
}
