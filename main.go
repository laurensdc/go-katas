package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/laurensdc/go-kata-leapyears/leapyear"
)

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

func main() {
	i := getArgAsInput()
	fmt.Println(leapyear.IsLeapYear(i))
}
