package counterstring

import (
	"bytes"
	"strconv"
)

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
