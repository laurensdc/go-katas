package main

import (
	"strconv"
)

func main() {

}

func Counterstring(n uint) string {
	if n == 0 {
		return ""
	}

	if n == 1 {
		return "*"
	}

	// e.g. "9999"
	nStr := strconv.FormatUint(uint64(n), 10)

	// e.g. "4"
	numberOfDigits := len(nStr)

	// e.g. Counterstring(9999 - 4 - 1)
	recursiveCounterstring := Counterstring(n - uint(numberOfDigits) - 1)

	return recursiveCounterstring + nStr + "*"

}
