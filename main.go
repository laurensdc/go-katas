package main

import (
	"bytes"
	"strconv"
)

func main() {
	Counterstring(uint(8))
}

func Counterstring(n uint) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "*"
	}

	var numbers []uint

	i := n
	for {
		if i <= 1 {
			break
		}

		numbers = append(numbers, i)

		numberOfDigits := numberOfDigits(i)

		// This calculates the next value for the "recursive" step
		i = i - numberOfDigits - 1
	}

	return constructString(numbers)
}

func constructString(numbers []uint) string {
	var buf bytes.Buffer

	if numbers[len(numbers)-1] == 3 {
		buf.WriteString("*")
	}

	// Append the numbers from numsToAppend in reverse order
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
