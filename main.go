package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/laurensdc/go-katas/counterstring"
	"github.com/laurensdc/go-katas/fizzbuzz"
	"github.com/laurensdc/go-katas/leapyear"
)

func main() {
	runLeapyear()
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

func runLeapyear() {
	i := getArgAsInput()
	fmt.Println(leapyear.IsLeapYear(i))
}

func getArgAsInput() int {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a year as argument")
		os.Exit(1)
	}

	arg := os.Args[1]
	if arg == "" {
		fmt.Println("Please provide a year as argument")
		os.Exit(1)
	}

	argI, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Could not convert %v to int\n", arg)
		os.Exit(1)
	}

	return argI
}
