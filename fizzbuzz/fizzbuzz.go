package fizzbuzz

import (
	"strconv"
)

func GenerateNumbers(n int) []string {
	result := make([]string, n)

	for i := 0; i < len(result); i++ {
		currentNumber := i + 1
		if currentNumber%3 == 0 && currentNumber%5 == 0 {
			result[i] = "fizzbuzz"
		} else if currentNumber%3 == 0 {
			result[i] = "fizz"
		} else if currentNumber%5 == 0 {
			result[i] = "buzz"
		} else {
			result[i] = strconv.Itoa(currentNumber)
		}
	}

	return result
}
