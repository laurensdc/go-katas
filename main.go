package main

import "strconv"

func main() {

}

func Counterstring(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "*"
	}
	return Counterstring(n-2) + strconv.Itoa(n) + "*"

}
