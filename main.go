package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	i := extractNumberFromArgs()
	fmt.Println(Counterstring(uint(i)))
}

func Counterstring(n uint) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "*"
	}

	numbers := generateNumberSlice(n)
	return constructString(numbers)
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

// Generate all the numbers that will be in the counterstring
func generateNumberSlice(n uint) []uint {
	var numbers []uint
	i := n

	for {
		if i <= 1 {
			break
		}

		numbers = append(numbers, i)
		numberOfDigits := numberOfDigits(i)
		i = i - numberOfDigits - 1
	}

	return numbers
}

// Reverse loop over numbers to generate the counterstring
func constructString(numbers []uint) string {
	var buf bytes.Buffer

	if numbers[len(numbers)-1] == 3 {
		buf.WriteString("*")
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		num := numbers[i]
		buf.WriteString(strconv.FormatUint(uint64(num), 10))
		buf.WriteString("*")
	}

	return buf.String()
}

func numberOfDigits(n uint) uint {
	nStr := strconv.FormatUint(uint64(n), 10)
	return uint(len(nStr))
}
