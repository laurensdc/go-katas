package main

import (
	"strconv"
)

func main() {

}

func Counterstring(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "*"
	}

	strlen := len(strconv.Itoa(n))

	return Counterstring(n-strlen-1) + strconv.Itoa(n) + "*"

}
