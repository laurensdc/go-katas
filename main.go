package main

import (
	"bytes"
	"strconv"
)

func main() {
	Counterstring(uint(201))

}

func Counterstring(n uint) string {
	var numbers []uint

	i := n
	for {
		if i <= 1 {
			break
		}

		numbers = append(numbers, i)

		iStr := strconv.FormatUint(uint64(i), 10)
		numberOfDigits := uint(len(iStr))

		// This calculates the next value for the "recursive" step
		// Ensure we don't underflow if currentN is small
		i = i - numberOfDigits - 1
	}

	var buf bytes.Buffer

	if i == 1 {
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
