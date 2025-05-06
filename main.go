package main

import (
	"fmt"
	"github.com/laurensdc/go-katas/fizzbuzz"
)

func main() {
	numbers := fizzbuzz.FizzBuzz()
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}
