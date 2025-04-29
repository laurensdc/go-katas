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
	if n >= 100 {
		return Counterstring(n-4) + strconv.Itoa(n) + "*"
	}
	if n >= 10 {
		return Counterstring(n-3) + strconv.Itoa(n) + "*"
	}
	return Counterstring(n-2) + strconv.Itoa(n) + "*"

}
