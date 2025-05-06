package main

import (
	"fmt"
	"github.com/laurensdc/go-katas/fizzbuzz"
)

func main() {
	runFizzBuzz()
}

func runFizzBuzz() {
	numbers := fizzbuzz.GenerateNumbers(100)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
